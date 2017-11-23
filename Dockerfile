FROM golang:1.9 as golang
RUN go get -u github.com/kardianos/govendor && \
        go get -u github.com/BonnierNews/logstash_exporter && \
        cd $GOPATH/src/github.com/BonnierNews/logstash_exporter && \
        govendor build +local

FROM busybox:1.27.2-glibc
COPY --from=golang /go/bin/logstash_exporter /
LABEL maintainer christoffer.kylvag@bonniernews.se
EXPOSE 9198
ENTRYPOINT ["/logstash_exporter"]  
