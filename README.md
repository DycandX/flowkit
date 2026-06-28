# flowkit

Scaffold project workflow in one command.

Generate `WORKFLOW.md`, CI pipelines, and git hooks — with auto-detected tech stack and workflow style.

```bash
flowkit init
```

## Install

**Go users:**
```bash
go install github.com/DycandX/flowkit@latest
```

**npm users:**
```bash
npx flowkit init
```

**Binary download:** download from [GitHub Releases](https://github.com/DycandX/flowkit/releases).

## Usage

```bash
flowkit init
```

Interactive prompts:
1. Project name
2. Main branch
3. Language (en/id)
4. Workflow style (GitFlow / GitHub Flow / Trunk-Based)
5. Generate CI? Hooks?

Non-interactive:
```bash
flowkit init --project-name "my-app" --main-branch master --force
```

## Output

| File | Description |
|------|-------------|
| `WORKFLOW.md` | Git workflow documentation |
| `.github/workflows/ci.yml` | CI pipeline per tech stack |
| `.github/workflows/pr-check.yml` | Enforces branch & commit convention |
| `.husky/pre-commit` | Pre-commit hook (JS) / `.githooks/` (non-JS) |
| `.husky/commit-msg` | Commit message lint hook |
| `commitlint.config.js` | Commit convention config (JS) |
| `.lintstagedrc.json` | Lint-staged config (JS) |
| `flowkit.json` | Saved config for re-run |

## Detected Stacks

next, react, vue, nuxt, node, go, rust, python, laravel, unknown

## Workflow Styles

- **GitFlow** — `main` + `develop`, feature branches → develop → release
- **GitHub Flow** — `main` only, feature branches → PR → main → deploy
- **Trunk-Based** — single branch, short-lived branches < 1 day

## License

MIT
