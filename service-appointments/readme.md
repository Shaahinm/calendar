go env -w CGO_ENABLED=1
go run .\cmd\main.go
go run .\cmd

# tools

https://github.com/avelino/awesome-go

# for sqlite

https://jmeubank.github.io/tdm-gcc/download/

# todo

# db relations

# structured logging

# test

# mux

# middlewares

# authentication

# opentelemetry

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

