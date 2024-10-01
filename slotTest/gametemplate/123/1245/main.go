package main

import (
	"fmt"
	"sync"

	pb_game "github.com/adimax2953/Shared/pb/slot/game"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/adimax2953/slot-math/gamelogic"
)

var pointLock sync.Mutex

var (
	csvPath                        = "D:\\Golang\\src\\gitlab.com\\bf_tech\\template\\slot\\datas"
	host, password, port, poolSize = "127.0.0.1", "", 6379, 10000
)

var opt = gamelogic.RedisOption{
	Host:          host,
	Password:      password,
	Port:          port,
	PoolSize:      poolSize,
	UseLocalCache: false,
}

var bet int32 = 10000

var mock_game gamelogic.Games
var err error

func init() {
	mock_game, err = gamelogic.NewGames("40C", opt, csvPath)
	if err != nil {
		LogTool.LogErrorf("NewGames ERR", "%v", err)
	}
}

func getPoint(i int32) int64 {
	var spinRequest = &pb_game.SlotGameRequest{
		PlayerInfo: &pb_game.PlayerInfo{
			PlayerID:   i,
			CountryID:  0,
			VendorID:   2,
			PlatformID: 4,
		},
		BetInfo: &pb_game.BetInfo{
			Bet:       bet,
			SingleBet: bet / 10,
		},
		GameRequestTriggerType: 0,
	}
	round, err := mock_game.Spin(spinRequest)
	if err != nil {
		fmt.Printf("Spin error: %s", err.Error())
	}
	// j_round, err := json.Marshal(round)
	// if err != nil {
	// 	fmt.Printf("Marshal error: %s", err.Error())
	// }

	return round.TotalWin
}
