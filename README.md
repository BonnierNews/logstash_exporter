# Logstash exporter
Prometheus exporter for the metrics available in Logstash since version 5.0.

## Usage

```bash
go get -u github.com/BonnierNews/logstash_exporter
cd $GOPATH/src/github.com/BonnierNews/logstash_exporter
make
./logstash_exporter -web.listen_address :1234 --logstash.endpoint http://localhost:1235
```

### Flags
Flag | Description | Default
-----|-------------|---------
--web.listen_address | Exporter bind address | :9198
--logstash.endpoint | Metrics endpoint address of logstash | http://localhost:9600

## Implemented metrics
* Node metrics
