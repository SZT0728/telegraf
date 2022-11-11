## Configuration

```toml @sample.conf
# Read metrics about custom-ceph usage
[[inputs.custom-ceph]]
  ## where the ceph binary
  ceph_binary = /etc/ceph/ceph.config
  ## the ceph user
  ceph_user = admin
```

## Data format to consume.
## Each data format has its own unique set of configuration options, read
## more about them here:
## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
data_format = "influx"