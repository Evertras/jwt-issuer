FROM golang:1.13-alpine AS go-builder

WORKDIR /go/src/github.com/Evertras/jwt-issuer
COPY . .
ARG BUILD_TAGS

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build -a -tags "netgo ${BUILD_TAGS}" \
        -ldflags '-w -extldflags "-static"' \
        -o /server \
        ./cmd/server/main.go

FROM scratch

COPY --from=go-builder /server server

ENTRYPOINT ["/server"]

