# 🧮 Simple Calculator

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Type-CLI_Tool-181717?style=for-the-badge&logo=windows-terminal&logoColor=white" alt="CLI Tool">
  <img src="https://img.shields.io/badge/Dependencies-Standard_Library-blue?style=for-the-badge" alt="Standard Library">
</p>

<p align="center">
  <b>A simple command-line calculator written in Go.</b>
  <br>
  Perform basic arithmetic operations and exponentiation directly from your terminal.
</p>

---

## 📌 About

**Simple Calculator** is a small command-line application written in **Go (Golang)** as part of my journey into learning the language.

The program takes two numbers as input and performs five mathematical operations: addition, subtraction, multiplication, division, and exponentiation.

This project was created to explore Go's syntax, functions, basic arithmetic, user input, and standard library while gaining practical programming experience.

## ✨ Features

- ➕ **Addition** — calculates the sum of two numbers.
- ➖ **Subtraction** — calculates the difference between two numbers.
- ✖️ **Multiplication** — calculates the product of two numbers.
- ➗ **Division** — calculates the quotient of two numbers.
- 🔢 **Exponentiation** — raises the first number to the power of the second.
- 🔣 **Floating-Point Support** — uses `float64` to work with decimal numbers.
- 🧩 **Separate Functions** — implements each arithmetic operation in its own function.
- 📦 **No External Dependencies** — uses Go's standard library.

## 🧮 Supported Operations

| Operation | Operator | Example |
|:--|:--:|--:|
| Addition | `+` | `5 + 3 = 8` |
| Subtraction | `-` | `5 - 3 = 2` |
| Multiplication | `*` | `5 * 3 = 15` |
| Division | `/` | `5 / 3 ≈ 1.67` |
| Exponentiation | `^` | `5 ^ 3 = 125` |

## 🛠️ Built With

- **[Go](https://go.dev/)** — programming language.
- **`fmt`** — handles console input and output.
- **`math`** — provides mathematical functions such as `math.Pow()`.

## 🚀 Getting Started

### Prerequisites

Make sure [Go](https://go.dev/dl/) is installed on your computer.

### Installation

**1. Clone the repository**

```bash
git clone https://github.com/w94e/Simple-Calculator.git
```

**2. Navigate to the project directory**

```bash
cd Simple-Calculator
```

**3. Run the application**

```bash
go run main.go
```

> If the source file has a different name, replace `main.go` with the actual filename.

## 💻 Example

```text
《 Простой калькулятор на Go 》

✦› Введите первое число: 5
✦› Введите второе число: 3

┌✦› Сумма чисел: 5 + 3 = 8
├✦› Разность чисел: 5 - 3 = 2
├✦› Произведение чисел: 5 * 3 = 15
├✦› Частное чисел: 5 / 3 = 1.6666666666666667
└✦› Возведение в степень: 5 ^ 3 = 125
```

## 📂 Project Structure

```text
Simple-Calculator/
└── main.go    # Main application and calculator functions
```

## 🎯 Learning Goals

This project serves as a practical exercise in several fundamental Go concepts:

- Variables and data types (`float64`)
- Functions and return values
- Console input and output
- Arithmetic operators
- Floating-point arithmetic
- The `math` standard library package
- Basic input validation
- Code organization and function decomposition

## ⚠️ Project Disclaimer

> This project is strictly educational and was created with limited knowledge of the Go programming language.
>
> The code may contain significant bugs, design flaws, and architectural shortcomings. It was developed as an experiment to explore Go's syntax, complexity, capabilities, and potential for further learning and development.
>
> **This project is not intended for production use.** The code is provided for educational and experimental purposes only and should not be considered production-ready.
>
> The project may be refactored, improved, or rewritten as my understanding of Go evolves.
>
> *Learn by building. Experiment by coding. Improve by making mistakes.*

## 👨‍💻 Author

**[w94e](https://github.com/w94e)**

- GitHub Profile: [@w94e](https://github.com/w94e)
- Repository: [Simple-Calculator](https://github.com/w94e/Simple-Calculator)

---

<p align="center">
  ⭐ If you find this project interesting, consider giving it a star!
  <br><br>
  <i>Built with Go 💙</i>
</p>
