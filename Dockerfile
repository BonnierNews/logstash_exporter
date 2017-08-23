FROM golang:1.8 as golang
RUN go get -u github.com/kardianos/govendor && \
        go get -u github.com/DagensNyheter/logstash_exporter && \
        cd $GOPATH/src/github.com/DagensNyheter/logstash_exporter && \
        govendor build +local

FROM quay.io/prometheus/busybox:glibc
COPY --from=golang /go/bin/logstash_exporter /
LABEL maintainer christoffer.kylvag@bonniernews.se
EXPOSE 9198
ENTRYPOINT ["/logstash_exporter"]  
