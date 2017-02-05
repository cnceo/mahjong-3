package hu_checker

import (
	"strings"
	//"mahjong/game_server/log"
	"mahjong/game_server/log"
)

type NewCheckerFunc func(conf *HuConfig) Checker

type Factory struct {
	newCheckerFunc 	map[string]NewCheckerFunc
	huConfigLst 	*HuConfigList
	allChecker		[]Checker
}

var factoryInst *Factory
func FactoryInst() *Factory{
	if factoryInst == nil {
		factoryInst = &Factory{
			newCheckerFunc:		make(map[string]NewCheckerFunc),
			huConfigLst:		NewHuConfigList(),
			allChecker:			make([]Checker, 0),
		}
	}
	return factoryInst
}

func (factory *Factory) GetAllChecker() []Checker {
	return factory.allChecker
}

func (factory *Factory) DebugAllChecker() {
	for _, checker := range factory.allChecker {
		log.Debug("factory's checker :", checker, checker.GetConfig())
	}
}

func (factory *Factory) createChecker(conf *HuConfig) Checker {
	name := strings.ToUpper(conf.Name)
	newCheckerFunc, ok := factory.newCheckerFunc[name]
	if !ok {
		panic("hu checker not support :" + name)
		return nil
	}
	return newCheckerFunc(conf)
}

func (factory *Factory) Init(conf string) error {
	if factory.huConfigLst.Len() > 0 {
		return nil
	}
	err := factory.huConfigLst.Init(conf)
	//log.Debug("after conf lst init")
	//factory.DebugAllChecker()
	for _, conf := range factory.huConfigLst.HuConfigLst {
		if !conf.IsEnabled {
			continue
		}
		checker := factory.createChecker(conf)
		if checker == nil {
			continue
		}
		factory.allChecker = append(factory.allChecker, checker)
	}
	return err
}

func (factory *Factory) register(name string, checkerFunc NewCheckerFunc) {
	name = strings.ToUpper(name)
	factory.newCheckerFunc[name] = checkerFunc
}

