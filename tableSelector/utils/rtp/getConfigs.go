package config

import (
	"encoding/csv"
	"os"
	"strings"

	LogTool "github.com/adimax2953/log-tool"
	"gitlab.baifu-tech.net/v3/slot-table/utils/readFiles"

	"gopkg.in/yaml.v3"
)

func GetRTPCfgFromYaml(targetDir string) (rtp_cfgs RTPConfigs, err error) {
	// var rtp_cfgs RTPConfigs
	f := strings.ReplaceAll(targetDir+"\\rtp_config.yaml", "\\", "/")
	b, err := os.ReadFile(f)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &rtp_cfgs)
	if err != nil {
		return
	}
	for name, lc := range rtp_cfgs.LimitConfigs {
		err = lc.ValueVLD()
		if err != nil {
			LogTool.LogFatalf("getRTPCfg", "模板:%v 不合規範:%#v", name, err)
			return
		}
	}
	return
}

func GetGameListFromCSV(targetDir string) (TemplateKeyMap, error) {
	var gameList TemplateKeyMap
	file := strings.ReplaceAll(targetDir+"\\gameList.csv", "\\", "/")
	f, err := os.Open(file)
	if err != nil {
		return gameList, err
	}
	// 建立 CSV Reader
	reader := csv.NewReader(f)
	// 讀取所有記錄
	records, err := reader.ReadAll()
	if err != nil {
		return gameList, err
	}

	//records: Country Platform	Vendor	GameCode	RoomType	Bet	TmpID
	tmpMap := make(map[string]map[string]interface{})
	for _, record := range records {
		// Country:Platform:Vendor
		venderKey := record[readFiles.CountryName] + ":" + record[readFiles.PlatformName] + ":" + record[readFiles.VendorName]
		//	                          RoomType_Bet
		if _, ok := tmpMap[venderKey]; !ok {
			tmpMap[venderKey] = make(map[string]interface{})
		}
		tmpMap[venderKey][record[readFiles.Bet]] = record[readFiles.TmpName]
	}
	gameList = tmpMap
	return gameList, nil
}

func GetDesignatedListFromCSV(targetDir string) (TemplateKeyMap, error) {
	var gameList TemplateKeyMap
	file := strings.ReplaceAll(targetDir+"\\designatedList.csv", "\\", "/")
	f, err := os.Open(file)
	if err != nil {
		return gameList, err
	}
	// 建立 CSV Reader
	reader := csv.NewReader(f)
	// 讀取所有記錄
	records, err := reader.ReadAll()
	if err != nil {
		return gameList, err
	}

	//records: Country Platform	Vendor	GameCode	RoomType	Bet	TmpID
	tmpMap := make(map[string]map[string]interface{})
	for _, record := range records {
		// Country:Platform:Vendor
		designatedKey := record[0] + ":" + record[1]
		//	                          RoomType_Bet
		if _, ok := tmpMap[designatedKey]; !ok {
			tmpMap[designatedKey] = make(map[string]interface{})
		}
		tmpMap[designatedKey][record[2]] = record[3]
	}
	gameList = tmpMap
	return gameList, nil
}
