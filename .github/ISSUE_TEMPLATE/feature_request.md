---
name: Feature Request
description: Usulkan fitur baru atau improvement
labels: ["enhancement"]
body:
  - type: markdown
    attributes:
      value: "Terima kasih sudah berkontribusi. Jelaskan ide kamu di bawah."
  - type: textarea
    id: problem
    attributes:
      label: Masalah
      description: Jelaskan masalah atau kebutuhan yang ingin dipecahkan
      placeholder: "Saya sering kesulitan ketika..."
    validations:
      required: true
  - type: textarea
    id: solution
    attributes:
      label: Solusi yang Diinginkan
      description: Gambaran solusi ideal menurut kamu
      placeholder: "Akan lebih baik jika flowkit bisa..."
    validations:
      required: true
  - type: textarea
    id: alternatives
    attributes:
      label: Alternatif yang Sudah Dicoba
      description: Apa yang sudah kamu lakukan untuk mengatasi masalah ini?
      placeholder: "Selama ini saya pakai cara manual..."
  - type: textarea
    id: example
    attributes:
      label: Contoh CLI Output (Mockup)
      description: Kalo ada gambaran command / output yang diinginkan
      render: shell
