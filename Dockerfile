FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY go.* /app/
RUN go mod download
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/ibuYemekApi

FROM alpine:3.15
COPY --from=builder /app/ibuYemekApi /app/
RUN apk update && apk add curl
CMD [ "/app/ibuYemekApi" ]
