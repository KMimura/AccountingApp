FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN go get -u github.com/gin-gonic/gin && go get -u github.com/jinzhu/gorm && go get -u github.com/go-sql-driver/mysql
RUN mkdir /usr/src && mkdir /usr/src/api
COPY ./api /usr/src/api
WORKDIR /usr/src/api
CMD ["go","run","main.go"]
