# Snippetbox

#### Runnnig the application (default port 4000)
```bash
go run ./cmd/web -dsn="user:pass@/snippetbox?parseTime=true"
```

#### Viewing all flags
```bash
go run ./cmd/web -help
```

#### Generating a self-signed TLS certificate
1. create a tls directory under the root of the project.

2. Locate go and find `generate_cert.go` under src/crypto/tls.

3. Run the following command to generate a self-signed certificate and private key:
```bash
cd tls
go run location_in_first_step_/generate_cert.go --rsa-bits=2048 --host=localhost
```

#### Running Tests
IDE's should automatically pickup the `*_test.go` files and run them. If not, run the following command:
```bash
go test ./cmd/web
```
