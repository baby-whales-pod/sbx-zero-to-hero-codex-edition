# Handoff — sbx-zero-to-hero-codex-edition

> Working notes to resume this workshop's authoring after an `sbx` restart / new session. Not part of the workshop content itself.

## What this repo is

A hands-on workshop for **Docker Sandboxes (`sbx`)** driving the **Codex** agent (OpenAI), authenticated with an **OpenAI API key** (`sbx secret set -g openai`). It was converted from an earlier edition that used the GitHub Copilot CLI with the OAuth device flow. Kits are authored with the **v2 kit schema** (`schemaVersion: "2"`). The sbx source of truth is checked out under `./sandboxes` (read-only reference).

## ▶️ Resume here first (open task)

We were debugging `sbx create --name tiny-service shell .` (lesson `03-02.secrets-injection.md`): the command printed two binding warnings and, per the user, **the sandbox was not created**.

- The two warnings (`credential for "github"/"openai" discovered but no domains allowed by your bindings; not injecting`) are **expected and harmless here**: the `shell` template is v2 and declares many services (incl. github/openai, which are stored globally but not bound). Confirmed in `sandboxes/sandboxlib/bindingconsent/service.go` that this path **warns but does NOT block** creation. `tiny-service` needs neither — its secret is the custom `TINY_SERVICE_TOKEN` created later in the lesson.
- So the "not created" symptom is a **separate** issue. Next diagnostic steps for the user after restart:
  - `sbx ls` (or the `sbx` dashboard) — `sbx create` does not attach, so it may already exist silently.
  - stale name collision → `sbx rm tiny-service` then retry.
  - real error → `sbx create -D --name tiny-service shell .` and read the last line.
- **Waiting on the user** to paste the last line of `sbx create -D …` or the output of `sbx ls`.

## Done this session

- Full conversion to the **Codex / API-key** edition: `sbx run codex`, base image `docker/sandbox-templates:codex-docker`, default sandbox `codex-hello-sbx`; all `copilot`→`codex` renames across every lesson, `README.md`, `PLAN.en.md`, `PLAN.fr.md`.
- Kits moved to **v2**: `kits/tools-kit/spec.yaml`, `kits/serve-kit/spec.yaml`, and the inline specs in `05-01.templates-and-kits.md`; `custom-template/Dockerfile` now `FROM docker/sandbox-templates:codex-docker`.
- Filled the previously-empty module-0 docs: `00-01.container-escape.md` (escape demo + Linux/macOS/Windows host paths), `00-02.sbx-introduction.md` (what/why + security architecture, with Mermaid + editable draw.io in `.assets/sbx-architecture.drawio` and `.assets/sbx-credential-injection.drawio`), `00-03.sbx-login.md` (what `sbx login` does).
- `00-00.requirements.md`: per-OS install (no Docker Desktop required) + **GitHub PAT injection** section. `02-01` now just links to it.
- Bindings-v2 troubleshooting extracted to `troubleshooting/openai-credential-no-domains-allowed.md` (linked from `01-01`).
- `01-01` gained a **Reconnect to the sandbox** section; `01-02` gained a **host vs sandbox filesystem** explanation; `> [!NOTE]` intros added to `01-01`, `01-02`, `02-01`, `03-01`; "why" sections added to `03-01` and `03-02`; `03-02` explains the `shell` template (agent-less) + that Go is baked into the base image.

## Conventions / directives (persistent)

- **No hard-wrapped prose** in any Markdown file: one continuous line per paragraph and per list item; newlines only for real structure (blank lines, list items, headings, code fences, blockquotes, tables, and GitHub `[!NOTE]`-style alert markers on their own line). This rule is committed in **`AGENTS.md`** (survives recreation) and also stored in agent memory.
- A throwaway reflow-checker lived at `/tmp/reflow.py` (gone after restart); the rule itself is in `AGENTS.md`.

## GitHub issues opened (repo `baby-whales-pod/sbx-zero-to-hero-codex-edition`, label `documentation`, type `Task`)

- #1 — verify sbx install steps on Linux and Windows (`00-00`).
- #2 — verify the container-escape demo on Linux and Windows (`00-01`).
- #3 — verify non-interactive `sbx login --username <user> --password-stdin` (`00-03`).

## Pending / offered but not done

- Add a note in `03-02` that the github/openai binding warnings are expected during `sbx create … shell` (offered, awaiting go-ahead).
- Possible issues: bindings-v2 / `credentials.yaml` flow is undocumented on docs.docker.com; validate the `yq` binding commands on a real install.
- `hello-sbx/README.md` is hard-wrapped but is the throwaway sample repo (not authored workshop content) — left untouched.
- `.assets/github-login-*.png` are orphaned Copilot-era screenshots — left in place, unreferenced.
- Nothing has been committed to git this whole session.

## Environment notes

- This is a Claude Code sandbox named `claude-sbx-zero-to-hero-codex-edition`.
- GitHub API works from here because the user set a **sandbox-scoped** secret: `sbx secret set claude-sbx-zero-to-hero-codex-edition github -t "$(gh auth token)"` (global `-g` only applies to newly-created sandboxes). Issues were created via the REST API through the proxy (`gh` itself reports "not logged in", which is expected).
