FROM golang:1.14-alpine AS build
WORKDIR /build

# RUN apk add git npm build-base

ADD . .

RUN go mod tidy

RUN go build -o ./bin/count-server ./cmd/server/main.go


FROM alpine:latest AS final
WORKDIR /app
COPY --from=build /build/bin .

RUN apk add ca-certificates

EXPOSE 8080

ENV VC_WS_ADDR='0.0.0.0:8080' \
    VC_LOG_LEVEL='4'

CMD ./count-server