# calcdemo
Golang testing - basics 

## Variations
### Basic Commands
`go test`
`go test -v`

### Detailed analysis
```
go test -coverage
go test -coverprofile cover.out
go tool cover -func ./cover.out 
go tool cover -html ./cover.out 

```

For example
```
vinayb@carbon ~/go/src/github.com/balamuru/calcdemo/calc $ go test -coverprofile cover.out
PASS
coverage: 66.7% of statements
ok      github.com/balamuru/calcdemo/calc       0.001s
vinayb@carbon ~/go/src/github.com/balamuru/calcdemo/calc $ go tool cover -func ./cover.out 
github.com/balamuru/calcdemo/calc/calc.go:3:    Add             100.0%
github.com/balamuru/calcdemo/calc/calc.go:8:    Sub             100.0%
github.com/balamuru/calcdemo/calc/calc.go:13:   mult            0.0%
total:                                          (statements)    66.7%
```