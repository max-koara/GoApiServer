FROM golang:alpine
ADD ./src /go/src/app
WORKDIR /go/src/app

RUN go mod init domain 
RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4/middleware 
CMD ["go", "run",  "main.go"]