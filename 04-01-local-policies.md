# Local policies
> `sbx` without ai governance policy

> [!NOTE]
> In this lesson you will inspect the network policy rules that apply to a sandbox, then deny and allow outbound domains — both globally and scoped to a single sandbox — and read the structured 403 explanation when a request is blocked.

**Local policies** control what a sandbox is allowed to access over the network. A rule can apply:
- **globally** to all sandboxes (default)
- or be **scoped** to a single sandbox with `--sandbox`

> Every blocked outbound request returns an **HTTP 403** with a structured explanation (`rule`, `origin`, `detail`).

We reuse the sandbox created earlier: `codex-hello-sbx`.

```bash
sbx ls
```

## 1. Look at the policies in place

Overview of the policies:
```bash
sbx policy ls
sbx policy ls --type network
sbx policy ls --type filesystem
```

The rules that apply to a given sandbox:
```bash
sbx policy ls codex-hello-sbx
```


## 2. Test an access *before* changing the policy

```bash
curl https://example.com
```

Now let's check from inside the sandbox with a real request:

```bash
sbx exec codex-hello-sbx -- curl -sS -o /dev/null -w "%{http_code}\n" https://example.com
```

## 3. Block a domain (deny) and observe the 403

Add a global **deny** rule:

```bash
sbx policy deny network example.com
```

Replay the request: it is now blocked.

```bash
sbx exec codex-hello-sbx -- curl -sS https://example.com
```

You should get a **403** with the explanatory body:

```text
Blocked by network policy: domain example.com
  rule:   "..." (domain, deny)
  origin: local policy
  detail: ...
```

> ✋ **deny** rules always take precedence over **allow** rules (check with the `sbx` dashboard).

Ask the agent (Codex) to try accessing the blocked domain:

```text
please run: curl -sS https://example.com
```

The agent will get the 403 and the body explaining that the network policy blocks the request.


## 4. Scope a rule to a single sandbox

A rule can apply to **only one** sandbox thanks to `--sandbox` (it is then added to that sandbox's `local` policy):

```bash
# blocked only for codex-hello-sbx
sbx policy deny network --sandbox codex-hello-sbx httpbin.org
```

Verification:
```bash
# blocked in the sandbox context
sbx policy check network --sandbox codex-hello-sbx httpbin.org

# but not globally, check with
sbx policy ls
```

Test the request: it is now blocked.

```bash
sbx exec codex-hello-sbx -- curl -sS https://httpbin.org
```

> 👋👀 have look to the `sbx` dashboard

## 5. Allow a domain (allow)

Conversely, open access to a domain:

```bash
# a specific domain
sbx policy allow network api.example.com

# all subdomains
sbx policy allow network "*.npmjs.org"

# only for one sandbox
sbx policy allow network --sandbox codex-hello-sbx api.example.com
```

## 6. Audit: see what was allowed / blocked

```bash
# all logs
sbx policy log

# logs for one sandbox
sbx policy log codex-hello-sbx

# the last 20 entries
sbx policy log --limit 20
```


## 7. Cleanup

List the rules to find their resource or `rule ID`:
```bash
sbx policy ls
```

Remove the rules added during the exercise:
```bash
# global rule, by resource
sbx policy rm network --resource example.com

# rule scoped to a sandbox
sbx policy rm network --sandbox codex-hello-sbx --resource httpbin.org
```

> To start over from scratch on the whole config: `sbx policy reset` (⚠️ stops running sandboxes and prompts you to initialize the global policy again via `sbx policy init`).
