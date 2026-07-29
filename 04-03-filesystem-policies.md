# Local filesystem policies
> `sbx` without ai governance policy

> [!NOTE]
> In this lesson you will inspect the filesystem policies that govern what a sandbox can read and write, understand the default "read anything, write anything" posture, and see how a read-only workspace is the practical, local way to restrict writes today.

Just like **network policies** (`04-01`) control what a sandbox can reach on the network, **filesystem policies** describe what a sandbox is allowed to **read** and **write** on disk.

There are two filesystem rule types:
- `filesystem:read`
- `filesystem:write`

By default, `sbx` ships two permissive **local** rules that apply to every sandbox:

| rule                         | type               | decision | resources |
|------------------------------|--------------------|----------|-----------|
| `default-fs-read-allow-all`  | `filesystem:read`  | allow    | `**`      |
| `default-fs-write-allow-all` | `filesystem:write` | allow    | `**`      |

> ✋ **Local scope = read-only on the CLI here.** Unlike network rules, you cannot *add* custom `allow`/`deny` filesystem rules with `sbx policy allow/deny` (those subcommands only accept `network`). Custom filesystem restrictions are delivered by **remote governance policies**. Locally, the way you restrict what a sandbox can write is by mounting a workspace **read-only** — see `04-02-workspaces.md`.

We reuse the sandbox from the previous lesson. If you removed it, recreate it:

```bash
sbx run codex . ../shared-docs:ro --name codex-workspaces
sbx ls
```

## 1. Look at the filesystem policies in place

Global overview:
```bash
sbx policy ls --type filesystem
```

You should see the two `default-fs-*` rules, both `allow **`.

The rules that apply to a given sandbox (global + scoped):
```bash
sbx policy ls codex-workspaces --type filesystem
```

## 2. Understand read vs write

- `filesystem:read` → can the sandbox *read* a path
- `filesystem:write` → can the sandbox *write* a path
- resource `**` → "any path"
- as with network, a **deny** rule always wins over an **allow** rule

So the default posture is: **read anything, write anything** — *within the mounted workspaces only*. The sandbox still cannot see host paths that were never mounted (that is enforced by the container boundary, not by a policy).

## 3. Read-only workspace = a write restriction in practice

In `04-02` the `:ro` workspace refused writes. That is the concrete, local filesystem control you have today:

```bash
# allowed: read the read-only workspace
sbx exec codex-workspaces -- cat ../shared-docs/REFERENCE.md

# refused: write the read-only workspace (Read-only file system)
sbx exec codex-workspaces -- sh -c 'echo x >> ../shared-docs/REFERENCE.md'
```

> Workspaces mount at the same absolute path as on the host; `sbx exec` starts in the primary workspace, so `../shared-docs` is the read-only sibling.

## 4. Audit note

Network decisions are logged and queryable:
```bash
sbx policy log codex-workspaces
```

Filesystem policy **logging is not implemented yet** — this returns a notice, not entries:
```bash
sbx policy log --type filesystem
# -> filesystem policy log entries are not supported yet
```

