FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app src/main.go

FROM scratch
COPY --from=builder /app/app .

EXPOSE 3000
CMD ["/app"]