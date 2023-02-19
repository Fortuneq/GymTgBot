FROM golang:1.19.5-alpine3.17 as build

WORKDIR /app

COPY ../bot .

RUN go build -o main cmd/app/main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=build /app/config.yml .
COPY --from=build /app/main .

ENTRYPOINT ["./main"]
