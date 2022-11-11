//go:generate ../../../tools/readme_config_includer/generator
package custom_ceph

import (
	_ "embed"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type CustomCeph struct {
	CephBinary string `toml:"ceph_binary"`
	CephUser   string `toml:"ceph_user"`
}

//go:embed sample.conf
var sampleConfig string

func (*CustomCeph) SampleConfig() string {
	return sampleConfig
}

func (*CustomCeph) Init() error {
	return nil
}

func (*CustomCeph) Gather(acc telegraf.Accumulator) error {
	acc.AddFields("state", map[string]interface{}{"value": "pretty good"}, map[string]string{"tag-key": "tag-value"})
	acc.AddGauge("test-measurement", map[string]interface{}{"gauge-value": 111}, map[string]string{"guage-tag-key": "guage-tag-value"})
	return nil
}

func init() {
	inputs.Add("custom-ceph", func() telegraf.Input { return &CustomCeph{} })
}
