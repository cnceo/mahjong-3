package hu_checker

import (
	"io/ioutil"
	"encoding/json"
)

type HuConfig struct {
	Name 		string      `json:"name"`		//胡的名字
	Desc		string      `json:"desc"`		//胡的中文名字
	Score		int			`json:"score"`		//胡所得分数
	IsEnabled	bool        `json:"is_enabled"`	//是否激活
}

func (config *HuConfig) ToString() string  {
	bytes, _ := json.Marshal(config)
	return string(bytes)
}

type HuConfigMap struct {
	data 		map[string]*HuConfig
}

func NewHuConfigMap() *HuConfigMap {
	return &HuConfigMap{
		data :  make(map[string]*HuConfig),
	}
}

func (configMap *HuConfigMap) Init(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	confLst := &struct {
		Hu []*HuConfig `json:"hu"`
	}{
		Hu : make([]*HuConfig, 0),
	}
	err = json.Unmarshal(bytes, confLst)
	if err != nil {
		return err
	}

	for _, conf := range confLst.Hu {
		configMap.data[conf.Name] = conf
	}
	return nil
}

func (configMap *HuConfigMap) GetHuConfig(name string) (*HuConfig, bool){
	value, ok := configMap.data[name]
	return value, ok
}

func (configMap *HuConfigMap) ToString() string {
	str := ""
	for _, value := range  configMap.data {
		str += value.ToString() + "\n"
	}
	return str
}