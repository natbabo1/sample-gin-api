FROM golang:1.24.5-alpine3.22 AS builder
WORKDIR /src
COPY . .
RUN go build -o /app ./cmd/server

FROM alpine:3.20
RUN adduser -D -g '' app
USER app
COPY --from=builder /app /app
EXPOSE 8080
CMD ["/app"]
