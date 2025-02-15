FROM --platform=$BUILDPLATFORM golang AS builder

WORKDIR /app

RUN --mount=target=. \
        --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg \
        GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 \
    go build -o /deploy/server/service ./cmd/service/main.go && \
    cp -r config /deploy/server/config

FROM alpine

WORKDIR /app

ENV VERV_NAME=makosh

COPY --from=builder /deploy/server/ .

EXPOSE 8080

ENTRYPOINT ["./service"]