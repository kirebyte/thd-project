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

## Why is the `sqlite` repository under `internal/`?

The `internal/` directory in Go signals:  
> “This package should not be imported from outside this module.”

The `sqlite` repository implementation is tied to the current app and domain.  
It's not a reusable library or part of a public interface — it's an internal infrastructure detail.

By placing it in `internal/`, we:
- Prevent accidental reuse outside the app
- Respect Go’s visibility semantics
- Keep a clear boundary between public contracts and private implementations

If we ever need to share or extract it, we can move it later.  
It's safer to start private and open up later, than the other way around.

In contrast, the `CarRepository` interface lives at the root because it defines a contract that may be consumed across packages or layers.

---

## Who owns the data model? Should it live in a shared library?

We believe the **application that exposes the contract should own the model**.

This gives full control over:
- Versioning
- Validation
- Evolution of business logic

Shared model libraries (a.k.a. "backend models") sound useful but usually cause more problems than they solve:

- Tight coupling between unrelated apps
- Breaking changes that impact multiple teams
- Unclear ownership and review responsibilities

Instead, we define models **per app**, in a clear and accessible package (like `/model`).  
This keeps contracts visible and reduces surprise dependencies.

If another service needs the model, they can:
- Use our OpenAPI definition
- Generate their own types
- Or copy the shape explicitly

---

## What about shared constants or string tables?

Sometimes teams try to centralize error messages or validation strings in shared files.

While this can seem clean, it introduces similar issues:
- Changes in wording can trigger redeployments of unrelated apps
- Shared files become bloated and hard to manage
- Translation or customization becomes harder

Our approach:
- **Error messages live with the package that generates them**
- If needed, we define them as constants for reuse *within the app*
- Universal messages (if any) should be extracted into a real standard library, not a string dump

We aim for **clarity, context, and independence**, even if that means a bit of duplication.

---