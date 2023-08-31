# Snippetbox

#### Runnnig the application (default port 4000)
Pass -addr flag to change the port like `-addr=":80"`.
```bash
go run ./cmd/web
```

#### Viewing all flags
```bash
go run ./cmd/web -help
```

```bash
go run ./cmd/web -addr=":4000" -dsn="user:pass@/snippetbox?parseTime=true"
```

#### Generating a self-signed TLS certificate
1. create a tls directory under the root of the project.

2. Locate go and find `generate_cert.go` under src/crypto/tls.

3. Run the following command to generate a self-signed certificate and private key:
```bash
cd tls
go run location_in_first_step_/generate_cert.go --rsa-bits=2048 --host=localhost
```