# go-log-headers
simple golang app to print request headers to stdout json logs

```
go build
./headers
curl localhost:8082
```

```
docker build -t go-headers:1.0 .
docker run --name go-headers go-headers:1.0
curl localhost:8082
```