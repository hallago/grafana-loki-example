# grafana-loki-example

An example logging system using Grafana Loki.

<h3 align="center">
<img src="img/Grafana-Loki.png" alt="struct" />
</h3>

## Usage

### Docker

```
// start
$ cd grafana-loki-example
$ docker-compose up -d

// stop
$ docker-compose stop
```

### Grafana WebPage

```
http://localhost:3000
User: admin
Pass: admin
```

## ScreenShot

<img src="img/2023-04-23-6.55.04.png" alt="01" />
<img src="img/2023-04-23-6.55.44.png" alt="02" />
<img src="img/2023-04-23-6.56.18.png" alt="03" />


## Container

* `server` - Log files are saved to `./server/logs`
* `loki` - Store logs received from Promtail.
* `promtail` - Collect string data from log files and send to Loki.
* `grafana` - Visualize data from Loki.

## Grafana Data Sources

- Loki: `http://loki:3100`