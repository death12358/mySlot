package main

import (
	"sync"

	rtp "gitlab.baifu-tech.net/v3/slot-table/utils/rtp"

	LogTool "github.com/adimax2953/log-tool"
	"gitlab.baifu-tech.net/v3/slot-table/utils/template"
)

func init() {
	readenv()
	m, err = template.NewTmpManager(redisHost, redisPassword, redisPort, poolSize)
	if err != nil {
		if err != nil {
			LogTool.LogErrorf("NewTmpManager", "%v", err)
		}
	}
}

// SetRTP
// func main() {
// 	var wg sync.WaitGroup
// 	for _, file := range files {
// 		fmt.Printf("上傳%#v\n", file)
// 		configPosition := filepath.Join(getCurrDirPath(), file)
// 		wg.Add(1)
// 		go func(f string) {
// 			UpdateLimitConfigs(configPosition, &wg)
// 		}(file)
// 		wg.Add(1)
// 		go func(f string) {
// 			UpdateGameList(configPosition, &wg)
// 		}(file)
// 	}
// 	wg.Wait()
// }

func UpdateLimitConfigs(targetDir string, wg *sync.WaitGroup) {
	m.SetGameCodeList(gameCode)
	// configManage, err := NewRTPConfigManager(host, password, port) // localhost
	r, err := rtp.GetRTPCfgFromYaml(targetDir)
	if err != nil {
		LogTool.LogErrorf("GetRTPCfgFromYaml", "%#v", err)
	}
	err = m.SetLimitConfigs(gameCode, r.LimitConfigs, wg)
	if err != nil {
		LogTool.LogErrorf("SetLimitConfig", "%#v", err)
	}
	wg.Done()
}

func UpdateGameList(targetDir string, wg *sync.WaitGroup) {
	gl, err := rtp.GetGameListFromCSV(targetDir)
	if err != nil {
		LogTool.LogErrorf("SetGameList", "%#v", err)
	}
	// LogTool.LogInfof("", "%#v", gl)
	wg.Add(1)
	go func(c string, l rtp.TemplateKeyMap, w *sync.WaitGroup) {
		err = m.SetGameList(c, l, w)
		if err != nil {
			LogTool.LogErrorf("SetGameList", "%#v", err)
		}
	}(gameCode, gl, wg)
	wg.Done()
}

func UpdateDesignatedList(targetDir string, wg *sync.WaitGroup) {
	m.DelHashAll(template.DESIGNATED_KEY)
	gl, err := rtp.GetDesignatedListFromCSV(targetDir)
	if err != nil {
		LogTool.LogErrorf("SetDesignatedList", "%#v", err)
	}
	// LogTool.LogInfof("", "%#v", gl)
	wg.Add(1)
	go func(c string, l rtp.TemplateKeyMap, w *sync.WaitGroup) {
		err = m.SetDesignatedList(c, l, w)
		if err != nil {
			LogTool.LogErrorf("SetGameList", "%#v", err)
		}
	}(gameCode, gl, wg)
	wg.Done()
}
