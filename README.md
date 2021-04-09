# Water Jug Challenge

2021-04-09

I try to cover production-like requirements below within a limit timeline

- Using web framework `echo`
- Authorization middleware (demo with weak plain text comparison in Header, production should use jwt token for authz)
- Structural logging
- Request body validation
- Unit testing

## Devlopment

```
go build
./water-jug


curl --request POST \
  --url http://localhost:8080/problem \
  --header 'content-type: application/json' \
  --header 'x-auth-token: 123456' \
  --data '{
	"x": 2,
	"y": 10,
	"z": 4
}'
```

## Test

```
go test -v ./service
go test -v ./controller
```
