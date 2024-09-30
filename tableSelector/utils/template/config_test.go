package template_test

// // Todo檢查Redis的部分
// import (
// 	"sync"
// 	"testing"

// 	rtp "gitlab.baifu-tech.net/v3/slot-table/rtp"

// 	LogTool "github.com/adimax2953/log-tool"
// )

// var wg sync.WaitGroup
// var err error

// var (
// 	mlock sync.Locker

// 	host, password, port, poolSize = "127.0.0.1", "", 6379, 50
// )

// func init() {
// 	m, err = NewTmpManager(host, password, port, poolSize)
// 	if err != nil {
// 		if err != nil {
// 			LogTool.LogErrorf("NewTmpManager", "%v", err)
// 		}
// 	}
// }

// func Test_SetLimitConfigsFromyaml(t *testing.T) {
// 	m.SetGameCodeList("302")
// 	// configManage, err := NewRTPConfigManager(host, password, port) // localhost
// 	targetDir := "D:\\Golang\\src\\gitlab.baifu-tech.net\\v3\\slot-table\\302\\dev\\DSG"
// 	r, err := rtp.GetRTPCfgFromYaml(targetDir)
// 	if err != nil {
// 		LogTool.LogErrorf("GetRTPCfgFromYaml", "%#v", err)
// 	}
// 	err = m.SetLimitConfigs("302", r.LimitConfigs, &wg)
// 	if err != nil {
// 		LogTool.LogErrorf("SetLimitConfig", "%#v", err)
// 	}
// 	wg.Wait()
// }

// // func Test_SetGameListFromCSV(t *testing.T) {
// // 	// configManage, err := NewRTPConfigManager(host, password, port) // localhost
// // 	targetDir := "D:\\Golang\\src\\gitlab.baifu-tech.net\\v3\\slot-table\\302\\dev\\DSG"
// // 	gl, err := rtp.GetGameListFromCSV(targetDir)
// // 	if err != nil {
// // 		LogTool.LogErrorf("SetGameList", "%#v", err)
// // 	}
// // 	// LogTool.LogInfof("", "%#v", gl)
// // 	err = m.SetGameList("302", gl)
// // 	if err != nil {
// // 		LogTool.LogErrorf("SetGameList", "%#v", err)
// // 	}
// // }

// // func Test_SetRTPConfig(t *testing.T) {
// // 	// configManage, err := NewRTPConfigManager(host, password, port) // localhost
// // 	targetDir := "D:\\Golang\\github\\bftrtpmodel\\slotProb\\slotGame"
// // 	path.SetRTPPath(targetDir)

// // 	GetGameInfoFromYaml()
// // 	configManage, err := NewRTPConfigManager(host, password, port)
// // 	if err != nil {
// // 		LogTool.LogErrorf("NewRTPConfigManager", "%#v", err)
// // 	}
// // 	for roomKey, gi := range gameInfoData.GameInfo {
// // 		gc := GameConfig{
// // 			GameInfo:       gi,
// // 		}
// // 		lc := rtpCfgData.LimitConfig[roomKey]
// // 		rc.GameConfigs = append(rc.GameConfigs, gc)

// // 		rc.RTPConfigs[roomKey] = RTPConfig{
// // 			LimitConfig: lc,
// // 		}
// // 	}
// // 	configManage.SetRTPConfig(config)
// // }

// // func Test_AutoWriteRedisData2(t *testing.T) {
// // 	var (
// // 		// roomTypeID = []int32{1, 2, 3}
// // 		// modelID  = []int32{1, 2, 3}
// // 		gameCode               = "302"
// // 		roomTypesMappingProdID = map[int32]int{
// // 			1: 1,
// // 			2: 1,
// // 			3: 1,
// // 		}
// // 	)
// // 	var (
// // 		countryID, platformID, vendorID                               int32 = 1, 1, 15
// // 		countryName, platformName, vendorName, gameName, roomTypeName       = "country", "platform", "vendor", "game", " roomTypeName"
// // 		rc                                                            GameRTPConfig
// // 	)

// // 	//configManage, err := NewRTPConfigManager(host, password, port)
// // 	configManage, err := NewRTPConfigManager(host, password, port) //dev
// // 	// configManage, err := NewRTPConfigManager(host, password, port) // localhost
// // 	if err != nil {
// // 		fmt.Println(err.Error())
// // 		return
// // 	}
// // 	rc.RTPConfigs = make(map[string]RTPConfig)
// // 	rc.GameConfigs = make([]GameConfig, 0)
// // 	// var configByRoomType = map[int32][]interface{}{
// // 	// 	//BaseBet,SRTP,     DSLoss,           DPrProfit,    MPrProfit,  SCfg        PrCfg
// // 	// 	1: {10, true, 9950, false, 1500, 99, true, 1500, true, 20000, 99, 9650, 1},
// // 	// 	2: {100, true, 9950, false, 3000, 99, true, 3000, true, 40000, 99, 9650, 1},
// // 	// 	3: {150, true, 9950, false, 4500, 99, true, 4500, true, 60000, 99, 9650, 1},
// // 	// 	// 4:    {200, true, 9950, false, 1, 99, true, 600000, true, 800000, 99, 9650, 1, 0, 10, 200, 2, 2500, 3},
// // 	// 	// 5:    {250, true, 9950, false, 1, 99, true, 750000, true, 1000000, 99, 9650, 1, 0, 10, 200, 2, 2500, 3},
// // 	// 	// 10:   {500, true, 9895, false, 1, 99, true, 1000000, true, 1500000, 99, 9645, 1, 0, 10, 125, 2, 2500, 3},
// // 	// 	// 20:   {1000, true, 9870, false, 1, 99, true, 1000000, true, 1500000, 99, 9640, 1, 0, 10, 115, 2, 2500, 3},
// // 	// 	// 30:   {1500, true, 9855, false, 1, 99, true, 1000000, true, 1500000, 99, 9635, 1, 0, 10, 110, 2, 2500, 3},
// // 	// 	// 40:   {2000, true, 9830, false, 1, 99, true, 1000000, true, 1500000, 99, 9630, 1, 0, 10, 100, 2, 2500, 3},
// // 	// 	// 50:   {2500, true, 9805, false, 1, 99, true, 1000000, true, 1500000, 99, 9625, 1, 0, 10, 90, 2, 2500, 3},
// // 	// 	// 100:  {5000, true, 9760, false, 1, 99, true, 1000000, true, 1500000, 99, 9620, 1, 0, 10, 70, 2, 2500, 3},
// // 	// 	// 150:  {7500, true, 9715, false, 1, 99, true, 1000000, true, 1500000, 99, 9615, 1, 0, 10, 50, 2, 2500, 3},
// // 	// 	// 300:  {15000, true, 9710, false, 1, 99, true, 1000000, true, 1500000, 99, 9610, 1, 0, 10, 50, 2, 2500, 3},
// // 	// 	// 500:  {25000, true, 9705, false, 1, 99, true, 1250000, true, 2000000, 99, 9605, 1, 0, 10, 50, 2, 2500, 3},
// // 	// 	// 1000: {50000, true, 9700, false, 1, 99, true, 2000000, true, 3000000, 99, 9600, 1, 0, 10, 50, 2, 2500, 3},
// // 	// 	// 2000: {100000, true, 9695, false, 1, 99, true, 3000000, true, 4000000, 99, 9595, 1, 0, 10, 50, 2, 2500, 3},
// // 	// }
// // 	for roomType, _ := range roomTypesMappingProdID {
// // 		configKey := fmt.Sprintf("%d_%d_%d_%s_%d", countryID, platformID, vendorID, gameCode, roomType)
// // 		config := ConfigByKey[configKey]
// // 		gc := GameConfig{
// // 			GameInfo: GameInfo{
// // 				CountryID:    countryID,
// // 				PlatformID:   platformID,
// // 				VendorID:     vendorID,
// // 				GameCode:     gameCode,
// // 				RoomType:     roomType,
// // 				GameName:     gameName,
// // 				PlatformName: platformName,
// // 				VendorName:   vendorName,
// // 				CountryName:  countryName,
// // 				RoomTypeName: roomTypeName,
// // 			},
// // 			ConfigKey: configKey,
// // 		}
// // 		rc.GameConfigs = append(rc.GameConfigs, gc)
// // 		rc.RTPConfigs[configKey] = RTPConfig{
// // 			LimitConfig: LimitConfig{
// // 				BaseBet:                         int64(config[0].(int)),
// // 				SysRTPLimitEnabled:              config[1].(bool),
// // 				SysRTPLimit:                     int32(config[2].(int)),
// // 				DailySysLossLimitEnabled:        config[3].(bool),
// // 				DailySysLossLimit:               int64(config[4].(int)),
// // 				DailyPlayerProfitLimitEnabled:   config[6].(bool),
// // 				DailyPlayerProfitLimit:          int64(config[7].(int)),
// // 				MonthlyPlayerProfitLimitEnabled: config[8].(bool),
// // 				MonthlyPlayerProfitLimit:        int64(config[9].(int)),
// // 			},
// // 			SysConfig: SysConfig{
// // 				ExpectedRTP: int32(config[11].(int)),
// // 				BaseProb:    config[12].(int),
// // 			},
// // 			PlayerConfig: PlayerConfig{
// // 				ExpectedRTP: 1,
// // 				Enabled:     false,
// // 			},
// // 		}
// // 	}
// // 	// token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50Ijoic3VwZXJhZG1pbiIsImV4cCI6MTcwMDIyMTc2MiwiaWF0IjoxNzAwMTg1NzYyLCJpc3MiOiJhdXRoLXNlcnZpY2UiLCJsYW5ndWFnZUlkIjoxLCJwbGF0Zm9ybXMiOlsxLDNdLCJwcmVmZXJlbmNlQ291bnRyeUlkIjowLCJyb2xlcyI6bnVsbCwidHlwZSI6MSwidWlkIjoxfQ.vYPJkTaDHoVu_ySwcb77Pdf1UGZHhDShzt82rj0MVdI`

// // 	// for i, s := range roomTypeID {
// // 	// 	l, _ := json.Marshal(rc.RTPConfigs[int(s)].LimitConfig)
// // 	// 	send(token, modelID[i], "limit", l)
// // 	// 	sc, _ := json.Marshal(rc.RTPConfigs[int(s)].SysConfig)
// // 	// 	send(token, modelID[i], "system", sc)
// // 	// 	pc, _ := json.Marshal(rc.RTPConfigs[int(s)].PlayerConfig)
// // 	// 	send(token, modelID[i], "player", pc)
// // 	// 	npc, _ := json.Marshal(rc.RTPConfigs[int(s)].NewPlayerConfig)
// // 	// 	send(token, modelID[i], "new-player", npc)
// // 	// 	b, _ := json.Marshal(rc.RTPConfigs[int(s)].BonusConfig)
// // 	// 	send(token, modelID[i], "bonus", b)
// // 	// 	h, _ := json.Marshal(rc.RTPConfigs[int(s)].HighScoreConfig)
// // 	// 	send(token, modelID[i], "high-score", h)
// // 	// }

// // 	err = configManage.SetRTPConfig(rc)
// // 	if err != nil {
// // 		fmt.Println(err.Error())
// // 	}
// // }

// //	func send(token string, modelID int32, configName string, data []byte) {
// //		req, _ := http.NewRequest("PUT", fmt.Sprintf("http://8.222.254.57:9400/v1/admin/slot/rtp/model/%v/%v", modelID, configName), bytes.NewBuffer(data))
// //		//req, _ := http.NewRequest("PUT", fmt.Sprintf("http://103.103.81.14:9400/v1/admin/slot/rtp/model/%v/%v", modelID, configName), bytes.NewBuffer(data))
// //		req.Header.Add("Authorization", token)
// //		client := &http.Client{}
// //		response, err := client.Do(req)
// //		if err != nil {
// //			fmt.Println("HTTP请求失败:", err)
// //			return
// //		}
// //		defer response.Body.Close()
// //		responseBody, err := ioutil.ReadAll(response.Body)
// //		if err != nil {
// //			fmt.Println("无法读取响应内容:", err)
// //			return
// //		}
// //		fmt.Println(modelID, configName, string(responseBody))
// //	}
// func Test_GetConfigIDListFromCSV(t *testing.T) {
// 	targetDir := "D:\\Golang\\src\\github.com\\adimax2953\\bftrtpmodel\\slotProb\\slotGame"
// 	rtp.GetGameListFromCSV(targetDir)
// }
