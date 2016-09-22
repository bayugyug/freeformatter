#start
FROM golang

#auth
MAINTAINER bayugyug<bayugyug@gmail.com>

#prepare
RUN apt-get update -y && apt-get install -y --no-install-recommends \
git \
&& rm -rf /var/lib/apt/lists/*

#environment 
ENV GOHOME=
ENV GOPATH=$GOHOME/go
ENV GOSRC=$GOPATH/src
ENV PATH=$PATH:$GOPATH/bin

RUN mkdir -p $GOPATH/bin  $GOSRC/github.com/bayugyug

RUN cd $GOSRC/github.com/bayugyug && git clone https://github.com/bayugyug/freeformatter.git

#compile
RUN cd $GOSRC/github.com/bayugyug/freeformatter && go get -v && go build -tags netgo -installsuffix netgo -installsuffix cgo -v -ldflags "-s -w -X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`"

#install
RUN cd $GOSRC && go install github.com/bayugyug/freeformatter

#ADD
RUN cd $GOSRC/github.com/bayugyug/freeformatter  && cp -f Docker/cleanup.sh $GOPATH/bin/

#executable
RUN chmod +x $GOPATH/bin/*

#cleanup
RUN $GOPATH/bin/cleanup.sh

WORKDIR $GOSRC/github.com/bayugyug/freeformatter

EXPOSE 7777 7778 7779 7780 7781 7782 7783 7784 7785 7786 7787 7788 7789 7790 7791 7792 7793 7794 7795 7796 7797 7798 7799 7800

#main entry
ENTRYPOINT [ "freeformatter" ]
