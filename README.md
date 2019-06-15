# grpc

![Gopher Share](https://github.com/go-stuff/images/blob/master/GOPHER_SHARE_640x320.png)

Testing grpc

# Environment Variables Required

```bash
MONGO_URI
MONGO_DATABASE
API_URI
```

# Test Certs

I have included some test certs in the package. You can generate them the following way.

``` bash
go run GOROOT/src/crypto/tls/generate_cert.go --host 127.0.0.1 --duration 17520h
```
