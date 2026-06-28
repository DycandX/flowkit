# TASK-001: Demo GIF di README

**Estimasi:** 30 menit

## Goal

Bikin demo GIF yang menunjukkan `npx @dycandx/flowkit init` dari awal sampai selesai, lalu tampilin di README.

## Files

- `docs/demo.gif` (new — generated)
- `README.md` (modify — tambah section Demo)

## Todo Checklist

- [ ] Install terminalizer
- [ ] Record demo
- [ ] Render GIF
- [ ] Update README dengan preview GIF

## AI Prompt

```
Saya punya Go CLI tool "flowkit" di https://github.com/DycandX/flowkit
Tool ini generate WORKFLOW.md, CI pipelines, dan git hooks.

Buatkan demo GIF / terminal recording yang menunjukkan:

1. User jalanin `npx @dycandx/flowkit init` di folder project Next.js
2. Prompt interaktif:
   - Project name: "my-app"
   - Main branch: enter (default master)
   - Language: pilih English
   - Workflow style: pilih GitFlow
3. Output generated files (WORKFLOW.md, CI, hooks, dll)
4. Done message dengan next steps

Syarat:
- GIF disimpan ke docs/demo.gif
- Update README.md: tambah section "## Demo" sebelum "## Table of Contents"
  dengan preview GIF: ![flowkit demo](docs/demo.gif)
- Background terminal gelap, font monospace
- Durasi maks 30 detik
- Fokus ke flow interaktif, jangan terlalu cepet

Untuk bikin GIF:
1. Install terminalizer: npm install -g terminalizer
2. Init config: terminalizer init
3. Rekam: terminalizer record demo
4. Render: terminalizer render demo --width 800 --height 400
5. Output: docs/demo.gif
```

## DoD

- [ ] `docs/demo.gif` ada, ukuran < 2MB
- [ ] README section "## Demo" dengan preview GIF
- [ ] GIF durasi < 30 detik
- [ ] Isi GIF jelas terbaca
