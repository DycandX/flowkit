# TASK-002: Issue Templates (Bug Report + Feature Request)

**Estimasi:** 15 menit

## Goal

User bisa lapor bug / minta fitur lewat template standar di GitHub Issues.

## Files

- `.github/ISSUE_TEMPLATE/bug_report.md` (new)
- `.github/ISSUE_TEMPLATE/feature_request.md` (new)

## Todo Checklist

- [ ] Buat folder `.github/ISSUE_TEMPLATE/`
- [ ] Buat `bug_report.md`
- [ ] Buat `feature_request.md`
- [ ] Push + verifikasi muncul di tab Issues

## AI Prompt

```
Buatkan GitHub Issue Templates untuk repo Go CLI "flowkit" (github.com/DycandX/flowkit).

flowkit adalah CLI tool yang generate WORKFLOW.md, CI pipelines, dan git hooks.
Stack: Go (cobra CLI, charmbracelet/huh prompts)
Distribusi: npm (@dycandx/flowkit) + binary dari GitHub Releases
Platform: Windows, macOS, Linux

Buat 2 file di .github/ISSUE_TEMPLATE/:

## 1. bug_report.md

YAML frontmatter:
--- 
name: Bug Report
description: Laporkan bug atau masalah pada flowkit
labels: ["bug"]
body:
  - type: markdown
    attributes:
      value: "Terima kasih sudah melaporkan bug. Isi form di bawah."
  - type: input
    id: version
    attributes:
      label: Versi flowkit
      description: Output dari flowkit --version
      placeholder: "v0.1.3"
    validations:
      required: true
  - type: dropdown
    id: os
    attributes:
      label: OS
      options: [Windows, macOS, Linux]
    validations:
      required: true
  - type: textarea
    id: steps
    attributes:
      label: Langkah Reproduksi
      description: Step by step untuk reproduksi bug
      placeholder: |-
        1. npx @dycandx/flowkit init
        2. Pilih project name
        3. ...
    validations:
      required: true
  - type: textarea
    id: expected
    attributes:
      label: Expected Behavior
  - type: textarea
    id: actual
    attributes:
      label: Actual Behavior
  - type: textarea
    id: logs
    attributes:
      label: Terminal Output
      render: shell
  - type: textarea
    id: screenshots
    attributes:
      label: Screenshot (optional)
      description: Upload screenshot jika perlu

## 2. feature_request.md

YAML frontmatter:
---
name: Feature Request
description: Usulkan fitur baru atau improvement
labels: ["enhancement"]
body:
  - type: textarea
    id: problem
    attributes:
      label: Masalah
      description: Jelaskan masalah yang ingin dipecahkan
    validations:
      required: true
  - type: textarea
    id: solution
    attributes:
      label: Solusi yang Diinginkan
    validations:
      required: true
  - type: textarea
    id: alternatives
    attributes:
      label: Alternatif yang Sudah Dicoba
  - type: textarea
    id: example
    attributes:
      label: Contoh CLI Output (Mockup)

Gunakan Bahasa Indonesia.
Format markdown.
```

## DoD

- [ ] File `.github/ISSUE_TEMPLATE/bug_report.md` ada
- [ ] File `.github/ISSUE_TEMPLATE/feature_request.md` ada
- [ ] Buka tab Issues → New Issue → template muncul
