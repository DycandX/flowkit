# Sprint v1.0 — Production Polish

> **Goal:** Flowkit layak dipake orang lain secara serius.
> **Durasi:** ~4 jam (7 task, bisa dikerjakan paralel)
> **Base branch:** `develop`

## Urutan Eksekusi

| # | Task | Estimasi | Dikerjakan |
|---|------|----------|------------|
| 1 | TASK-007: Enable GitHub Discussions | 5 menit | Manual di UI GitHub |
| 2 | TASK-002: Issue Templates | 15 menit | AI Agent |
| 3 | TASK-004: Auto-Completion | 15 menit | AI Agent |
| 4 | TASK-006: Banner + --help Polish | 20 menit | AI Agent |
| 5 | TASK-003: E2E Test | 1 jam | AI Agent |
| 6 | TASK-005: Test Coverage > 80% | 2 jam | AI Agent |
| 7 | TASK-001: Demo GIF | 30 menit | AI Agent |

## Cara Pake

1. Buka file TASK sesuai urutan
2. Copy bagian **AI Prompt** → kirim ke AI agent (Claude/ChatGPT/Copilot)
3. Verifikasi hasil sesuai **DoD**
4. Commit ke branch task → merge ke `develop`

## Branch Convention

```
feat/sprint-v1.0/task-001-demo-gif
feat/sprint-v1.0/task-002-issue-templates
feat/sprint-v1.0/task-003-e2e-test
...
```

Atau kalo mau 1 branch aja:

```
feat/sprint-v1.0
```

## Definition of Done (Global)

- [ ] `npx @dycandx/flowkit init` works tanpa error di 3 OS
- [ ] CI hijau (ubuntu, windows, macos)
- [ ] `go test ./... -count=1` PASS
- [ ] `go test -cover ./...` > 80%
- [ ] Semua task punya prompt yang bisa di-copas
