//go:build !custom || inputs || inputs.custom_ceph

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/custom_ceph" // register plugin
