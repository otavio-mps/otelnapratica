# Pacote `api`

A ser documentado.

Mas por enquanto, se os protobufs precisarem ser gerados novamente:

```terminal
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./api/plan.proto
```
