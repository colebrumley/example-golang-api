# Example HTTP REST API in Go

Small demo/boilerplate of golang's `net/http` package.

**Flags**
```sh
  -cert="": TLS cert file
  -key="": TLS key file
  -listen=":80": Address to listen on
  -password="": Enable basic authentication
  -tls=false: Use TLS (requires -cert and -key)
```
**To run:** `go run example.go basicauth.go -h`

**To build:** `go build example.go basicauth.go`
