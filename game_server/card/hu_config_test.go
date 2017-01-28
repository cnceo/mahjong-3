package card

import (
	"testing"
	"github.com/bmizerany/assert"
	"io/ioutil"
	"strings"
)

func TestHuConfigMap_Init(t *testing.T) {
	confMap := NewHuConfigMap()
	err := confMap.Init("./hu_config.json")
	assert.Equal(t, err, nil)

	bytes, err := ioutil.ReadFile("./hu_config.json")
	content := string(bytes)
	for key, value := range confMap.config {
		idx := strings.Index(content, key)
		assert.Equal(t, idx != -1, true)
		assert.Equal(t, value.IsEnabled, true)
	}

	conf, ok := confMap.GetHuConfig("Q1S_HU")
	assert.Equal(t, conf.Name, "Q1S_HU")
	assert.Equal(t, ok, true)

}
