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
    docker build --tag gcr.io/esc-api-384517/cloud-run:latest .

docker-push:
    docker build --tag gcr.io/esc-api-384517/cloud-run:latest .
    docker push gcr.io/esc-api-384517/cloud-run:latest

infra: docker-push destroy-infra
    cd infra; terraform apply -auto-approve

destroy-infra:
    cd infra; terraform destroy -auto-approve
