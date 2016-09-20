FROM busybox:ubuntu-14.04

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

COPY ./freeformatter /freeformatter

ENTRYPOINT [ "/freeformatter" ]

