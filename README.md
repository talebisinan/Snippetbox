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

#### Creating a test DB
```bash
CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

```
CREATE USER 'test_web'@'localhost';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'localhost';
ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';
```
#### Running Tests
IDE's should automatically pickup the `*_test.go` files and run them. If not, run the following command:
```bash
go test ./...
```

#### Coverage
```bash
go test ./... -cover
```