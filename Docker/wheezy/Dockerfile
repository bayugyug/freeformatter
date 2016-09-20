#start
#FROM golang
#FROM google/debian:jessie
FROM debian:wheezy

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

#prepare
RUN apt-get update -y && apt-get install -y --no-install-recommends \
curl tar git ca-certificates apt-transport-https go \
&& rm -rf /var/lib/apt/lists/*

ENV GOHOME=/mongers
ENV GOROOT=/usr/local/go
ENV GOPATH=$GOHOME/go
ENV GOSRC=$GOPATH/src
ENV PATH=$PATH:$GOPATH/bin:/usr/local/go/bin


RUN mkdir -p $GOSRC $GOSRC/github.com/bayugyug $GOPATH/bin && cp cleanup.sh $GOPATH/bin/


RUN cd $GOSRC/github.com/bayugyug && \
    git clone https://github.com/bayugyug/freeformatter.git && \
    cd $GOSRC/github.com/bayugyug/freeformatter && \
    go get -v && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags netgo -installsuffix netgo -v -ldflags "-X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" . && \
    cd $GOSRC && \
    go install github.com/bayugyug/freeformatter && \
    cd $GOSRC/github.com/bayugyug/freeformatter  && \
    cp -f freeformatter Docker/cleanup.sh $GOPATH/bin/

#executable
RUN chmod +x $GOPATH/bin/* && \
    $GOPATH/bin/cleanup.sh && \
    rm -fr $GOSRC

#main entry
ENTRYPOINT [ "/mongers/go/bin/freeformatter" ]
