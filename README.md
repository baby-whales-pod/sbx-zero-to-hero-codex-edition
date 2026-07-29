# sbx-zero-to-hero-codex-edition

✋ This hands-on uses `sbx` and the **Codex** agent (OpenAI). Authentication is done with an **OpenAI API key** stored as an `sbx` secret — the key is injected by the proxy at request time and the sandbox only ever sees a `proxy-managed` placeholder.

> Before you start, store your OpenAI API key once:
> ```bash
> sbx secret set -g openai
> ```
> (Paste your `sk-...` key when prompted.)
>
> 📖 Reference: <https://docs.docker.com/ai/sandboxes/agents/codex/>
