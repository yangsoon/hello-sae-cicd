FROM golang:1.19 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM alpine/curl:3.14
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000
ENTRYPOINT ["./main"]


