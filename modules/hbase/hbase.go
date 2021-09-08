package hbase

import (
	"errors"
	"time"

	"github.com/netdata/go.d.plugin/pkg/web"

	"github.com/netdata/go.d.plugin/agent/module"
)

func init() {
	creator := module.Creator{
		Create: func() module.Module { return New() },
	}

	module.Register("hbase", creator)
}

// New creates HBase with default values.
func New() *HBase {
	config := Config{
		HTTP: web.HTTP{
			Request: web.Request{
				URL: "http://127.0.0.1:50070/jmx",
			},
			Client: web.Client{
				Timeout: web.Duration{Duration: time.Second}},
		},
	}

	return &HBase{
		Config: config,
	}
}

type nodeType string

const (
	masterType nodeType = "master"
	regionServerType nodeType = "regionserver"
)

type hbaseVersion string

const (
	v1_6 hbaseVersion = "1_6"
	v2_0 hbaseVersion = "2_0"
)

// Config is the HBase module configuration.
type Config struct {
	web.HTTP `yaml:",inline"`
}

// HBase HBase module.
type HBase struct {
	module.Base
	Config `yaml:",inline"`

	nodeType
	client  *client
}

// Cleanup makes cleanup.
func (HBase) Cleanup() {}

func (h HBase) createClient() (*client, error) {
	httpClient, err := web.NewHTTPClient(h.Client)
	if err != nil {
		return nil, err
	}

	return newClient(httpClient, h.Request), nil
}

func (h HBase) determineNodeType() (nodeType, error) {
	var raw rawJMX
	err := h.client.doOKWithDecodeJSON(&raw)
	if err != nil {
		return "", err
	}

	if raw.isEmpty() {
		return "", errors.New("empty response")
	}

	master := raw.findMaster()
	if (master != nil) {
		return nodeType(masterType), nil
	}

	region := raw.findRegionServer()
	if (region != nil) {
		return nodeType(regionServerType), nil
	}

	return "", errors.New("Unknown node type!")
}

// Init makes initialization.
func (h *HBase) Init() bool {
	cl, err := h.createClient()
	if err != nil {
		h.Errorf("error on creating client : %v", err)
		return false
	}
	h.client = cl

	return true
}

// Check makes check.
func (h *HBase) Check() bool {
	t, err := h.determineNodeType()
	if err != nil {
		h.Errorf("error on node type determination : %v", err)
		return false
	}
	h.nodeType = t

	return len(h.Collect()) > 0
}

// Charts returns Charts.
func (h HBase) Charts() *Charts {
	switch h.nodeType {
	default:
		return nil
	case masterType:
		return masterServerCharts()
	case regionServerType:
		return regionserverCharts()
	}
}

// Collect collects metrics.
func (h *HBase) Collect() map[string]int64 {
	mx, err := h.collect()

	if err != nil {
		h.Error(err)
	}

	if len(mx) == 0 {
		return nil
	}

	return mx
}
