# Changelog

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
