package client

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/enablers"
	"github.com/zubairhamed/betwixt/objdefs/oma"
	"github.com/zubairhamed/betwixt/registry"
	"testing"
)

func TestClient(t *testing.T) {

	registry := registry.NewDefaultObjectRegistry()
	cli := NewDefaultClient(":0", "localhost:5683", registry)
	assert.NotNil(t, cli, "Error instantiating client")
	assert.NotNil(t, registry, "Error instantiating registry")

	cases1 := []struct {
		in LWM2MObjectType
	}{
		{oma.OBJECT_LWM2M_SERVER},
		{oma.OBJECT_LWM2M_DEVICE},
		{oma.OBJECT_LWM2M_SECURITY},
	}

	for _, c := range cases1 {
		err := cli.EnableObject(c.in, nil)

		assert.NotNil(t, err, "Object should already be enabled: ", c.in)
	}

	cases2 := []struct {
		in LWM2MObjectType
		en ObjectEnabler
	}{
		{oma.OBJECT_LWM2M_ACCESS_CONTROL, enablers.NewNullEnabler()},
		{oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, enablers.NewNullEnabler()},
		{oma.OBJECT_LWM2M_FIRMWARE_UPDATE, enablers.NewNullEnabler()},
		{oma.OBJECT_LWM2M_LOCATION, enablers.NewNullEnabler()},
		{oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, enablers.NewNullEnabler()},
	}

	for _, c := range cases2 {
		err := cli.EnableObject(c.in, c.en)

		assert.Nil(t, err, "Error enabling object: ", c.in)
	}

	cases3 := []struct {
		in LWM2MObjectType
	}{
		{oma.OBJECT_LWM2M_SERVER},
		{oma.OBJECT_LWM2M_ACCESS_CONTROL},
		{oma.OBJECT_LWM2M_DEVICE},
		{oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
		{oma.OBJECT_LWM2M_LOCATION},
		{oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
	}

	for _, c := range cases3 {
		o := cli.GetObject(c.in)
		assert.NotNil(t, o, "Error getting object: ", c)
		assert.NotNil(t, o.GetEnabler(), "Error getting object enabler: ", c)
	}

	cli.AddObjectInstances(oma.OBJECT_LWM2M_SECURITY, 0, 1, 2)

	assert.Equal(t, len(cli.GetObject(oma.OBJECT_LWM2M_SECURITY).GetInstances()), 3)
}