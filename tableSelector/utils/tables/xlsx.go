package tables

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// Sheet --> 資料內容(map)
type ExcelData map[string]DataMap

// 數據名稱(first column) --> 數據內容([]string)
type DataMap map[string][]string

// fileName: "path/XXX.xlsx"
// map: Sheet名稱 --> 檔案內容(map)
func GetExcelData(fileName string) (excelData ExcelData, err error) {
	excelData = map[string]DataMap{}
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Excel讀檔失敗:%v", err)
	}

	for _, sheet := range f.GetSheetList() {
		dataMap := make(map[string][]string)
		if _, ok := excelData[sheet]; !ok {
			excelData[sheet] = make(map[string][]string)
		}

		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, fmt.Errorf("Excel讀取資料失敗:%v", err)
		}

		for _, row := range rows {
			// 将第一column的值作为Key
			key := row[0]
			dataMap[key] = make([]string, 0)
			// 将Key对应的后续列的数据存储到Map中
			for _, value := range row[1:] {
				dataMap[key] = append(dataMap[key], value)
			}
		}
		excelData[sheet] = dataMap
	}
	return
}
