#start
#FROM golang:alpine
FROM alpine:3.4

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

RUN apk add --no-cache ca-certificates

RUN set -ex \
    && apk add --no-cache --virtual .build-deps \
        bash \
        git \
        gcc \
        musl-dev \
        openssl \
        go

RUN apk add --update ca-certificates # Certificates for SSL

ENV GOPATH /go
ENV GOSRC $GOPATH/src
ENV GOBIN $GOPATH/bin
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin

#prepare
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

#home
RUN mkdir -p $GOSRC/github.com/bayugyug && cd $GOSRC/github.com/bayugyug && git clone https://github.com/bayugyug/freeformatter.git

#compile
RUN cd $GOSRC/github.com/bayugyug/freeformatter && go get -v && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -ldflags "-s -w -X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" .

#ADD
RUN cd $GOSRC/github.com/bayugyug/freeformatter && cp -f freeformatter $GOBIN/

#executable
RUN chmod +x $GOBIN/*

#cleanup
RUN du -sh /go/* /* && ls -ltrha $GOBIN/*

RUN apk del .build-deps go git gcc musl-dev && \
    rm -rf /var/lib/{apt,dpkg,cache,log}/ &&  \
    rm -rf /var/{cache,log} && \
    rm -fr /go/pkg  /go/src

#cleanup
RUN du -sh /go/* /*

EXPOSE 7777 7778 7779 7780 7781 7782 7783 7784 7785 7786 7787 7788 7789 7790 7791 7792 7793 7794 7795 7796 7797 7798 7799 7800

#main entry
ENTRYPOINT [ "/go/bin/freeformatter" ]
