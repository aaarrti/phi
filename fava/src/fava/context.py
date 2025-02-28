"""Specify types for the flask application context."""
from __future__ import annotations

import flask

from .core import FavaLedger
from .util.date import Interval


class Context:
    """The allowed context values."""

    beancount_file_slug: str | None
    conversion: str
    interval: Interval
    ledger: FavaLedger


g: Context = flask.g  # type: ignore
