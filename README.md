# Truth Table DSL

**truth-table-dsl** is a small domain-specific language (DSL) and CLI toolset for working with Boolean logic.  
It parses logical expressions, lists the variables used, and generates full truth tables.

Built with [ANTLR 4](https://www.antlr.org/) (Go target) and Go.

---

## ✨ Features

- **Operators** (with precedence & associativity):
  - `!` — NOT (highest precedence, unary)
  - `&` — AND
  - `|` — OR (inclusive)
- **Literals:** `true`, `false`
- **Parentheses** for grouping
- **Identifiers:** variables like `A`, `B`, `x1`, `foo`

### Tools
- **`tt`** — generate a full truth table for an expression
- **`vars`** — list all variables used in an expression

---

## 📦 Build

Clone and build from source:

```powershell
git clone https://github.com/syedazeez337/truth-table-dsl.git
cd truth-table-dsl

# build executables
go build -o .\bin\tt.exe   .\cmd\tt
go build -o .\bin\vars.exe .\cmd\vars
````

Executables will appear in `./bin`.

---

## 🚀 Usage

### Generate a truth table

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

### List variables

```powershell
.\bin\vars.exe "A & !B | C"
```

Output:

```
OK
vars: A, B, C
```

### Stdin mode

Both tools also read from stdin:

```powershell
echo "(true | A) & !false" | .\bin\tt.exe
```

---

## 🔧 Grammar

ANTLR4 grammar (current):

```antlr
grammar TruthExpr;

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
├── parser/         # Generated Go parser (antlr4 -Dlanguage=Go …)
├── internal/engine # Shared Go code (parse, eval, truth table)
├── cmd/
│   ├── tt/         # Truth table CLI
│   └── vars/       # Variable lister CLI
└── bin/            # Built executables
```

---

## 📖 Background

This project is an educational DSL for Boolean logic.
It demonstrates how to:

* Write an ANTLR grammar
* Generate a Go parser
* Evaluate Boolean expressions
* Enumerate variables and build truth tables

Useful for learning propositional logic, ANTLR, and compiler-style DSLs in Go.

---

## 📜 License

MIT. See [LICENSE](LICENSE).

```