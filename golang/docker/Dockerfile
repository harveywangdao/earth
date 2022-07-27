FROM golang
MAINTAINER thomas 1272777053@qq.com

COPY . $GOPATH/src/

WORKDIR $GOPATH/src/github.com/beego/beeweb

RUN go build

EXPOSE 8080

ENTRYPOINT ["./beeweb"]
