# Truthy — Boolean Expression DSL with Truth Tables

**Truthy** is a small domain-specific language (DSL) and CLI toolset for working with Boolean logic.  
It lets you parse logical expressions, list the variables used, and generate full truth tables.

---

## ✨ Features

- **Operators** (with proper precedence & associativity):
  - `!` — NOT
  - `&` — AND
  - `|` — OR (inclusive)
- **Literals:** `true`, `false`
- **Parentheses** for grouping
- **Variable identifiers:** `A`, `B`, `x1`, `foo`, …

Tools included:
- `tt` — generate full truth table for an expression
- `vars` — list all variables used in an expression

---

## 📦 Installation

Build from source:

```powershell
# Clone the repo
git clone https://github.com/syedazeez337/truthy.git
cd truthy

# Build executables (Windows example)
go build -o .\bin\tt.exe   .\cmd\tt
go build -o .\bin\vars.exe .\cmd\vars
````

Executables will be in `./bin`.

---

## 🚀 Usage

### `tt` — truth table generator

```powershell
.\bin\tt.exe "A & !B | C"
```

Output:

```
A   B   C   expr
F   F   F   F
F   F   T   T
F   T   F   F
F   T   T   T
T   F   F   T
T   F   T   T
T   T   F   F
T   T   T   T
```

### `vars` — variable lister

```powershell
.\bin\vars.exe "A & !B | C"
```

Output:

```
OK
vars: A, B, C
```

### Stdin mode

Both tools also read from stdin if no argument is given:

```powershell
echo "(true | A) & !false" | .\bin\tt.exe
```

---

## 🔧 Grammar

ANTLR4 grammar (current):

```antlr
expr
    : '!' expr             # Not
    | expr '&' expr        # And
    | expr '|' expr        # Or
    | '(' expr ')'         # Parens
    | TRUE                 # TrueLit
    | FALSE                # FalseLit
    | IDENT                # Var
    ;

TRUE  : [Tt] [Rr] [Uu] [Ee] ;
FALSE : [Ff] [Aa] [Ll] [Ss] [Ee] ;
IDENT : [A-Za-z_] [A-Za-z_0-9]* ;
WS    : [ \t\r\n]+ -> skip ;
LINE_COMMENT : '//' ~[\r\n]* -> skip ;
```

---

## 🏗 Project Layout

```
.
├── grammar/        # ANTLR grammar (.g4)
├── parser/         # Generated Go sources (antlr4 -Dlanguage=Go …)
├── internal/engine # Shared Go code (parse, eval, truth table)
├── cmd/
│   ├── tt/         # Truth table CLI
│   └── vars/       # Variable lister CLI
└── bin/            # Compiled executables
```

---

## 📖 Background

This project is an educational DSL for Boolean logic.
It parses expressions, finds all variables, and generates complete truth tables.
Useful for learning propositional logic, compilers, or just playing with Boolean algebra.

---

## 📜 License

MIT. See [LICENSE](LICENSE).

```