FROM golang:1.19 as builder
WORKDIR /doki
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-w' -o app ./cmd/main.go

FROM alpine
WORKDIR /
COPY --from=builder /doki/app ./app
ENTRYPOINT ["/app"]