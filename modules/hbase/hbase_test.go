package hbase

import (
	"github.com/netdata/go.d.plugin/agent/module"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	testMasterNodeData, _ = ioutil.ReadFile("testdata/masternode.json")
)

func Test_readTestData(t *testing.T) {
	assert.NotNil(t, testMasterNodeData)
}

func TestNew(t *testing.T) {
	assert.Implements(t, (*module.Module)(nil), New())
}

func TestHDFS_Init(t *testing.T) {
	job := New()

	assert.True(t, job.Init())
}