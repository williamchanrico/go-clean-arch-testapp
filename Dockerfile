FROM golang:1.12-alpine as build

WORKDIR /src

COPY . /src/

RUN apk add --no-cache git protobuf make && \
    go get github.com/golang/protobuf/protoc-gen-go && \
    protoc --proto_path=grpc --go_out=plugins=grpc:grpc grpc/*.proto && \
    GOPATH="/go" GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/xtest cmd/main.go && \
    echo 'nobody:x:65534:' > /src/group.nobody && \
    echo 'nobody:x:65534:65534::/:' > /src/passwd.nobody && \
    GRPC_HEALTH_PROBE_VERSION=v0.2.2 && \
    wget -q -O /bin/grpc-health-probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc-health-probe

FROM gcr.io/distroless/static

EXPOSE 9000
EXPOSE 50051

COPY --from=build /src/group.nobody /etc/group
COPY --from=build /src/passwd.nobody /etc/passwd
USER nobody:nobody

COPY --from=build /bin/xtest /bin/xtest
# Add grpc-health-probe to use with readiness and liveness probes
COPY --from=build /bin/grpc-health-probe /bin/grpc-health-probe

ENTRYPOINT ["/bin/xtest"]
