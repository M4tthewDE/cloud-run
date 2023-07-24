set dotenv-load

run:
    go run cmd/server/main.go

client CMD:
    go run cmd/client/main.go "{{CMD}}"

clean:
    rm main

build: 
    CGO_ENABLED=0 GOOS=linux go build cmd/server/main.go

debug:
    dlv debug main.go

docker-build: 
    docker build --tag ghcr.io/m4tthewde/cloud-run:latest .

docker-release version:
    git checkout {{version}}
    docker build --tag ghcr.io/m4tthewde/cloud-run:{{version}} .
    docker push ghcr.io/m4tthewde/cloud-run:{{version}}

