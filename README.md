# grpc

![Gopher Share](https://github.com/go-stuff/images/blob/master/GOPHER_SHARE_640x320.png)

Testing grpc

# Environment Variables Required

The following are default values if none are specified:

```bash
MONGO_URI = "mongodb://localhost:27017"
MONGO_DATABASE = "test"
API_URI = "localhost:6000"
```

# Certs

I have included some test certs in the package to connect with `go-stuff\web`. You can generate them the following way:

``` bash
go run GOROOT/src/crypto/tls/generate_cert.go --host 127.0.0.1 --duration 17520h
```
