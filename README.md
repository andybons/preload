# Preload

Test case for double-download bug with link rel="preload" and `responseType` set
to `blob`.

```sh
go run main.go
```

Visit [localhost:8080](http://localhost:8080)

Optionally, a `timeout` search parameter can be specified for how long to wait
(in ms) before the request for the resource is made.
