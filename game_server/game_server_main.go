package main

import (
	"bufio"
	"os"
	"mahjong/game_server/log"
	"mahjong/game_server/playing"
	"mahjong/game_server/hu_checker"
	"strconv"
)

func help() {
	log.Debug("help info")
	log.Debug(playing.PlayerOperateDrop, " : playing.PlayerOperateDrop")
	log.Debug(playing.PlayerOperateChi, " : playing.PlayerOperateChi")
	log.Debug(playing.PlayerOperatePeng, " : playing.PlayerOperatePeng")
	log.Debug(playing.PlayerOperateGang, " : playing.PlayerOperateGang")
	log.Debug(playing.PlayerOperateZiMo, " : playing.PlayerOperateZiMo")
	log.Debug(playing.PlayerOperateDianPao, " : playing.PlayerOperateDianPao")
	//log.Debug(playing.PlayerOperateEnterRoom)
	//log.Debug(playing.PlayerOperateLeaveRoom)
}

func main() {
	running := true

	//init hu checker factory
	huFactory := hu_checker.FactoryInst()
	err := huFactory.Init("./hu_checker/hu_config.json")
	if err != nil {
		log.Debug("hu factory init", err)
		return
	}

	//init room
	conf := playing.NewRoomConfig()
	err = conf.Init("./playing/room_config.json")
	if err != nil {
		log.Debug("room config init", err)
		return
	}
	room := playing.NewRoom(conf)
	room.Start()
	allChecker := huFactory.GetAllChecker()

	robots := []*playing.Player{playing.NewPlayer(allChecker), playing.NewPlayer(allChecker), playing.NewPlayer(allChecker)}
	for _, robot := range robots {
		robot.OperateEnterRoom(room)
	}
	curPlayer := playing.NewPlayer(allChecker)
	curPlayer.OperateEnterRoom(room)

	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		cmd := string(data)
		if cmd == "h" {
			help()
		}
		c, _ := strconv.Atoi(cmd)
		switch playing.PlayerOperateType(c) {
		case playing.PlayerOperateDrop:
			//curPlayer.OperateDrop()
		case playing.PlayerOperateChi:
		case playing.PlayerOperatePeng:
		case playing.PlayerOperateGang:
		case playing.PlayerOperateZiMo:
		case playing.PlayerOperateDianPao:

		}
	}
}