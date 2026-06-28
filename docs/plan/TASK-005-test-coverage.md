# TASK-005: Test Coverage > 80%

**Estimasi:** 2 jam

## Goal

`go test -cover ./...` > 80%.

## Files

- `internal/detector/detector_test.go` (new)
- `internal/config/config_test.go` (new)
- `internal/generator/generator_test.go` (update — tambah test case)

## Todo Checklist

- [ ] Test detector untuk semua stack (next, react, vue, nuxt, node, go, rust, python, laravel, unknown)
- [ ] Test detector untuk file tidak terbaca / corrupt
- [ ] Test config JSON marshal/unmarshal
- [ ] Test saveConfig function
- [ ] Test force overwrite vs skip
- [ ] Test non-JS stack (.githooks/ path)
- [ ] `go test -cover ./...` > 80%

## AI Prompt

```
Tambah test coverage untuk Go package di internal/detector dan internal/config,
dan update existing test di internal/generator.

Target: go test -cover ./... > 80%

## 1. internal/detector/detector_test.go

Package detector_test (bisa pakai package detector — whitebox)

Test Detect() function:

func TestDetectNextJS(t *testing.T) {
    tmp := t.TempDir()
    orig, _ := os.Getwd()
    os.Chdir(tmp)
    defer os.Chdir(orig)
    exec.Command("git", "init").Run()
    os.WriteFile("package.json", []byte(`{"dependencies":{"next":"14"}}`), 0644)
    if got := Detect(); got != StackNext {
        t.Errorf("Detect() = %v, want %v", got, StackNext)
    }
}

Buat test yang sama untuk:
- react (package.json dengan "react")
- vue (package.json dengan "vue")
- nuxt (package.json dengan "nuxt")
- node (package.json tanpa framework)
- go (go.mod exist)
- rust (Cargo.toml exist)
- python (pyproject.toml exist)
- laravel (composer.json exist)
- unknown (dir kosong)

Gunakan t.TempDir() + os.Chdir() + git init untuk tiap test case.
Jalankan subtest parallel: t.Run("next", func(t *testing.T) { ... })

## 2. internal/config/config_test.go

Package config

Test JSON marshal/unmarshal:
```go
func TestConfigMarshal(t *testing.T) {
    cfg := &Config{
        ProjectName: "test",
        MainBranch:  "main",
        Stack:       "go",
    }
    data, err := json.Marshal(cfg)
    if err != nil { t.Fatal(err) }
    var cfg2 Config
    if err := json.Unmarshal(data, &cfg2); err != nil { t.Fatal(err) }
    if cfg2.ProjectName != cfg.ProjectName {
        t.Errorf("ProjectName = %q, want %q", cfg2.ProjectName, cfg.ProjectName)
    }
}
```

Test FeaturesConfig default values.

## 3. Update internal/generator/generator_test.go

Tambah test:
- TestSaveConfig: verify flowkit.json tertulis dengan project_name yang benar
- TestForceOverwrite: generate 2x dengan force=true → file ke-2 nge-overwrite
- TestSkipExisting: generate 2x dengan force=false → file ke-2 di-skip
- TestNonJSStack: stack "go" → .githooks/ bukan .husky/

Jalankan: go test -cover -count=1 ./internal/...
```

## DoD

- [ ] `go test -cover ./...` > 80%
- [ ] Semua test PASS
- [ ] Detector tested untuk 10 stack case
