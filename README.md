# Learning Go language

## Fundamentals

- Executabel programms starts by main package.
- Code in Go is organized in packages. To use them you need to declare in one or many **imports**.
- The start point of a Go program is the function **main**.
- Inside the same package you can’t have two functions with the same name, in our case **main**.
  
### Commands

- `go help <command>` show docs from an especific command of GO.
- `godoc --http=:6060` open offline docs of Go on port *6060*.
- `go vet <filepath>` “Vet examines Go source code and reports suspicious constructs, such as
Printf calls whose arguments do not align with the format string.”
- `go get -u <package-name>` install package to your $GOPATH; ex: `go get -u github.com/go-sql-driver/mysql` install mysql driver.

### Constants and Variables

- Like *Javascript* we have `const` and `var`. But we can use `:=` to create and assing value to a variable that not exists yet.
- Every time you create a new variable, it need to be used in some place. Unlike other languages, Go doesn't accept unused variables and breaks when compiling.

### Console Prints

- `fmt.Print()` print in console without `\n` in the end of the string.
- `fmt.Println()` print in console with `\n` in the end of the string.
- `fmt.Printf()` print in console without `\n` in the end of the string and add the possibility of adding variables to the string.
- `%d` for integer; `%f` for float, `%t` for boolean, `%s` for string. `%v` for general use.

### Types

Some types of Go is:

- `int`
- `float64`
- `uint8`
- `bool`
- `string`

### Functions

Functions are pieces of code with one responsability that execute one part of the code. In Go, functions are most like in other C-like languages, like Typescript.

### Pointers

Pointers are memory references to a variable. In Go, pointers are most like in other C-like languages.
