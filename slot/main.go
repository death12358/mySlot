package main

import (
	LogTool "github.com/adimax2953/log-tool"
	"github.com/adimax2953/slot-math/gamelogic"
)

var dfh *dfhControl

type dfhControl struct {
	gameLogic *gamelogic.Games
}

func main() {
	// err := dfh.InitController()

	g := GetController()
	LogTool.LogInfof("", "%#v", g.Spin)
}
func GetController() *dfhControl {
	if dfh == nil {
		dfh = &dfhControl{}
		dfh.InitController()
	}
	return dfh
}

func (c *dfhControl) InitController() error {
	var opt gamelogic.RedisOption
	s := pkg_config.GetSetting()
	if s == nil {
		opt.UseLocalCache = true
	} else {
		opt = gamelogic.RedisOption{
			Host:          s.RedisMathHost,
			Password:      s.RedisMathPassword,
			Port:          s.RedisMathPort,
			PoolSize:      s.RedisMathPoolSize,
			UseLocalCache: pkg_config.IsTest(),
		}
	}
	g, err := gamelogic.NewGames(shared_game_code.DFH.String(), opt, internal.DatasDirPath)
	if err != nil {
		return err
	}

	dfh.gameLogic = g

	return nil
}

func (c *dfhControl) Spin(inRequest *pb_game.SlotGameRequest) (*pb_game.SlotGameResults, error) {
	rslt, err := c.gameLogic.Spin(inRequest)

	return rslt, err
}

// GetPay() 取賠付資訊
func (c *dfhControl) GetPay() *pb_game.PayData {
	return c.gameLogic.GetPay()
}

// GetProbabilitySetting 大富豪沒東西
func (c *dfhControl) GetProbabilitySetting() map[int]string {
	return map[int]string{}
}
