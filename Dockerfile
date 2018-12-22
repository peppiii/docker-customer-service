## Versi Golang yang dipakai
FROM golang:1.9

## Creating Container
MAINTAINER Febry Rizky Wardani febry.rizky@cbncloud.co.id
ADD . /go/src/dockercustomers

## compiling package yang dibutuhkan
RUN go get github.com/gin-gonic/gin
RUN go get -d github.com/go-sql-driver/mysql
RUN go get -d github.com/jinzhu/gorm
##RUN go get -u -v golang.org/x/crypto/bcrypt
RUN go install customers

#ENTRYPOINT /go/bin/dockerlogins

## setup environment variable yang akan dipakai
ENV DB_HOST 
ENV DB_PORT 
ENV DB_USER 
ENV DB_PASS

## Expose hasil golang berdasarkan port yang disetting
EXPOSE 8081

RUN mkdir -p /go/src/dockercustomers
COPY . /go/src/dockercustomers
WORKDIR /go/src/dockercustomers

CMD go run customer.go
