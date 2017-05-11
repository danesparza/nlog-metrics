# nlog-metrics [![CircleCI](https://circleci.com/gh/danesparza/nlog-metrics.svg?style=svg)](https://circleci.com/gh/danesparza/nlog-metrics)
Tracks logging metrics using InfluxDB

## Quick start 

* Make sure you've already got [InfluxDB](https://docs.influxdata.com/influxdb/v1.2/introduction/installation/) up and running.  In InfluxDb, [create the database](https://docs.influxdata.com/influxdb/v1.2/guides/writing_data/#creating-a-database-using-the-http-api) that you want to log to.
* Get the [latest release](https://github.com/danesparza/nlog-metrics/releases/latest) for your platform (it's just a single executable) 
* Generate a config file:
```
nlog-metrics config > config.yaml
```

Update the config file sections with information for your NLog database and your InfluxDB database:
``` yaml
sqlserver:
  server: servernamehere
  database: system_logging
  user: username
  password: password
influxdb:
  server: http://localhost:8086
  database: applogs
  measurement: metrics
```

When you're ready to start logging metrics, just run 

```
nlog-metrics start
```

Note: To run this as a Windows service, check out the install.bat/uninstall.bat files in the dist directory (also packaged with the windows release zip file)
