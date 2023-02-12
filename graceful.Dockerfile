FROM golang:1.16-alpine as builder
WORKDIR /app/
COPY ./ ./
RUN go build -o go-server graceful/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/go-server ./
CMD ["./go-server"]