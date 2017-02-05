package hu_checker

import (
	"io/ioutil"
	"encoding/json"
	"sort"
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

type HuConfigList struct {
	HuConfigLst 	[]*HuConfig `json:"hu_config_lst"`
}

func NewHuConfigList() *HuConfigList {
	return &HuConfigList{
		HuConfigLst : make([]*HuConfig, 0),
	}
}

func (confLst *HuConfigList) Less(i, j int) bool {
	if confLst.HuConfigLst[i].Score > confLst.HuConfigLst[j].Score {
		return false
	}
	return true
}

func (confLst *HuConfigList) Len() int {
	return len(confLst.HuConfigLst)
}

func (confLst *HuConfigList) Swap(i, j int) {
	tmp := confLst.HuConfigLst[i]
	confLst.HuConfigLst[i] = confLst.HuConfigLst[j]
	confLst.HuConfigLst[j] = tmp
}

func (confLst *HuConfigList) Init(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, confLst)
	if err != nil {
		return err
	}

	sort.Sort(sort.Reverse(confLst))
	return nil
}