# go-code-coverage

Trying out test coverage tools in Go

### run tests with code coverage

```sh
go test ./... -race -covermode=atomic -coverprofile coverage
```

### Print out individual method coverage

```sh
go tool cover -func=coverage
```

### Visualize code coverage with an HTML report

```sh
go tool cover -html=coverage
```
