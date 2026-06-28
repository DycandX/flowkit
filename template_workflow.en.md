# {{PROJECT_NAME}} — Git Workflow & Best Practices

> **Universal template.** Replace all `{{VAR}}` with your project's values.

---

## Branching Strategy

```
{{MAIN_BRANCH}}                production — live
  └─ develop                   integration & staging
       ├─ feat/xxx             new feature
       ├─ fix/xxx              bug fix
       ├─ perf/xxx             performance
       ├─ sec/xxx              security
       ├─ chore/xxx            tooling, CI, refactor
       └─ docs/xxx             documentation
```

### Ground Rules

1. **Task/fix branches merge directly to `develop` via terminal.** No PR needed.
2. **Branch from `develop`, merge into `develop`.**
3. **Only `develop` merges into `{{MAIN_BRANCH}}` — via PR on GitHub.** (after sprint complete)
4. **Short-lived branches.** Max 2-3 days. Split into smaller branches if needed.
5. **One commit per logical change.** Don't wait until everything is done to commit.

---

## Sprint Workflow

```
┌──────────────────────────────────────────────────────────────┐
│                      One Sprint                              │
│                                                              │
│  Task 1:  feat/xxx ──(terminal merge)──→ develop             │
│  Task 2:  fix/xxx  ──(terminal merge)──→ develop             │
│  Task 3:  perf/xxx ──(terminal merge)──→ develop             │
│  ...                                                          │
│  Final:   develop ──(PR)──→ {{MAIN_BRANCH}} ──(tag)──→ deploy│
└──────────────────────────────────────────────────────────────┘
```

| Step | Target | When | Method |
|------|--------|------|--------|
| Task branch → `develop` | Integration & testing | Each task completed | Terminal merge |
| `develop` → `{{MAIN_BRANCH}}` | Production release | Each sprint completed | PR on GitHub |

**Important rules:**
- **Never merge directly to `{{MAIN_BRANCH}}`.** `develop` is staging/QA. `{{MAIN_BRANCH}}` only receives tested code via `develop`.
- **1 task = 1 branch.** Don't combine 2 different tasks in 1 branch.

---

## Naming Convention

Format: `<prefix>/<kebab-case-description>`

| Prefix | Usage | Example |
|--------|-------|---------|
| `feat/` | New feature | `feat/user-auth` |
| `fix/` | Bug fix | `fix/login-validation` |
| `perf/` | Performance | `perf/db-indexing` |
| `sec/` | Security | `sec/rls-policies` |
| `chore/` | Tooling, CI, refactor | `chore/sentry-setup` |
| `docs/` | Documentation | `docs/api-guide` |

---

## Commit Message Convention

Use **Conventional Commits**:

```
<type>(<scope>): <description>
```

| Type | When |
|------|------|
| `feat` | New feature |
| `fix` | Bug fix |
| `perf` | Performance |
| `sec` | Security |
| `refactor` | Refactor (no behavior change) |
| `chore` | Tooling, CI, dependencies |
| `docs` | Documentation |
| `style` | Formatting (not CSS) |

**Rules:**
- Use English
- Imperative mood: "add" not "added"
- Max 72 characters for subject
- One commit = one logical change

Examples:
```
feat(auth): add OAuth2 login flow
fix(api): validate request body before processing
perf(db): add index to users.email column
```

---

## Task Workflow

### 1. Start a New Task

```bash
git checkout develop
git pull
git checkout -b <prefix>/<description>
```

### 2. During Development

```bash
git add <file>
git commit -m "<type>(<scope>): <message>"

# If develop has moved forward, rebase
git fetch origin
git rebase origin/develop
```

#### Resolve Rebase Conflicts

```bash
# Check conflicted files
git status
# Edit file → remove markers <<<<<<< ======= >>>>>>>
# Stage + continue
git add <file>
git rebase --continue
# If stuck, cancel:
git rebase --abort
```

**IMPORTANT:** `git push --force-with-lease` after rebase.

### 3. Merge to Develop (Terminal)

```bash
git checkout develop
git pull
git merge <branch-name>
git push origin develop
git branch -d <branch-name>
git push origin --delete <branch-name>
```

Make sure `{{BUILD_COMMAND}} && {{LINT_COMMAND}}` passes before merging.

### 4. Release to {{MAIN_BRANCH}} (via PR)

```bash
git checkout {{MAIN_BRANCH}}
git pull
git merge develop
git tag v<major>.<minor>.<patch>
git push origin {{MAIN_BRANCH}} --tags
```

Or create a PR from `develop` to `{{MAIN_BRANCH}}` on GitHub.

---

## Best Practices

### 1. Keep Branch Up-to-Date
```bash
git fetch origin
git rebase origin/develop
```

### 2. Never Rebase a Shared Branch
**Never** rebase `develop` or `{{MAIN_BRANCH}}`. Rebase is for **personal branches only**.

### 3. Rebase > Merge (personal branch)
Rebase keeps history linear. Merge creates unnecessary merge commits.

### 4. Commit Small & Often
Each logical change = one commit. Easier to revert and cherry-pick.

### 5. Self-Review Before Pushing
```bash
git diff develop...HEAD
```
Make sure: no debug code, no console.log, consistent naming, proper error handling.

### 6. Clean Up Merged Branches
```bash
git branch -d <branch-name>
git push origin --delete <branch-name>
```

---

## Code Quality Checklist (before commit)

- [ ] `{{LINT_COMMAND}}` — no errors
- [ ] `{{BUILD_COMMAND}}` — no errors
- [ ] No `any` types (when possible)
- [ ] No commented-out code
- [ ] No `console.log`
- [ ] Input validation at trust boundaries
- [ ] Proper error handling

---

## Project-Specific Variables

| Variable | Value |
|----------|-------|
| `{{PROJECT_NAME}}` | Fill with project name |
| `{{MAIN_BRANCH}}` | `main` or `master` |
| `{{STACK}}` | Next.js, Laravel, etc. |
| `{{ROOT_PATH}}` | Absolute project path |
| `{{BASE_BRANCH}}` | `develop` |
| `{{INSTALL_COMMAND}}` | `npm install` / `composer install` / `pip install` |
| `{{DEV_COMMAND}}` | `npm run dev` / `php artisan serve` |
| `{{BUILD_COMMAND}}` | `npm run build` |
| `{{LINT_COMMAND}}` | `npm run lint` / `npx eslint` |

---

*Template: SynapseCS Git Workflow. Adapt to your project's needs.*
