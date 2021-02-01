# Learning Go

## Fundamentals

- Executabel programms starts by main package.
- Code in Go is organized in packages. To use them you need to declare in one or many **import**.
- The start point of a Go program is the function **main**.
- Inside the same package you can't have two functions with the same name, in our case **main**.
  
### Commands

- `go help <command>` show docs from an especific command of GO.
- `godoc --http=:6060` open offline docs of Go on port *6060*.
- `go vet <filepath>` "Vet examines Go source code and reports suspicious constructs, such as
Printf calls whose arguments do not align with the format string."
- `go get -u <package-name>` install package to your $GOPATH; ex: `go get -u github.com/go-sql-driver/mysql` install mysql driver.
