# Demo Terminal Output

Untuk generate GIF dari output ini, gunakan terminalizer / vhs / asciinema.

## Non-Interactive Mode

```text
$ mkdir my-app && cd my-app
$ git init
Initialized empty Git repository in /my-app/.git/
$ echo '{"dependencies":{"next":"14"}}' > package.json

$ npx @dycandx/flowkit init -n "my-app" -b master -l en -w gitflow

→ Detected stack: next


Generating workflow files...
✓ flowkit.json (config saved)
✓ WORKFLOW.md
✓ .github/workflows/ci.yml
✓ .github/workflows/pr-check.yml
✓ .husky/pre-commit
✓ .husky/commit-msg
✓ commitlint.config.js
✓ .lintstagedrc.json

✅ Done! Next steps:
  1. Review generated files
  2. git add .
  3. git commit -m "chore: add workflow scaffold"

  To activate git hooks:
    npm install         (activates Husky)

$ ls -la
.rw-r--r--    WORKFLOW.md
.rw-r--r--    flowkit.json
.rw-r--r--    commitlint.config.js
.rw-r--r--    .lintstagedrc.json
drwxr-xr-x    .github/
drwxr-xr-x    .husky/
```

## Interactive Mode (huh prompts)

```text
$ npx @dycandx/flowkit init

? Project name:  my-app
? Main branch:   master (default)
? Language:      English (selected)
? Workflow style: GitFlow (selected)

→ Detected stack: next

Generating workflow files...
✓ WORKFLOW.md
✓ .github/workflows/ci.yml
✓ .github/workflows/pr-check.yml
✓ .husky/pre-commit

✅ Done! Next steps:
  ...
```

## Cara Generate GIF

### Option 1: terminalizer (npm)

```bash
npm install -g terminalizer
terminalizer record demo    # record live session
terminalizer render demo    # output: my-app.gif
```

### Option 2: vhs (Go)

```bash
go install github.com/charmbracelet/vhs@latest
vhs < demo.tape
```

### Option 3: asciinema + svg-term-cli

```bash
asciinema rec demo.cast
asciinema play demo.cast
npx svg-term-cli --in demo.cast --out demo.gif
```
