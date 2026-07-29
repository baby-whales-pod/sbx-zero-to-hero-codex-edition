# sbx — Zero to Hero · Workshop Plan (EN)

> ⚠️ **Living document.** This plan describes the intended scope and flow of the hands-on. The content is a work in progress: modules, ordering, commands and examples **may change**, and the workshop can be **adapted** on the fly to the audience, the available time, and the `sbx` version in use (some features are still **experimental**). Treat section durations as indicative, not contractual.
>
> In particular:
> - The content **must be adapted to the AI agent that is selected**. This edition uses the **Codex** agent (OpenAI), but `sbx` can run other agents — the login flow, prompts and some commands will differ accordingly.
> - The **number of steps and modules will vary with the length of the workshop**: a short session covers only the first modules, a full session goes all the way through. Trainers pick, trim or reorder steps to fit the time available.

## What is this workshop?

A practical, "learn by doing" introduction to **Docker Sandboxes (`sbx`)** — a way to run AI coding agents (here the **Codex** agent from OpenAI) inside isolated, policy-controlled VMs instead of directly on your machine.

Throughout the hands-on you drive a real agent inside a sandbox, publish ports, inject environment variables and secrets, control what the sandbox can reach on the network and on disk, and finally package and share your own reproducible sandbox environments.

> ✋ This hands-on uses `sbx` and the **Codex** agent authenticated with an **OpenAI API key**. Store it once with `sbx secret set -g openai`; the proxy injects it at request time and the sandbox only ever sees a `proxy-managed` placeholder.

## Who is it for?

- Developers and platform/DevOps engineers curious about running AI agents safely.
- No deep Docker expertise required, but basic terminal comfort is expected.
- The demo app is a tiny **Go** microservice — no Go knowledge needed.

## Learning goals

By the end, participants can:
1. Create a sandbox and run an AI agent inside it.
2. Let the agent build, run and expose a small service, then commit/push to GitHub.
3. Inject environment variables and **secrets** the agent never actually sees in clear text.
4. Restrict a sandbox with **network** and **filesystem** policies.
5. Customize and share sandbox environments via **kits**, **snapshots** and **templates**.

## Prerequisites

- A GitHub account (handle) and a repository named `hello-sbx` to work in.
- An **OpenAI API key** (`sk-...`), stored once with `sbx secret set -g openai`.
- `sbx` installed and set up — see <https://docs.docker.com/ai/sandboxes/#get-started>.
- VS Code + a terminal.

## Workshop outline

### Module 0 — Introduction & setup *(planned / evolving)*
- Why sandboxes: the "container escape" / blast-radius problem for AI agents.
- What `sbx` is and how you log in.
- Requirements check.

### Module 1 — Create and explore a sandbox
- `sbx run codex` — create your first sandbox; the OpenAI API key is injected by the proxy (no in-sandbox login).
- Walk into the sandbox shell, inspect its filesystem, run the agent on simple tasks.

### Module 2 — Interacting with GitHub
- Ask the agent to scaffold a small Go microservice with a `/hello` JSON endpoint.
- **Publish a port** to the host (`sbx ports … --publish`) and test with `curl`.
- Commit & push the code (via the GitHub token injected by the proxy), and create a GitHub issue from the agent with the `gh` CLI.

### Module 3 — Environment variables & secrets
- Inject a persistent environment variable into a running sandbox.
- **Custom secret injection**: the sandbox sees only a *placeholder*, while the proxy substitutes the real value into outbound requests — demonstrated against a local `tiny-service` that validates a bearer token.

### Module 4 — Policies: controlling access
- **Network policies** (`04-01`): allow/deny domains, global vs per-sandbox scope, observe the structured **HTTP 403**, audit with `sbx policy log`.
- **Workspaces** (`04-02`): mount host directories read-write / read-only, prove the rest of the host stays invisible.
- **Filesystem policies** (`04-03`): read vs write rules, and how read-only workspaces are the local way to restrict writes.

### Module 5 — Customize & share environments
- **Kits** (`05-01`): declarative builder/launcher add-ons authored with the **v2 kit schema**; share via ZIP, OCI registry, or git. *(experimental)*
- **Snapshots** (`05-02`): freeze a prepared sandbox into a reusable template image.
- **Custom template from a Dockerfile** (`05-03`): build a reproducible base image.
- **Build a template inside a sandbox** (`05-04`): when the host has no Docker, use the sandbox's own Docker engine.

### Topics on the roadmap *(may or may not be covered)*
AI governance policies, audit, **clone mode**, git worktrees, and multiple workspaces — these are being drafted and may be added or adapted per session.

## Format & logistics

- **Style:** fully hands-on; each participant works in their own `hello-sbx` repo.
- **Duration:** modular — a short session can cover Modules 1–3; a full session goes through Module 5. Trainers pick and reorder modules to fit the group.
- **Cleanup:** most lessons end with cleanup commands (`sbx rm …`, `sbx template rm …`).

---
*Because `sbx` is evolving quickly, always cross-check commands against the official docs: <https://docs.docker.com/ai/sandboxes/>. This plan will be updated as the tooling and the hands-on mature.*
