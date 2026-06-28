# Contributing to flowkit

Thanks for your interest. Keep it simple.

## 1. Branching

- Branch from `develop`, merge to `develop`
- Naming: `<prefix>/<kebab-case>`

  | Prefix | When |
  |--------|------|
  | `feat/` | New feature |
  | `fix/` | Bug fix |
  | `perf/` | Performance |
  | `sec/` | Security |
  | `chore/` | Tooling, CI, deps |
  | `docs/` | Documentation |

## 2. Commits

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>
```

Examples:

```
feat(cli): add --dry-run flag
fix(generator): escape template delimiters correctly
docs(readme): add quick-start section
```

English. Imperative mood. Max 72 chars.

## 3. Development

```bash
# Test
go test ./internal/... -count=1

# Lint
go vet ./...

# Build
go build -o /dev/null .
```

## 4. PR

- One branch = one logical change
- Rebase before push
- No merge commits — use squash merge when merging to `develop`
