# TASK-004: Enable Auto-Completion (bash/zsh/fish)

**Estimasi:** 15 menit

## Goal

User bisa generate completion script untuk shell mereka.

## Files

- `cmd/root.go` (modify)

## Todo Checklist

- [ ] Enable cobra completion di root command
- [ ] Test `flowkit completion bash` output
- [ ] Test `flowkit completion zsh` output
- [ ] Build + verify

## AI Prompt

```
Tambahkan shell completion support ke Go CLI "flowkit" di cmd/root.go.

Cobra udah built-in completion command. Tinggal enable.

Edit cmd/root.go:

1. Tambahkan properti ke rootCmd:
```go
var rootCmd = &cobra.Command{
    Use:   "flowkit",
    Short: "...",
    Long:  "...",
    SilenceUsage: true,
    Version:      Version,
    CompletionOptions: cobra.CompletionOptions{
        DisableDefaultCmd: false,
    },
}
```

2. Atau cara lain — jangan set CompletionOptions, karena cobra secara default
   udah enable completion command. Yang perlu dilakukan adalah:

   Di fungsi Execute(), pastikan completion command terdaftar:
```go
func Execute() {
    rootCmd.AddCommand(completionCmd) // TAPI ini udah otomatis
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

Cobra secara default menambahkan "completion [bash|zsh|fish|powershell]" command.
Tinggal pastikan ga di-disable.

3. Test:
   go run . completion bash    → output shell script
   go run . completion zsh     → output shell script
   go run . completion fish    → output shell script

4. Build & verify:
   go build -o /dev/null .
   go vet ./cmd/...

Jangan ubah file lain.
```

## DoD

- [ ] `go run . completion bash` output bash completion script
- [ ] `go run . completion zsh` output zsh completion script
- [ ] `flowkit --help` tampilin subcommand "completion"
