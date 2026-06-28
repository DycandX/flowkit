# flowkit

> Scaffold project workflow in one command — `WORKFLOW.md`, CI pipelines, git hooks.

[![CI](https://github.com/DycandX/flowkit/actions/workflows/ci.yml/badge.svg)](https://github.com/DycandX/flowkit/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/DycandX/flowkit)](go.mod)
[![License](https://img.shields.io/github/license/DycandX/flowkit)](LICENSE)
[![npm](https://img.shields.io/npm/v/@dycandx/flowkit)](https://www.npmjs.com/package/@dycandx/flowkit)

```text
$ flowkit init

? Project name: my-app
? Main branch: master
? Language: en
? Workflow style: GitFlow

✓ WORKFLOW.md
✓ .github/workflows/ci.yml
✓ .github/workflows/pr-check.yml
✓ .husky/pre-commit
✓ .husky/commit-msg
✓ commitlint.config.js
✓ flowkit.json
```

## Table of Contents

- [Why flowkit](#why-flowkit)
- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [CLI Reference](#cli-reference)
- [What Gets Generated](#what-gets-generated)
- [Stack Detection](#stack-detection)
- [Workflow Styles](#workflow-styles)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Why flowkit

Every project needs a git workflow, but teams waste hours on the same setup:

- **Document it** → someone forgets to read it
- **Copy-paste templates** → error-prone (`sed` breaking on paths)
- **Scaffold tools** → too opinionated or JS-only

flowkit solves this by generating **three layers** in one command:

| Layer | What | File |
|-------|------|------|
| 📄 Document | Git workflow reference | `WORKFLOW.md` |
| 🤖 Automate | CI pipeline that runs on push | `.github/workflows/ci.yml` |
| 🔒 Enforce | Git hooks + PR checks | `.husky/*` / `.githooks/*`, `pr-check.yml` |

Zero runtime dependency. Single binary. Cross-platform.

## Features

- **Interactive CLI** — prompts with sensible defaults, no flags to memorize
- **Stack auto-detection** — reads `package.json`, `Cargo.toml`, `go.mod`, and more
- **3 workflow styles** — GitFlow, GitHub Flow, Trunk-Based
- **CI templates per stack** — correct pipeline for Next.js, Go, Rust, Python, Laravel, etc.
- **PR enforcement** — branch naming and commit message convention checks
- **Git hooks** — pre-commit linting + commit-msg validation
- **Config re-run** — saves `flowkit.json` so you can regenerate later
- **Cross-platform** — Windows, macOS, Linux (single binary or `npx`)

## Installation

### Go (recommended)

```bash
go install github.com/DycandX/flowkit@latest
```

### npm

```bash
npx @dycandx/flowkit init
```

Globally:

```bash
npm install -g @dycandx/flowkit
```

### Binary

Download from [GitHub Releases](https://github.com/DycandX/flowkit/releases), extract, and add to `PATH`.

### From source

```bash
git clone https://github.com/DycandX/flowkit.git
cd flowkit
go build -o /usr/local/bin/flowkit .
```

## Quick Start

```bash
cd /path/to/your-project

flowkit init
```

Answer the prompts:

```text
? Project name: my-app
? Main branch (master): master
? Language: en
? Workflow style: GitFlow
```

Review the generated files, then:

```bash
git add .
git commit -m "chore: add workflow scaffold"
git push
```

Your CI pipeline runs automatically. Every future commit must pass lint, build, and conventional commit format.

## CLI Reference

### `flowkit init`

Scaffold workflow files for the current directory.

```text
Usage:
  flowkit init [flags]

Flags:
  -f, --force   Overwrite existing files
  -h, --help    Help for init
```

### `flowkit init --non-interactive`

```bash
flowkit init \
  --project-name "my-app" \
  --main-branch master \
  --language en \
  --workflow-style gitflow
```

**Aliases:** `--project-name` also accepts `-n`, `--main-branch` also `-b`, `--language` also `-l`, `--workflow-style` also `-w`.

## What Gets Generated

```
my-project/
├── WORKFLOW.md                    Git workflow documentation
│
├── .github/
│   └── workflows/
│       ├── ci.yml                 CI pipeline (stack-aware)
│       └── pr-check.yml           Enforces branch + commit conventions
│
├── .husky/                        (JS/Node projects)
│   ├── pre-commit                 Runs lint-staged before each commit
│   └── commit-msg                 Validates conventional commit format
│
├── .githooks/                     (non-JS projects)
│   ├── pre-commit                 Shell-based debug-code check
│   └── commit-msg                 Regex-based conventional commit check
│
├── commitlint.config.js           Conventional commit rules (JS only)
├── .lintstagedrc.json             Lint-staged configuration (JS only)
├── flowkit.json                   Saved config for re-generation
└── WORKFLOW.md                    Reusable workflow doc
```

### output example

**`WORKFLOW.md`:** Branching diagram, naming conventions, commit rules, step-by-step task workflow, code quality checklist — all customized with your project name and commands.

**`ci.yml`:** GitHub Actions workflow that runs lint, build, and test on every push. Cached dependencies. Correct per stack (npm, cargo, go, pip).

**`pr-check.yml`:** Validates branch names match `feat/`, `fix/`, `chore/`, etc. and commit messages follow conventional commits format.

**`pre-commit`:** Stops bad commits before they happen. JS projects get `lint-staged`; non-JS projects get a shell script that flags debug code.

**`commit-msg`:** Ensures every commit message matches `<type>(<scope>): <description>`.

## Stack Detection

flowkit reads your project's language files and picks the right CI template:

| File | Detected Stack | CI Template |
|------|----------------|-------------|
| `package.json` + `"next"` | Next.js | npm ci → lint → build |
| `package.json` + `"react"` | React | npm ci → lint → build |
| `package.json` + `"vue"` | Vue | npm ci → lint → build |
| `package.json` + `"nuxt"` | Nuxt.js | npm ci → lint → build |
| `package.json` (no framework) | Node.js | npm ci → lint → build |
| `Cargo.toml` | Rust | cargo check → clippy → build → test |
| `go.mod` | Go | go vet → build → test |
| `pyproject.toml` / `requirements.txt` | Python | pip install → ruff → pytest |
| `composer.json` | Laravel | composer install → artisan lint → phpunit |
| _none of the above_ | Unknown | Generic (configurable) |

## Workflow Styles

### GitFlow

```
master ── production
  └─ develop ── integration
       ├─ feat/xxx
       ├─ fix/xxx
       └─ ...
```

Feature branches merge to `develop` via terminal. `develop` merges to `master` via PR at sprint end.

**Best for:** teams with scheduled releases, multiple environments.

### GitHub Flow

```
master ── production
  └─ feat/xxx ──(PR)──→ master
```

Everything branches from `master`. All merges go through PR. Deploy immediately after merge.

**Best for:** continuous deployment, small teams.

### Trunk-Based

```
master ── only long-lived branch
  └─ <short> ──(merge)──→ master
```

Branches live less than 1 day. Feature flags for incomplete work. Pair programming encouraged.

**Best for:** senior teams, high-throughput CI, microservices.

## Examples

### Next.js project

```bash
cd my-next-app
flowkit init
# Detects Next.js, generates Next.js CI, uses Husky hooks
```

### Go microservice

```bash
cd my-service
go mod init github.com/user/my-service
flowkit init
# Detects Go, generates Go CI, uses .githooks/ shell hooks
```

### Monorepo with mixed stacks

Run `flowkit init` in the root directory. If stack detection finds multiple candidates, it prioritizes JS/TS and prints a note.

### Re-running with saved config

```bash
flowkit init --force   # re-generates all files from flowkit.json
```

## Contributing

1. Fork the repo
2. Create a feature branch from `develop`: `git checkout -b feat/my-feature`
3. Commit with conventional commits
4. Push and create a PR to `develop`

### Development

```bash
git clone https://github.com/DycandX/flowkit.git
cd flowkit
go run . init              # test the CLI
go test ./internal/...      # run tests
go vet ./...                # static analysis
```

### Adding a template

1. Create the file under `internal/generator/templates/`
2. Register it in `internal/generator/generator.go` in the `filesFor()` function
3. Run `go test ./internal/generator/ -v` to verify
4. Commit with message: `feat(templates): add <description>`

## Acknowledgments

Inspired by the need for consistent, enforceable git workflows across projects of all stacks. Built with Go, cobra, and promptui.

## License

MIT — see [LICENSE](LICENSE).

---

<p align="center">
  <sub>Made with ❤️ by <a href="https://github.com/DycandX">DycandX</a></sub>
</p>
