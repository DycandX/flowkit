# TASK-003: E2E Test (Real Init)

**Estimasi:** 1 jam

## Goal

Test `flowkit init` beneran jalan di temp directory + verifikasi semua file tergenerate.

## Files

- `internal/generator/e2e_test.go` (new)
- `go.mod`, `go.sum` (updated — jika perlu dependensi baru)

## Todo Checklist

- [ ] Buat `e2e_test.go` di package `generator`
- [ ] Test init Next.js (via `package.json` mock) → semua file JS tergenerate
- [ ] Test init Go project → `.githooks/` bukan `.husky/`
- [ ] Test skip existing files (force=false)
- [ ] Test force overwrite (force=true)
- [ ] `go test ./internal/generator/ -v -count=1` PASS

## AI Prompt

```
Buat E2E test untuk Go CLI "flowkit" di internal/generator/e2e_test.go.

Package: generator (sama dengan existing file generator_test.go)

Test harus:

## Test 1: Init Next.js project
1. Buat temp directory via t.TempDir()
2. Simpan working directory asli: origDir, _ := os.Getwd()
3. Pindah ke temp dir: os.Chdir(tmpDir)
4. Defer: os.Chdir(origDir)
5. Init git repo: exec.Command("git", "init").Run()
6. Buat package.json: os.WriteFile("package.json", []byte(`{"dependencies":{"next":"14"}}`), 0644)
7. Deteksi stack: stack := string(detector.Detect()) — harus return "next"
8. Buat config: cfg := &config.Config{ProjectName: "test-app", Stack: "next", ...}
9. Generate: generator.GenerateAll(cfg, true)
10. Verifikasi file ADA:
    - WORKFLOW.md
    - .github/workflows/ci.yml
    - .github/workflows/pr-check.yml
    - .husky/pre-commit
    - .husky/commit-msg
    - commitlint.config.js
    - .lintstagedrc.json
11. Verifikasi isi flowkit.json valid JSON

## Test 2: Init Go project (non-JS)
1. Temp dir baru
2. os.Chdir
3. Buat go.mod: os.WriteFile("go.mod", []byte("module test"), 0644)
4. Git init
5. Deteksi stack — harus "go"
6. Generate with force=true
7. Verifikasi:
    - .githooks/pre-commit ADA (bukan .husky/)
    - .githooks/commit-msg ADA
    - .husky/ TIDAK ADA
    - commitlint.config.js TIDAK ADA

## Test 3: Skip existing files
1. Temp dir + git init + package.json
2. Generate pertama (force=false) — semua file terbuat
3. Generate kedua (force=false) — harus skip semua file
   Output: "⏭  WORKFLOW.md already exists, skipping" etc.

## Test 4: Force overwrite
1. Temp dir + git init + package.json
2. Generate pertama
3. Hapus WORKFLOW.md
4. Generate kedua dengan force=true
5. WORKFLOW.md harus terbuat lagi

Import:
- github.com/DycandX/flowkit/internal/detector
- github.com/DycandX/flowkit/internal/generator
- github.com/DycandX/flowkit/internal/config
- testing, os, os/exec, path/filepath, encoding/json

Jalankan dengan: go test ./internal/generator/ -v -count=1 -run TestE2E
```

## DoD

- [ ] Semua test PASS
- [ ] Test Next.js → file `.husky/` tergenerate
- [ ] Test Go → file `.githooks/` tergenerate, file JS-specific tidak ada
- [ ] Skip existing files works (tidak overwrite tanpa --force)
