#TODO: dockerize the app
FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /go/src/app
COPY . .
RUN go build -o main main.go

FROM alpine:3.18.0
WORKDIR /app
COPY --from=builder /go/src/app/main /app
CMD [ "/app/main" ]