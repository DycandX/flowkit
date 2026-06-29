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
      description: Output dari `flowkit --version` atau `npx @dycandx/flowkit --version`
      placeholder: "v0.1.3"
    validations:
      required: true
  - type: dropdown
    id: os
    attributes:
      label: OS / Environment
      options:
        - Windows
        - macOS
        - Linux
        - GitHub Actions CI
        - Lainnya
    validations:
      required: true
  - type: dropdown
    id: install
    attributes:
      label: Cara Install
      options:
        - npx @dycandx/flowkit
        - npm install -g @dycandx/flowkit
        - go install
        - Binary download
        - Build dari source
    validations:
      required: true
  - type: textarea
    id: steps
    attributes:
      label: Langkah Reproduksi
      description: Step by step untuk reproduksi bug
      placeholder: |-
        1. cd /my-project && flowkit init
        2. Pilih "my-app" sebagai project name
        3. Error muncul...
      value: |-
        1.
        2.
        3.
    validations:
      required: true
  - type: textarea
    id: expected
    attributes:
      label: Yang Diharapkan
      description: Apa yang seharusnya terjadi
    validations:
      required: true
  - type: textarea
    id: actual
    attributes:
      label: Yang Terjadi
      description: Apa yang sebenarnya terjadi (paste error message)
    validations:
      required: true
  - type: textarea
    id: logs
    attributes:
      label: Terminal Output
      description: Paste full terminal output di sini
      render: shell
  - type: textarea
    id: screenshots
    attributes:
      label: Screenshot (optional)
      description: Upload screenshot jika perlu
