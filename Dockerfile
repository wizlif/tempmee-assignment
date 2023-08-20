FROM golang:1.19.3-alpine3.16 as builder
WORKDIR /app
COPY . .
ARG SERVER_PORT
RUN apk update && apk add build-base
RUN go mod download
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/db/migration ./db/migration
COPY app.env .
COPY db/migration ./db/migration

EXPOSE $SERVER_PORT
CMD [ "/app/main" ]
