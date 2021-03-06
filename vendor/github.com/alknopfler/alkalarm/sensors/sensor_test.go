package sensors

import (
	"testing"
	"github.com/stretchr/testify/assert"
	cfg "github.com/alknopfler/alkalarm/config"
)

func TestRegisterSensor(t *testing.T) {
	data:=cfg.Sensor{
		Code:"testCode",
		TypeOf: "presence",
		Zone: "ventanaTest"}
	err:=Register(data)
	assert.NoError(t,err)
}

func TestQuerySensors(t *testing.T) {
	code,err:=Query("testCode")
	assert.NoError(t,err)
	assert.Equal(t,code,cfg.Sensor{Code:"testCode", TypeOf:"presence", Zone:"ventanaTest"})
}

func TestQuerySensorsAll(t *testing.T) {
	_,err:=QueryAll()
	assert.NoError(t,err)
}

func TestSensorExists(t *testing.T) {
	res:=Exists("testCode")
	assert.True(t,res)
	res2:=Exists("codeeeeeefail")
	assert.False(t,res2)
}

func TestUnregisterSensor(t *testing.T) {
	err:=Unregister("testCode")
	assert.NoError(t,err)
}