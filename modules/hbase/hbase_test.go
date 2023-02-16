package hbase

import (
	"github.com/netdata/go.d.plugin/agent/module"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func TestHDFS_InitErrorOnCreatingClientWrongTLSCA(t *testing.T) {
	job := New()
	job.Client.TLSConfig.TLSCA = "testdata/tls"

	assert.False(t, job.Init())
}

func TestHBase_Check(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write(testMasterNodeData)
			}))
	defer ts.Close()

	job := New()
	job.URL = ts.URL
	require.True(t, job.Init())

	assert.True(t, job.Check())
	assert.NotZero(t, job.nodeType)
}

func TestHBase_CheckMasterNode(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write(testMasterNodeData)
			}))
	defer ts.Close()

	job := New()
	job.URL = ts.URL
	require.True(t, job.Init())

	assert.True(t, job.Check())
	assert.Equal(t, masterType, job.nodeType)
}