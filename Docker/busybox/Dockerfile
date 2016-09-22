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

EXPOSE 7777 7778 7779 7780 7781 7782 7783 7784 7785 7786 7787 7788 7789 7790 7791 7792 7793 7794 7795 7796 7797 7798 7799 7800

ENTRYPOINT [ "/freeformatter" ]

