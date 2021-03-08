FROM golang:latest
MAINTAINER brook1711 "brook1711@163.com"  
WORKDIR $GOPATH/src/web_v1
ENV PATH $JAVA_HOME/bin:$PATH
ADD ./web_v1/ $GOPATH/src/web_v1  
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go get -v github.com/beego/bee 
RUN go mod tidy && bee pack -be GOOS=linux

EXPOSE 8080