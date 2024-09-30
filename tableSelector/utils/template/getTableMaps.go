package template

import (
	"encoding/csv"
	"os"
	"strings"

	"gitlab.baifu-tech.net/v3/slot-table/utils/readFiles"
)

type TableMap map[string]interface{}
type TableMaps struct {
	HighPrTableMap   TableMap
	NormalPrTableMap TableMap
	LowPrTableMap    TableMap
}

func GetTableMapFromCSV(targetDir string) (tableMaps TableMaps, err error) {
	tableMaps.HighPrTableMap, tableMaps.NormalPrTableMap, tableMaps.LowPrTableMap = make(TableMap), make(TableMap), make(TableMap)
	file := strings.ReplaceAll(targetDir+"\\gameList.csv", "\\", "/")
	f, err := os.Open(file)
	if err != nil {
		return
	}
	// 建立 CSV Reader
	reader := csv.NewReader(f)
	// 讀取所有記錄
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	//records: Country Platform	Vendor	GameCode	RoomType	Bet	TmpID payTableName  deadTableName
	for _, record := range records {
		if record[readFiles.TmpName] == "TmpName" {
			continue
		}
		tmpKey := record[readFiles.CountryName] + ":" + record[readFiles.PlatformName] + ":" + record[readFiles.VendorName] + ":" + record[readFiles.Bet]

		// 或是填0, 不過這樣會繼不少沒用的資料...
		// 順序寫死的問題是: 如果不同遊戲有不同的表順序會有問題,變得會需要填0來維持位置
		// if hasKthElement(record, 7) {
		tableMaps.HighPrTableMap[tmpKey] = record[readFiles.HighTableName]
		// }
		tableMaps.NormalPrTableMap[tmpKey] = record[readFiles.NormalTableName]
		tableMaps.LowPrTableMap[tmpKey] = record[readFiles.LowTableName]
	}
	return
}

func hasKthElement(slice []string, k int) bool {
	// 检查切片长度是否大于 k
	return k < len(slice)
}
