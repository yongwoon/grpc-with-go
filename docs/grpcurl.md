# grpcurl

- install

```bash
brew install grpcurl
```

## common

- サーバー内に実装されているサービス一覧の確認

```bash
grpcurl -plaintext localhost:8080 list
```

- `GreetingService` service が持つ method list

```bash
grpcurl -plaintext localhost:8080 list GreetingService
```

## Call `GreetingService.hello`

```bash
 grpcurl -plaintext -d '{"name": "hsaki"}' localhost:8080 GreetingService.Hello
```
