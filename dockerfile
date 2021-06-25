FROM golang:1.16-alpine AS builder
WORKDIR /
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./gin ./src/main.go

FROM alpine:latest AS runner
WORKDIR /root/
ENV GIN_MODE=release
COPY --from=builder ./gin ./gin
COPY --from=builder ./.env ./.env
CMD ["./gin"]