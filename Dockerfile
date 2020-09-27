FROM golang:1.15-alpine AS builder
WORKDIR /build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:latest
COPY --from=builder /build/app /go/bin/app
EXPOSE 8080
CMD ["/go/bin/app"]