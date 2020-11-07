FROM golang:1.13-alpine AS build-env
WORKDIR /go/src/github.com/SaKu2110/sign-server
COPY ./ ./
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/SaKu2110/sign-server/server /usr/local/bin/server
ENV DB_USER auth_user
ENV DB_PASS password
ENV DB_IP mysql
ENV DB_PORT 3306
ENV DB_NAME auth
ENV DB_MAX_CONNS 100
ENV DB_MAX_IDLE 100
ENV DB_MAX_LIFETIME 100

EXPOSE 9000
CMD ["/usr/local/bin/server"]
