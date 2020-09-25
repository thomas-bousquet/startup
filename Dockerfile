FROM golang:1.15-alpine AS build
WORKDIR /builder
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app

FROM golang:1.15-alpine
COPY --from=build /builder /go/bin/builder
EXPOSE 8080
CMD ["/go/bin/builder/app"]