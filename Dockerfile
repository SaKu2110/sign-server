FROM golang:latest

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

RUN go get github.com/SaKu2110/sign-server
WORKDIR /go/src/github.com/SaKu2110/sign-server

RUN go get github.com/gin-gonic/gin
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/mysql

ENV DB_USER keeper
ENV DB_PASS admin0000
ENV DB_IP mysql
ENV DB_PORT 3306
ENV DB_NAME sign

EXPOSE 8080

CMD ["go", "run", "/go/src/github.com/SaKu2110/sign-server/main.go"]
