FROM golang:1.11 as golang
ARG GODEP_VERSION=v0.5.0

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/${GODEP_VERSION}/dep-linux-amd64 && \
        chmod +x /usr/local/bin/dep && \
        go get -u github.com/BonnierNews/logstash_exporter && \
        cd $GOPATH/src/github.com/BonnierNews/logstash_exporter && \
        dep ensure && \
        make

FROM busybox:1.30.0-glibc
COPY --from=golang /go/src/github.com/BonnierNews/logstash_exporter/logstash_exporter /
LABEL maintainer christoffer.kylvag@bonniernews.se
EXPOSE 9198
ENTRYPOINT ["/logstash_exporter"]  
