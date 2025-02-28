<script lang="ts">
  import AutocompleteInput from "../AutocompleteInput.svelte";
  import { _ } from "../i18n";
  import { accounts, links, payees, tags, years } from "../stores";
  import { account_filter, fql_filter, time_filter } from "../stores/filters";

  $: fql_filter_suggestions = [
    ...$tags.map((tag) => `#${tag}`),
    ...$links.map((link) => `^${link}`),
    ...$payees.map((payee) => `payee:"${payee}"`),
  ];

  function valueExtractor(value: string, input: HTMLInputElement) {
    const match = value
      .slice(0, input.selectionStart || undefined)
      .match(/\S*$/);
    return match ? match[0] : value;
  }
  function valueSelector(value: string, input: HTMLInputElement) {
    const selectionStart = input.selectionStart ?? 0;
    const match = input.value.slice(0, selectionStart).match(/\S*$/);
    return match
      ? `${input.value.slice(
          0,
          selectionStart - match[0].length
        )}${value}${input.value.slice(selectionStart)}`
      : value;
  }

  let account_filter_value = "";
  let fql_filter_value = "";
  let time_filter_value = "";
  account_filter.subscribe((v) => {
    account_filter_value = v;
  });
  fql_filter.subscribe((v) => {
    fql_filter_value = v;
  });
  time_filter.subscribe((v) => {
    time_filter_value = v;
  });

  function submit() {
    account_filter.set(account_filter_value);
    fql_filter.set(fql_filter_value);
    time_filter.set(time_filter_value);
  }
</script>

<form on:submit|preventDefault={submit}>
  <AutocompleteInput
    bind:value={time_filter_value}
    placeholder={_("Time")}
    suggestions={$years}
    key="f t"
    clearButton={true}
    setSize={true}
    on:blur={submit}
    on:select={submit}
  />
  <AutocompleteInput
    bind:value={account_filter_value}
    placeholder={_("Account")}
    suggestions={$accounts}
    key="f a"
    clearButton={true}
    setSize={true}
    on:blur={submit}
    on:select={submit}
  />
  <AutocompleteInput
    bind:value={fql_filter_value}
    placeholder={_("Filter by tag, payee, ...")}
    suggestions={fql_filter_suggestions}
    key="f f"
    clearButton={true}
    setSize={true}
    {valueExtractor}
    {valueSelector}
    on:blur={submit}
    on:select={submit}
  />
  <button type="submit" />
</form>

<style>
  form {
    display: flex;
    flex-wrap: wrap;
    padding-top: 7px;
    margin: 0;
    color: var(--text-color);

    --placeholder-color: var(--header-placeholder-color);
    --placeholder-background: var(--header-placeholder-background);
  }

  form > :global(span) {
    max-width: 18rem;
    margin: 0 4px 6px 0;
  }

  form :global(input) {
    padding: 8px 25px 8px 10px;
    background-color: var(--background);
    border: 0;
    outline: none;
  }

  form :global([type="text"]:focus) {
    background-color: var(--background);
  }

  [type="submit"] {
    display: none;
  }

  @media print {
    form {
      display: none;
    }
  }
</style>
