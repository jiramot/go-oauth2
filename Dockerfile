FROM golang:1.17-alpine as builder
RUN apk update && apk upgrade && \
    apk --update add git make bash

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
RUN make build

FROM alpine
LABEL org.opencontainers.image.source="ghcr.io/jiramot/go-oauth2"

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata

WORKDIR /app
EXPOSE 8080 8081 8082
COPY --from=builder /app/bin/admin /app/admin
COPY --from=builder /app/bin/public /app/public
COPY --from=builder /app/bin/app_gallery /app/app_gallery
COPY --from=builder /app/config.yaml /app/config.yaml

CMD /app/public