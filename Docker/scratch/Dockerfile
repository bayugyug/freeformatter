FROM scratch


#environment
ENV GOHOME=
ENV GOPATH=$GOHOME/go
ENV PATH=$PATH:$GOPATH/bin

COPY ./freeformatter /freeformatter

ENTRYPOINT [ "/freeformatter" ]

