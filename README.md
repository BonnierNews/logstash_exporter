# Logstash exporter
Prometheus exporter for the metrics available in Logstash since version 5.0.

## Usage
```bash
go get -u github.com/kardianos/govendor
go get -u github.com/DagensNyheter/logstash_exporter
cd $GOPATH/src/github.com/DagensNyheter/logstash_exporter
govendor build +local
./logstash_exporter -exporter.bind_address :1234 -logstash.endpoint http://localhost:1235
```

### Flags
Flag | Description | Default
-----|-------------|---------
-exporter.bind_address | Exporter bind address | :9198
-logstash.endpoint | Metrics endpoint address of logstash | http://localhost:9600

## Implemented metrics
* Node metrics