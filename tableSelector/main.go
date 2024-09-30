package main

import (
	"fmt"
	"path/filepath"
	"sync"

	LogTool "github.com/adimax2953/log-tool"
)

func main() {
	var wg sync.WaitGroup
	for _, file := range files {
		fmt.Printf("上傳%#v\n", file)
		err = gameCodeCheck(file, gameCode)
		if err != nil {
			LogTool.LogErrorf("", "%#v", err)
			break
		}
		configPosition := filepath.Join(getCurrDirPath(), file)
		// parentDir := filepath.Dir(configPosition)

		wg.Add(1)
		go func(f string) {
			UpdateLimitConfigs(configPosition, &wg)
		}(file)

		wg.Add(1)
		go func(f string) {
			UpdateDesignatedList(configPosition, &wg)
		}(file)

		//wg.Wait()???

		// 會檢查模板是否存在,有新模板要上傳時不能開goroutine同步進行
		wg.Add(1)
		go func(f string) {
			UpdateGameList(configPosition, &wg)
		}(file)

		// 會檢查模板是否存在,有新模板要上傳時不能開goroutine同步進行
		wg.Add(1)
		go func(f string) {
			UpdateTableMap(configPosition, &wg)
		}(file)
	}
	wg.Wait()
}

func gameCodeCheck(file, gameCode string) error {
	fileGameCode := file[:len(gameCode)]
	if fileGameCode != gameCode {
		return fmt.Errorf("gameCode錯誤! file: %v, GameCode: %v", file, gameCode)
	}
	return nil
}
