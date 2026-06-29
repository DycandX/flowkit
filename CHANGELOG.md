# Changelog

## v0.2.1 (2026-06-29)

### Fixed

- Dynamic version in CLI banner (remove hardcoded v0.1.3)
- Banner alignment with version text

## v0.2.0 (2026-06-29)

### Added

- Demo GIF, Issue Templates, E2E Tests
- Shell auto-completion (bash/zsh/fish/powershell)
- GitHub Discussions enabled
- Test coverage > 80% (detector 91.9%, generator 82.3%)

### Fixed

- Issue templates format (YAML → Markdown)
- Banner polish + init --help examples

## v0.1.3 (2026-06-28)

### Changed

- promptui → charmbracelet/huh (no bell sound, better UX)
- CI matrix: test on Windows + macOS

### Added

- `--version` flag
- `--config` flag (re-read flowkit.json)
- Project name sanitization (escape special chars)

## v0.1.2 (2026-06-28)

### Fixed

- run.js HTTP 302 redirect handling for binary download

## v0.1.1 (2026-06-28)

### Fixed

- Goreleaser name_template to match npm binary downloader
- Scoped npm package `@dycandx/flowkit` (was blocked by `flow-kit` name conflict)
- npm publish with `--access=public` for scoped packages

## v0.1.0 (2026-06-28)

Initial release.

### Added

- `flowkit init` interactive command — cobra CLI, survey prompts
- Stack auto-detection: next, react, vue, nuxt, node, go, rust, python, laravel, unknown
- 3 workflow styles: GitFlow, GitHub Flow, Trunk-Based
- CI templates per stack (GitHub Actions)
- PR branch/commit convention enforcement (pr-check.yml)
- Pre-commit hooks (Husky for JS, .githooks/ for non-JS)
- Commit-msg hook with conventional commit validation
- commitlint.config.js + .lintstagedrc.json for JS projects
- flowkit.json config save for re-run
- Non-interactive flags: --project-name, --main-branch, --language, --workflow-style, --ci, --hooks
- MIT License
- GoReleaser multi-platform builds
- GitHub Actions CI + Release workflows
- npm distribution wrapper
