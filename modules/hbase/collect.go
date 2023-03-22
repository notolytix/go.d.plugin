package hbase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/netdata/go.d.plugin/pkg/stm"
)

type (
	rawData map[string]json.RawMessage
	rawJMX  struct {
		Beans []rawData
	}
)

func (r rawJMX) isEmpty() bool {
	return len(r.Beans) == 0
}

func (r rawJMX) find(f func(rawData) bool) rawData {
	for _, v := range r.Beans {
		if f(v) {
			return v
		}
	}
	return nil
}

func (r rawJMX) findJvm() rawData {
	f := func(data rawData) bool { return string(data["modelerType"]) == "\"JvmMetrics\"" }
	return r.find(f)
}

func (r rawJMX) findMaster() rawData {
	f := func(data rawData) bool { return string(data["modelerType"]) == "\"Master,sub=IPC\"" }
	return r.find(f)
}

func (r rawJMX) findMasterMetrics() rawData {
	f := func(data rawData) bool { return string(data["modelerType"]) == "\"Master,sub=Server\"" }
	return r.find(f)
}

func (r rawJMX) findRegionServer() rawData {
	f := func(data rawData) bool { return string(data["modelerType"]) == "\"RegionServer,sub=IPC\"" }
	return r.find(f)
}

func (h *HBase) collect() (map[string]int64, error) {
	var raw rawJMX
	err := h.client.doOKWithDecodeJSON(&raw)
	if err != nil {
		return nil, err
	}

	if raw.isEmpty() {
		return nil, errors.New("empty response")
	}

	mx := h.collectRawJMX(raw)

	return stm.ToMap(mx), nil
}

func (h HBase) collectRawJMX(raw rawJMX) *metrics {
	var mx metrics
	switch h.nodeType {
	default:
		panic(fmt.Sprintf("unsupported node type : '%s'", h.nodeType))
	case masterType:
		h.collectMaster(&mx, raw)
	case regionServerType:
		h.collectRegionServer(&mx, raw)
	}
	return &mx
}

func (h HBase) collectMaster(mx *metrics, raw rawJMX) {
	err := h.collectJVM(mx, raw)
	if err != nil {
		h.Debugf("error on collecting jvm : %v", err)
	}
	err = h.collectMasterMerrics(mx, raw)
	if err != nil {
		h.Debugf("error on collecting master metrics : %v", err)
	}
}

func (h HBase) collectRegionServer(mx *metrics, raw rawJMX) {
	err := h.collectJVM(mx, raw)
	if err != nil {
		h.Debugf("error on collecting region server metrics : %v", err)
	}
}

func (h HBase) collectJVM(mx *metrics, raw rawJMX) error {
	v := raw.findJvm()
	if v == nil {
		return nil
	}

	var jvm jvmMetrics
	err := writeJSONTo(&jvm, v)
	if err != nil {
		return err
	}

	mx.Jvm = &jvm
	return nil
}

func (h HBase) collectMasterMerrics(mx *metrics, raw rawJMX) error {
	v := raw.findMasterMetrics()
	if v == nil {
		return nil
	}

	var master masterMetrics
	err := writeJSONTo(&master, v)
	if err != nil {
		return err
	}

	mx.Master = &master
	return nil
}

func writeJSONTo(dst interface{}, src interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}
