package main

import (
	"sync"

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

// SetTableMaps
// func main() {
// 	var wg sync.WaitGroup
// 	for _, file := range files {
// 		fmt.Printf("上傳%#v\n", file)
// 		configPosition := filepath.Join(getCurrDirPath(), file)
// 		wg.Add(1)
// 		go func(f string) {
// 			UpdateTables(configPosition, &wg)
// 		}(file)
// 		wg.Add(1)
// 		go func(f string) {
// 			UpdateTableMap(configPosition, &wg)
// 		}(file)
// 	}
// 	wg.Wait()
// }

func UpdateTableMap(targetDir string, wg *sync.WaitGroup) {
	defer wg.Done()
	tableMaps, err := template.GetTableMapFromCSV(targetDir)
	if err != nil {
		LogTool.LogErrorf("GetTableMapFromCSV", "%#v", err)
	}
	tMap := map[string]template.TableMap{
		"HighPrTable":   tableMaps.HighPrTableMap,
		"NormalPrTable": tableMaps.NormalPrTableMap,
		"LowPrTable":    tableMaps.LowPrTableMap,
	}
	err = m.SetTableMapToRedis(gameCode, tMap, wg)
	if err != nil {
		LogTool.LogErrorf("SetTable", "%#v", err)
	}
}
