# Pacote `cmd/payments`

A ser documentado.

Mas por enquanto, aqui vão alguns exemplos: 

```terminal
$ curl localhost:8084/payments
$ curl -X POST localhost:8084/payments -d '{"id": "some-uuid", "subscription_id":"jpkroehling", "amount":99, "status":"FAILED"}'
```
