go env -w CGO_ENABLED=1
go run .\cmd\main.go
go run .\cmd
go mod tidy

# tools
[link](https://github.com/avelino/awesome-go)


# Sqlite
[link](https://jmeubank.github.io/tdm-gcc/download/)

# todo

# db relations

# structured logging. [^1]

# test

# mux

# middlewares

# authentication

# opentelemetry

# Dockerfile

# Wrapper (HttpClient) for API call

go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
go get go.opentelemetry.io/otel/sdk/resource
go get go.opentelemetry.io/otel/trace

# Validation

currently we have two validation, first, we can validate url params (query and path) using the gorilla mux package (using regex)
we can force a parameter to be of a certain type, for example, we can force an id parameter to be an integer.
`go /todo/{id:[0-9]+}`
Second, we can validate the request body using the go-playground validator package.
a sample is added to the todos-controller (POST)

# generic service vs specific service
```go
//todoService := service.NewService[models.Todo]()
todoService := service.NewTodoService()
//todo := models.Todo{Title: vars["category"], Description: "description"}
t := todoService.Create(model)
controllers.Ok(&w, t)
```

# translation
```go
template := map[string]interface{}{
		"Name": "Shaahin",
}
log.Println(lang.GetWithTemplate(keys.Test, template))
```
> description
to change the locate


[^1]: Log with structured format and variables
