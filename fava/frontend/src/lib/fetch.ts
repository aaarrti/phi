import { isJsonObject, object, unknown } from "./validation";

class FetchError extends Error {}

/** Wrapper around fetch with some default options */
export function fetch(
  input: string,
  init: RequestInit = {}
): Promise<Response> {
  return window.fetch(input, { credentials: "same-origin", ...init });
}

/**
 * Handles JSON content for a Promise returned by fetch, also handling an HTTP
 * error status.
 */
async function handleJSON(response: Response): Promise<unknown> {
  if (!response.ok) {
    throw new FetchError(response.statusText);
  }
  const data: unknown = await response.json();
  if (!isJsonObject(data)) {
    throw new FetchError("Invalid response: not an object");
  }
  if (!data.success) {
    throw new FetchError(
      typeof data.error === "string"
        ? data.error
        : "Invalid response: missing error"
    );
  }
  return data;
}

const response_validator = object({ data: unknown });

export async function fetchJSON(
  input: string,
  init?: RequestInit
): Promise<unknown> {
  const res = await fetch(input, init).then(handleJSON);
  const parsed = response_validator(res);
  if (parsed.success) {
    return parsed.value.data;
  }
  throw new FetchError("Invalid response: missing data key.");
}

/**
 * Handles text content for a Promise returned by fetch, also handling an HTTP
 * error status.
 */
export async function handleText(response: Response): Promise<string> {
  if (!response.ok) {
    const msg = await response.text();
    throw new Error(msg || response.statusText);
  }
  return response.text();
}