# {{PROJECT_NAME}} — Git Workflow & Best Practices

> **Template universal.** Ganti semua `{{VAR}}` sesuai project sebelum dipakai.

---

## Branching Strategy

```
{{MAIN_BRANCH}}                production — live
  └─ develop                   integration & staging
       ├─ feat/xxx             fitur baru
       ├─ fix/xxx              bug fix
       ├─ perf/xxx             performance
       ├─ sec/xxx              security
       ├─ chore/xxx            tooling, CI, refactor
       └─ docs/xxx             dokumentasi
```

### Ground Rules

1. **Task/fix branch merge langsung ke `develop` via terminal.** No PR needed.
2. **Branch dari `develop`, merge ke `develop`.**
3. **Hanya `develop` yang merge ke `{{MAIN_BRANCH}}` — via PR di GitHub.** (setelah sprint selesai)
4. **Branch pendek umurnya.** Maksimal 2-3 hari. Kalo lebih, split jadi branch lebih kecil.
5. **Tiap logical change satu commit.** Jangan nunggu sampe selesai semua baru commit.

---

## Workflow per Sprint

```
┌──────────────────────────────────────────────────────────────┐
│                      Satu Sprint                             │
│                                                              │
│  Task 1:  feat/xxx ──(merge terminal)──→ develop             │
│  Task 2:  fix/xxx  ──(merge terminal)──→ develop             │
│  Task 3:  perf/xxx ──(merge terminal)──→ develop             │
│  ...                                                          │
│  Final:   develop ──(PR)──→ {{MAIN_BRANCH}} ──(tag)──→ deploy│
└──────────────────────────────────────────────────────────────┘
```

| Langkah | Tujuan | Kapan | Metode |
|---------|--------|-------|--------|
| Task branch → `develop` | Integrasi & testing | Tiap 1 task selesai | Merge terminal |
| `develop` → `{{MAIN_BRANCH}}` | Rilis production | Setelah 1 sprint selesai | PR di GitHub |

**Aturan penting:**
- **Jangan pernah merge langsung ke `{{MAIN_BRANCH}}`.** `develop` adalah staging/QA. `{{MAIN_BRANCH}}` cuma diisi kalo udah tested lewat `develop`.
- **Tiap 1 task = 1 branch.** Jangan gabung 2 task beda dalam 1 branch.

---

## Naming Convention

Format: `<prefix>/<kebab-case-description>`

| Prefix | Penggunaan | Contoh |
|--------|-----------|--------|
| `feat/` | Fitur baru | `feat/user-auth` |
| `fix/` | Bug fix | `fix/login-validation` |
| `perf/` | Performance | `perf/db-indexing` |
| `sec/` | Security | `sec/rls-policies` |
| `chore/` | Tooling, CI, refactor | `chore/sentry-setup` |
| `docs/` | Dokumentasi | `docs/api-guide` |

---

## Commit Message Convention

Gunakan **Conventional Commits**:

```
<type>(<scope>): <deskripsi>
```

| Type | Kapan |
|------|-------|
| `feat` | Fitur baru |
| `fix` | Bug fix |
| `perf` | Performance |
| `sec` | Security |
| `refactor` | Refactor (tanpa ubah behavior) |
| `chore` | Tooling, CI, dependency |
| `docs` | Dokumentasi |
| `style` | Formatting (bukan CSS) |

**Aturan:**
- Bahasa Inggris untuk deskripsi
- Imperative mood: "add" bukan "added"
- Maksimal 72 karakter subjek
- Satu commit = satu logical change

Contoh:
```
feat(auth): add OAuth2 login flow
fix(api): validate request body before processing
perf(db): add index to users.email column
```

---

## Task Workflow

### 1. Mulai Task Baru

```bash
git checkout develop
git pull
git checkout -b <prefix>/<deskripsi>
```

### 2. Selama Ngerjain

```bash
git add <file>
git commit -m "<type>(<scope>): <message>"

# Kalo develop udah maju, rebase
git fetch origin
git rebase origin/develop
```

#### Resolve Konflik Saat Rebase

```bash
# Cek file conflict
git status
# Edit file → hapus marker <<<<<<< ======= >>>>>>>
# Stage + continue
git add <file>
git rebase --continue
# Kalo bingung, cancel:
git rebase --abort
```

**PENTING:** `git push --force-with-lease` setelah rebase.

### 3. Merge ke Develop (Terminal)

```bash
git checkout develop
git pull
git merge <nama-branch>
git push origin develop
git branch -d <nama-branch>
git push origin --delete <nama-branch>
```

Pastikan `{{BUILD_COMMAND}} && {{LINT_COMMAND}}` lulus sebelum merge.

### 4. Release ke {{MAIN_BRANCH}} (via PR)

```bash
git checkout {{MAIN_BRANCH}}
git pull
git merge develop
git tag v<major>.<minor>.<patch>
git push origin {{MAIN_BRANCH}} --tags
```

Atau bikin PR dari `develop` ke `{{MAIN_BRANCH}}` di GitHub.

---

## Best Practices

### 1. Keep Branch Up-to-Date
```bash
git fetch origin
git rebase origin/develop
```

### 2. Jangan Rebase Shared Branch
**Jangan pernah** rebase `develop` atau `{{MAIN_BRANCH}}`. Rebase cuma di **personal branch**.

### 3. Rebase > Merge (personal branch)
Rebase bikin history linear. Merge bikin banyak merge commit.

### 4. Commit Kecil & Sering
Tiap logical change = satu commit. Gampang revert, gampang cherry-pick.

### 5. Self-Review Sebelum Push
```bash
git diff develop...HEAD
```
Pastikan: no debug code, no console.log, naming konsisten, error handling ada.

### 6. Bersihin Branch Bekas
```bash
git branch -d <branch-name>
git push origin --delete <branch-name>
```

---

## Code Quality Checklist (sebelum commit)

- [ ] `{{LINT_COMMAND}}` — no errors
- [ ] `{{BUILD_COMMAND}}` — no errors
- [ ] No `any` types (kalo bisa)
- [ ] No commented-out code
- [ ] No `console.log`
- [ ] Input validation ada di trust boundaries
- [ ] Error handling proper

---

## Project-Specific Variables

| Variable | Value |
|----------|-------|
| `{{PROJECT_NAME}}` | Isi nama project |
| `{{MAIN_BRANCH}}` | `main` / `master` |
| `{{STACK}}` | Next.js, Laravel, dll |
| `{{ROOT_PATH}}` | Path absolut project |
| `{{BASE_BRANCH}}` | `develop` |
| `{{INSTALL_COMMAND}}` | `npm install` / `composer install` / `pip install` |
| `{{DEV_COMMAND}}` | `npm run dev` / `php artisan serve` |
| `{{BUILD_COMMAND}}` | `npm run build` / `npm run build` |
| `{{LINT_COMMAND}}` | `npm run lint` / `npx eslint` |

---

*Template: SynapseCS Git Workflow. Adaptasi sesuai kebutuhan project.*
