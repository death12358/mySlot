package readFiles

import (
	"github.com/adimax2953/Shared/utilities"
)

const (
	dataKeyIndex = 0
	// gameDir      = "gametemplate"
)

// CreateDataMap 建一個讀檔格式(橫式的csv檔,機率表...)
func CreateDataMap(path string) (DataMap, error) {
	records, err := utilities.OpenCSV(path)
	if err != nil {
		return nil, err
	}
	dataMap := make(DataMap)
	//逐行處理記錄
	for _, record := range records {
		// 逐欄位讀取記錄中的值,整理資料
		var k string
		for r := 0; r < len(record); r++ {
			record[r] = record[r]
		}

		if len(record) > dataKeyIndex {
			k = record[dataKeyIndex]
		}
		for i := dataKeyIndex + 1; i < len(record); i++ {
			if record[i] == "" {
				dataMap[k] = record[dataKeyIndex+1 : i]
				break
			} else if i == len(record)-1 {
				dataMap[k] = record[dataKeyIndex+1:]
				break
			}
		}
	}
	return dataMap, nil
}

// GetDataByKey 取該row
func (rm DataMap) GetDataByKey(inKey string) ([]string, bool) {
	data, ok := rm[inKey]
	return data, ok
}

// GetDataByIndex 取該row index裡的資料
func (rm DataMap) GetDataByIndex(inKey string, index int) (string, bool) {
	data, ok := rm[inKey]

	if ok && len(data)-1 >= index {
		return data[index], ok
	}
	return "", ok
}

type ProbabilitySetting struct {
	ID   int
	File string
}
