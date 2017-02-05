package hu_checker

import (
	"testing"
	"github.com/bmizerany/assert"
)

func TestHuConfig_Init(t *testing.T) {
	confLst := NewHuConfigList()
	err := confLst.Init("./hu_config.json")
	assert.Equal(t, err, nil)

	length := len(confLst.HuConfigLst)
	for idx := 0; idx < length; idx++ {
		conf := confLst.HuConfigLst[idx]
		if idx+1 < length-1 {
			tmp := confLst.HuConfigLst[idx+1]
			assert.Equal(t, conf.Score >= tmp.Score, true)
		}
	}
}
