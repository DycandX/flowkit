# TASK-006: Banner + --help Polish

**Estimasi:** 20 menit

## Goal

Banner ASCII lebih keren, `--help` output lebih informatif dengan contoh usage.

## Files

- `cmd/root.go` (modify — ASCII banner)
- `cmd/init.go` (modify — help text + flag ordering)

## Todo Checklist

- [ ] Update ASCII banner di root.go
- [ ] Tambah contoh usage di init.go Long description
- [ ] Urutin flag biar rapi di --help
- [ ] Build + verify

## AI Prompt

```
Polish CLI output untuk "flowkit" (Go CLI, cobra, charmbracelet/huh).

Repo: github.com/DycandX/flowkit

## 1. Update banner di cmd/root.go

Current banner (const banner di root.go):
```
 ╔══════════════════════════════════════╗
 ║            flowkit                    ║
 ║    Workflow scaffold in 1 command     ║
 ╚══════════════════════════════════════╝
```

Ganti dengan banner dari font "ANSI Shadow" yang dihasilkan oleh:
https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=flowkit

Hasilnya kira-kira:
```
  ______ _       _     _  __
 |  ____| |     | |   | |/ /
 | |__  | | ___ | |_  | ' / ___ _   _
 |  __| | |/ _ \| __| |  < / _ \ | | |
 | |    | | (_) | |_  | . \ (_) | |_| |
 |_|    |_|\___/ \__| |_|\_\___/ \__, |
                                  __/ |
                                 |___/
```

Pastikan banner dimasukkan sebagai const string di cmd/root.go.
Gunakan raw string literal.

## 2. Update init command help di cmd/init.go

Ubah Long field jadi:

```go
Long: `Scaffold WORKFLOW.md, CI pipeline, and git hooks for your project.

Examples:
  flowkit init                        Interactive mode with prompts
  flowkit init -n "my-app"            Non-interactive, all defaults
  flowkit init --config flowkit.json  Re-run with saved config
  flowkit init -n "app" -w github-flow --hooks=false
`,
```

## 3. Urutin flag registration di init()

Flags harus urut:
1. --project-name / -n
2. --main-branch / -b
3. --language / -l
4. --workflow-style / -w
5. --ci
6. --hooks
7. --force / -f
8. --config

## 4. Test

go build -o bin\flowkit.exe .
bin\flowkit.exe --help

Output harus:
- Banner ASCII
- Examples
- Flags urut rapi

Jangan ubah file lain.
```

## DoD

- [ ] `flowkit --help` tampil banner ASCII
- [ ] `flowkit init --help` tampil examples
- [ ] Flag urut rapi (project-name duluan)
