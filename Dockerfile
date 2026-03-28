FROM golang:1.25-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app/bin/app .
EXPOSE 8090
CMD ["./app"]
