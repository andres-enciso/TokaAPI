
FROM golang:1.25-alpine AS build
WORKDIR /src
ENV GOPROXY=https://proxy.golang.org,direct
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod go mod tidy


RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./api/main.go

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata && adduser -D -H -u 10001 appuser
WORKDIR /app
COPY --from=build /bin/app /app/app
EXPOSE 8080
USER appuser
ENTRYPOINT ["/app/app"]
