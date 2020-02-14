FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN go get -u github.com/gin-gonic/gin && mkdir /usr/src/api
COPY ./api /usr/src/api
WORKDIR /usr/src/api
CMD ["go","run","main.go"]
