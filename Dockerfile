FROM golang:1.6
LABEL maintainer christoffer.kylvag@bonniernews.se
EXPOSE 9198
RUN go get -u github.com/kardianos/govendor && \
        go get -u github.com/DagensNyheter/logstash_exporter && \
        cd $GOPATH/src/github.com/DagensNyheter/logstash_exporter && \
        govendor build +local && \
        mv /go/bin/logstash_exporter /
ENTRYPOINT ["/logstash_exporter"]  
