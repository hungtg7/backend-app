# grpc-api-app


## Running

Running app `app/app_data_monitoring/main.go`. 

You can configure
the port used with the `$PORT` environment variable, and to serve on HTTP set
`$SERVE_HTTP=true`. Default will serve on HTTPs.

```
$ go run app/app_data_monitoring/main.go
```

## Getting started

After cloning the repo, there are a couple of initial steps;

1. Install the generate dependencies with `make install`.
   This will install `buf`, `protoc-gen-go`, `protoc-gen-go-grpc`, `protoc-gen-grpc-gateway` and
   `protoc-gen-openapiv2` which are necessary for us to generate the Go and swagger files.
1. Replace `yourscmprovider.com/youruser/yourrepo` with your cloned repo path:
   ```bash
   $ find . -path ./vendor -prune -o -type f \( -name '*.go' -o -name '*.proto' \) -exec sed -i -e "s;github.com/hungtran150/grpc-api-app;yourscmprovider.com/youruser/yourrepo;g" {} +
   ```
1. Finally, generate the files with `make generate`.

Now you can run the web server with `go run app/app_data_monitoring/main.go`.

## Docker build
```bash
docker build --build-arg SLACK_WEB_HOOK=$SLACK_WEB_HOOK -t api-app .
```

```bash
docker run --rm -p 11000:11000 -e SLACK_WEB_HOOK=$SLACK_WEB_HOOK api-app ./main
```
