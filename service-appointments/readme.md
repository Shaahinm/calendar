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