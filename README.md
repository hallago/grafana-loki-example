# grafana-loki-example

An example logging system using Grafana Loki.

## API Server

- port 80
- Log files are saved to `./api/runtime/logs`

## Loki

Store logs received from Promtail.

- port 3100

## Promtail

Collect string data from log files and send to Loki.

## Grafana

Visualize data from Loki.

