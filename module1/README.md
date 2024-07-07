# Build and run app

```bash
cd src
go mod init demo # creates dependencies file with module name 'demo'

cd ..
pwd # should be in 'module1' directory
go build -o bin/app src/main.go # builds app
bin/app # runs app
```