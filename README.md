# sbx — Zero to Hero · Workshop (Codex edition)

> University of Edinburgh 2026-08-25
> 
> - 🟢 [00-00.requirements](./00-00.requirements)
> - 🟢 [01-01.create-the-sandbox.md](./01-01.create-the-sandbox.md)
> - 🟢 [01-02.walk-into-the-sandbox.md](./01-02.walk-into-the-sandbox.md)
> - 🟢 [02-01.github-interaction.md](./02-01.github-interaction.md)
> - 🟢 [03-02.secrets-injection.md](./03-02.secrets-injection.md)
> - 🟢 [07-01.local-mcp-gateway.md](./07-01.local-mcp-gateway.md)
> 
> If you finish before the end
> 
> - 🟠 [03-01.env-vars-injection.md](./03-01.env-vars-injection.md)
> - 🟠 [05-01.kits-add-your-tools.md](./05-01.kits-add-your-tools.md)

<hr>

✋ This hands-on uses `sbx` and the **Codex** agent (OpenAI). Authentication is done with an **OpenAI API key** stored as an `sbx` secret — the key is injected by the proxy at request time and the sandbox only ever sees a `proxy-managed` placeholder.

> Before you start, store your OpenAI API key once:
> ```bash
> sbx secret set -g openai
> ```
> (Paste your `sk-...` key when prompted.)
>
> 📖 Reference: <https://docs.docker.com/ai/sandboxes/agents/codex/>

> ⚠️ **Living document.** This README describes the intended scope and flow of the hands-on. The content is a work in progress: modules, ordering, commands and examples **may change**, and the workshop can be **adapted** on the fly to the audience, the available time, and the `sbx` version in use (some features are still **experimental**). Treat section durations as indicative, not contractual.
>
> In particular:
> - The content **must be adapted to the AI agent that is selected**. This edition uses the **Codex** agent (OpenAI), but `sbx` can run other agents — the login flow, prompts and some commands will differ accordingly.
> - The **number of steps and modules will vary with the length of the workshop**: a short session covers only the first modules, a full session goes all the way through. Trainers pick, trim or reorder steps to fit the time available.

## What is this workshop?

A practical, "learn by doing" introduction to **Docker Sandboxes (`sbx`)** — a way to run AI coding agents (here the **Codex** agent from OpenAI) inside isolated, policy-controlled VMs instead of directly on your machine.

Throughout the hands-on you drive a real agent inside a sandbox, publish ports, inject environment variables and secrets, control what the sandbox can reach on the network and on disk, and finally package and share your own reproducible sandbox environments.

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

## Viewing this material in VS Code

Two VS Code extensions are needed to get the full experience. Both are already listed in [`.vscode/extensions.json`](./.vscode/extensions.json), so VS Code offers to install them when you open this folder ("Show Recommendations").

| What | Extension | Why |
|---|---|---|
| **Diagrams** — `.assets/*.drawio.svg` | **Draw.io Integration** (`hediet.vscode-drawio`) | Open, zoom and **edit** the diagrams. They are SVG files with the draw.io model embedded, so the markdown preview displays them without any extension, but you need this extension to modify them (or open them at <https://app.diagrams.net>). |
| **Slides** — `presentation/*.md` | **Marp for VS Code** (`marp-team.marp-vscode`) | The decks are [Marp](https://marp.app/) markdown (`marp: true` front-matter). Without the extension the preview shows raw markdown instead of slides. |

```bash
# install from the command line if you prefer
code --install-extension hediet.vscode-drawio
code --install-extension marp-team.marp-vscode
```

> 💡 The lesson diagrams also keep their original **mermaid** source, commented out right above each image, in case you want to compare or regenerate them.
>
> 💡 To export a deck without VS Code: `npx @marp-team/marp-cli presentation/00-intro.md --pdf`.

## Workshop outline

### Module 0 — Introduction & setup *(planned / evolving)*
- Why sandboxes: the "container escape" / blast-radius problem for AI agents — [`00-01.container-escape.md`](./00-01.container-escape.md).
- What `sbx` is ([`00-02.sbx-introduction.md`](./00-02.sbx-introduction.md)) and how you log in ([`00-03.sbx-login.md`](./00-03.sbx-login.md)).
- Requirements check — [`00-00.requirements.md`](./00-00.requirements.md).

### Module 1 — Create and explore a sandbox
- `sbx run codex` — create your first sandbox; the OpenAI API key is injected by the proxy (no in-sandbox login) — [`01-01.create-the-sandbox.md`](./01-01.create-the-sandbox.md).
- Walk into the sandbox shell, inspect its filesystem, run the agent on simple tasks — [`01-02.walk-into-the-sandbox.md`](./01-02.walk-into-the-sandbox.md).

### Module 2 — Interacting with GitHub
- Ask the agent to scaffold a small Go microservice with a `/hello` JSON endpoint, **publish a port** to the host (`sbx ports … --publish`) and test with `curl`, then commit & push (via the GitHub token injected by the proxy) and create a GitHub issue from the agent with the `gh` CLI — [`02-01.github-interaction.md`](./02-01.github-interaction.md).

### Module 3 — Environment variables & secrets
- Inject a persistent environment variable into a running sandbox — [`03-01.env-vars-injection.md`](./03-01.env-vars-injection.md).
- **Custom secret injection**: the sandbox sees only a *placeholder*, while the proxy substitutes the real value into outbound requests — demonstrated against a local `tiny-service` that validates a bearer token — [`03-02.secrets-injection.md`](./03-02.secrets-injection.md).

### Module 4 — Policies: controlling access
- **Network policies**: allow/deny domains, global vs per-sandbox scope, observe the structured **HTTP 403**, audit with `sbx policy log` — [`04-01-local-policies.md`](./04-01-local-policies.md).
- **Workspaces**: mount host directories read-write / read-only, prove the rest of the host stays invisible — [`04-02-workspaces.md`](./04-02-workspaces.md).
- **Filesystem policies**: read vs write rules, and how read-only workspaces are the local way to restrict writes — [`04-03-filesystem-policies.md`](./04-03-filesystem-policies.md).

### Module 5 — Customize & share environments
- **Kits**: declarative builder/launcher add-ons authored with the **v2 kit schema**; share via ZIP, OCI registry, or git *(experimental)* — [`05-01.kits-add-your-tools.md`](./05-01.kits-add-your-tools.md).
- **Snapshots**: freeze a prepared sandbox into a reusable template image — [`05-02.snapshots.md`](./05-02.snapshots.md).
- **Custom template from a Dockerfile**: build a reproducible base image — [`05-03.custom-template-dockerfile.md`](./05-03.custom-template-dockerfile.md).
- **Build a template inside a sandbox**: when the host has no Docker, use the sandbox's own Docker engine — [`05-04.build-template-in-sandbox.md`](./05-04.build-template-in-sandbox.md).

### Module 6 — Advanced Git workspaces
- **Clone mode** (`--clone`): the agent works on a private in-container clone of your repo instead of your live files; sync commits back with the `sandbox-<name>` remote — [`06-01.clone-mode.md`](./06-01.clone-mode.md).


### Module 7 — MCP: giving the agent tools
- **Local MCP gateway**: force the local MCP gateway, register public no-auth MCP servers (DeepWiki, Context7) and use them from the agent — [`07-01.local-mcp-gateway.md`](./07-01.local-mcp-gateway.md). *(requires a recent `sbx` build)*

### Module 8 — SSH access
- **SSH into a sandbox**: turn a sandbox into a first-class SSH host (`ssh <name>.sbx`) with `sbx setup ssh`, run shells and one-shot commands, and open it in VS Code / Cursor Remote-SSH — [`08-01.ssh-access.md`](./08-01.ssh-access.md). *(experimental, off by default; requires a recent `sbx` build)*

### Topics on the roadmap *(may or may not be covered)*
AI governance policies, audit, git worktrees, and multiple workspaces — these are being drafted and may be added or adapted per session.

## Format & logistics

- **Style:** fully hands-on; each participant works in their own `hello-sbx` repo.
- **Duration:** modular — a short session can cover Modules 1–3; a full session goes through Module 5. Trainers pick and reorder modules to fit the group.
- **Cleanup:** most lessons end with cleanup commands (`sbx rm …`, `sbx template rm …`).

---
*Because `sbx` is evolving quickly, always cross-check commands against the official docs: <https://docs.docker.com/ai/sandboxes/>. This workshop was validated on `sbx` v0.37.0 and will be updated as the tooling and the hands-on mature.*
