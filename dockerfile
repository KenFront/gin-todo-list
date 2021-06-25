FROM golang:1.16-alpine AS builder
WORKDIR /
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app ./src/main.go

FROM alpine:latest
WORKDIR /root/
ENV GIN_MODE=release
COPY --from=builder ./app ./app
COPY --from=builder ./.env ./.env
CMD ["./app"]