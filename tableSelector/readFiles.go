package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	gotool "github.com/adimax2953/go-tool"
	LogTool "github.com/adimax2953/log-tool"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
	"gitlab.baifu-tech.net/v3/slot-table/utils/template"
)

var (
	redisHost     string
	redisPassword string
	redisPort     int
	poolSize      int = 300
	gameCode      string
	redisPort_int int
	err           error
	files         []string
	wg            sync.WaitGroup
	mlock         sync.Locker
	m             *template.TmpManager
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

// 讀取 .env 檔案
func readenv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("無法讀取 .env 檔案:", err)
		return
	}

	// 讀取環境變數
	redisHost = os.Getenv("REDIS_HOST")
	redisPassword = os.Getenv("REDIS_PASSWORD")
	redisPort, err = strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		LogTool.LogErrorf("readenv()", "%#v", err)
	}
	gameCode = os.Getenv("GAMECODE")
	files = strings.Split(os.Getenv("FILES"), ",")
}
func getCurrDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		LogTool.LogError("getCurrDirPath", err.Error())
	}
	return dir
}

func envToInt32Arr(s string) []int32 {
	// 載入 .env 檔案
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}

	// 讀取 INTEGER_ARRAY 變數
	integerArrayStr := os.Getenv(s)
	if integerArrayStr == "" {
		fmt.Printf("No %s  found in .env file\n", s)
		return nil
	}

	// 將字串拆分為切片
	integerArrayStrSlice := strings.Split(integerArrayStr, ",")

	// 轉換為整數切片
	integerArray := make([]int32, len(integerArrayStrSlice))
	for i, s := range integerArrayStrSlice {
		integer, err := gotool.StrToInt32(s)
		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", s, err)
			return nil
		}
		integerArray[i] = integer
	}

	return integerArray
}

func readExcel(DirPath string) map[string][]string {
	var dataMap = make(map[string][]string)
	// 打開 Excel 文件
	f, err := excelize.OpenFile(DirPath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	// 獲取第一個工作表的名稱
	sheetName := f.GetSheetName(f.GetActiveSheetIndex())

	// 獲取工作表中的所有行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
	}

	// 迭代行並打印每個單元格的值
	for _, row := range rows {
		dataMap[row[0]] = row[1:]
	}
	return dataMap
}
