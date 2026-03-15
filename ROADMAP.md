# Roadmap

This roadmap captures priorities for improving adoption and production readiness.

## Near term

- Stabilize API semantics and error behavior for composed runners.
- Add more examples for common service stacks (HTTP + queue + metrics).
- Add benchmark coverage for startup/shutdown orchestration.

## Mid term

- Add context-aware lifecycle helpers (`Runner` adapters with `context.Context`).
- Add optional observability hooks for startup/shutdown timing.
- Improve Windows signal behavior documentation and tests.

## Long term

- Publish migration notes for major versions.
- Grow integrations with common Go stacks and templates.
