# 🧠 Architecture FAQ

Quick notes on why things are structured the way they are.

---

## Why does `cmd/dbinit` use `log.Println` instead of our custom logger?

`cmd/dbinit` is a standalone bootstrap tool. We use `log.Println` and `log.Fatal` because:
- They're always visible, no matter what `LOG_LEVEL` is.
- They don't depend on extra init logic.
- It's safer during early startup where logging must never fail.

---

## Then why does it still import `settings`?

Because `settings` is our single source of truth for configuration across all binaries.

Even for standalone tools, this keeps:
- Environment variable names consistent
- Default values shared
- Configuration logic DRY

---

## Could we move everything to `internal`?

We only move what's meant to be reused *internally* by the app.  
`cmd/dbinit` is a one-shot tool, so it doesn't need to go inside `internal/`.

`settings` is used broadly, so it stays as a shared package.

---

## Should we log in JSON, or use levels like `INFO`, `DEBUG`?

We're using a minimal custom logger for now.  
If needed later, we can add:
- Structured logging
- JSON output
- Timestamps
- Log levels

For now, we keep it simple.

---

## Why SQLite instead of in-memory?

The original spec allowed in-memory, but we chose SQLite for:
- Realistic persistence
- Better local dev experience
- Cleaner separation of storage concerns

Still no external dependency: we use `modernc.org/sqlite` (pure Go).

---
