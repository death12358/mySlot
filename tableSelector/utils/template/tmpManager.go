package template

import (
	"fmt"
	"strconv"
	"sync"

	goredis "github.com/adimax2953/go-redis"
	"github.com/adimax2953/go-redis/src"
	// "github.com/adimax2953/bftrtpmodel/slotProb/slotRTP/config"
	// rtpCfg "github.com/adimax2953/bftrtpmodel/slotProb/slotRTP/config"
	// LogTool "github.com/adimax2953/log-tool"
)

var (
	once sync.Once
)

type TmpManager struct {
	myScriptor *src.MyScriptor
}

var m *TmpManager

func NewTmpManager(host, password string, port, poolSize int) (*TmpManager, error) {
	var err error
	DBKEY_SCRIPT_int, _ := strconv.Atoi(DBKEY_SCRIPT)
	once.Do(func() {
		m = &TmpManager{}
		opt := &goredis.Option{
			Host:     host,
			Port:     port,
			Password: password,
			DB:       DBKEY_SCRIPT_int,
			PoolSize: poolSize,
		}
		s := &goredis.Scriptor{}
		s, err = goredis.NewDB(opt, opt.DB, "Bft|0.0.1", &src.LuaScripts)
		if err != nil {
			return
		}
		m.myScriptor = &src.MyScriptor{
			Scriptor: s,
		}
	})
	return m, err
}

func (m *TmpManager) SetTableList(gameCode string, tableName string, tableKeyType string, wg *sync.WaitGroup) error {
	defer wg.Done()
	keys := []string{DBKEY_SLOTTABLE, SLOT_TABLE_KEY + ":" + gameCode, tableKeyType}
	var tables []string
	res, err := m.myScriptor.GetListAll(keys, []string{TABLE_LIST_KEY})
	if err != nil {
		return err
	}
	for _, r := range *res {
		tables = append(tables, r.Value)
	}
	exist := false
	for _, ti := range tables {
		if ti == tableName {
			exist = true
			break
		}
	}
	if !exist {
		_, err = m.myScriptor.NewList(keys, []string{TABLE_LIST_KEY, "R", tableName})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m TmpManager) SetTableMap(gameCode string, tableMap TableMap, tableKeyType string, wg *sync.WaitGroup) error {
	tableNameChecked := make(map[interface{}]bool)
	for _, tableName := range tableMap {
		if tableNameChecked[tableName] {
			continue
		}
		tableNameChecked[tableName] = true
		wg.Add(1)
		go func(gc, tn, tkType string, wg *sync.WaitGroup) error {
			err := m.CheckTableList(gameCode, tableName.(string), tableKeyType, wg)
			if err != nil {
				return err
			}
			return nil
		}(gameCode, tableName.(string), tableKeyType, wg)
	}

	keys := []string{DBKEY_SLOTTABLE, SLOT_TABLE_KEY + ":" + gameCode, tableKeyType, TABLE_MAP_KEY}
	_, err := m.myScriptor.UpdateHashBatch(keys, map[string]interface{}(tableMap))
	if err != nil {
		return err
	}
	return nil
}

func (m *TmpManager) CheckTableList(gameCode string, tableName string, tableKeyType string, wg *sync.WaitGroup) error {
	defer wg.Done()
	keys := []string{DBKEY_SLOTTABLE, SLOT_TABLE_KEY + ":" + gameCode, tableKeyType}
	var tables []string
	res, err := m.myScriptor.GetListAll(keys, []string{TABLE_LIST_KEY})
	if err != nil {
		return err
	}
	for _, r := range *res {
		tables = append(tables, r.Value)
	}
	exist := false
	for _, ti := range tables {
		if ti == tableName {
			exist = true
			break
		}
	}
	if !exist {
		return fmt.Errorf("Table %s doesn't exist", tableName)

	}
	return nil
}

// func (m *TmpManager) SetTmpConfig(gameTmpConfig *GameRTPConfig) error {
// 	var err error
// 	for tmpName, tmp := range gameTmpConfig.TemplateConfigs {
// 		tmp.TmpName = tmpName
// 		err = m.SetRTPConfig(tmp, gameTmpConfig.GameConfigs[0].GameInfo)
// 		if err != nil {
// 			LogTool.LogErrorf("m.SetRTPConfig", "%v", err)
// 		}
// 		err = m.SetTableMaps(tmp)
// 		if err != nil {
// 			LogTool.LogErrorf("m.SetTables", "%v", err)
// 		}
// 	}
// 	for _, gc := range gameTmpConfig.GameConfigs {
// 		gi := gc.GameInfo
// 		err = m.SetGameInfo(gc.GameInfo)
// 		if err != nil {
// 			LogTool.LogErrorf("m.SetRTPConfig", "%v", err)
// 		}
// 		err = m.SetTmpChoiseMap(gi, gameTmpConfig.TemplateConfigs[gc.ConfigName])
// 		if err != nil {
// 			LogTool.LogErrorf("m.SetTmpChoiseMap", "%v", err)
// 		}
// 	}

// 	// err = m.DeleteNotUseTmp()
// 	// if err != nil {
// 	// 	LogTool.LogErrorf("m.DeleteNotUseTmp()", "%v", err)
// 	// }
// 	return nil

// }

// func (m TmpManager) SetGameInfo(gi config.GameInfo) error {
// 	err := m.RtpConfigManager.SetGameInfo([]rtpCfg.GameConfig{{GameInfo: gi}})
// 	return err
// }
// func (m TmpManager) SetTableMaps(tmp *Template) error {
// 	err := m.TableManager.SetTableMapToRedis(tmp)
// 	return err
// }

// func (m TmpManager) SetTables(tmp *Template) error {
// 	err := m.TableManager.SetTableToRedis(tmp)
// 	return err
// }

// func (m TmpManager) SetRTPConfig(tmp *Template, gameInfo rtpCfg.GameInfo) error {

// 	rtpCfgs := make(map[string]rtpCfg.RTPConfig)
// 	rtpCfgs[tmp.TmpName] = rtpCfg.RTPConfig{LimitConfig: *tmp.LimitControll.LimitConfig}
// 	err := m.RtpConfigManager.SetLimitConfig(rtpCfgs, gameInfo)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (m TmpManager) SetTmpChoiseMap(gi config.GameInfo, tmp *Template) error {
// 	venderkey := fmt.Sprintf("%s:%d:%d:%d", gi.GameCode, gi.CountryID, gi.PlatformID, gi.VendorID)
// 	betKey := fmt.Sprintf("%d:%d", gi.RoomType, gi.Bet)
// 	templateKeyMap := make(map[string]map[string]interface{})
// 	templateKeyMap[venderkey][betKey] = tmp.TmpName
// 	configManage := m.RtpConfigManager
// 	_, err := configManage.SetTmpMap(gi.GameCode, templateKeyMap)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (m *TmpManager) DeleteNotUseTmp() error {
// 	err := m.DeleteNotUseRTPConfig()
// 	if err != nil {
// 		return err
// 	}
// 	gameListKeys := []string{DBKEY_SLOTRTP, SLOT_RTP_KEY, GAME_LIST_KEY}
// 	GameCodes, err := m.TableManager.myScriptor.GetListAll(gameListKeys, []string{GAME_CODE_KEY})
// 	if err != nil {
// 		return err
// 	}
// 	time.Sleep(time.Second * 3)
// 	// 針對每款遊戲(gameCode)遍歷出使用的configID, 並刪除未被config使用的tables
// 	for _, gameCodeMap := range *GameCodes {
// 		gameCodeRes := gameCodeMap.Value
// 		haveUseConfigIDList_ForResGameCode, err := m.findUseConfigIDList(gameListKeys, gameCodeRes)
// 		if err != nil {
// 			return err
// 		}
// 		err = m.DeleteNotUseTables(gameCodeRes, haveUseConfigIDList_ForResGameCode)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (m *TmpManager) findUseConfigIDList(gameListKeys []string, gameCodeRes string) ([]string, error) {
// 	var haveUseConfigIDList_ForResGameCode []string
// 	configIDs, err := m.TableManager.myScriptor.GetHashAll(gameListKeys, []string{gameCodeRes})
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, configIDMap := range *configIDs {
// 		cfgIDRes := configIDMap.Value
// 		if err != nil {
// 			return nil, err
// 		}

// 		exist := isIDinList(cfgIDRes, haveUseConfigIDList_ForResGameCode)
// 		if !exist {
// 			haveUseConfigIDList_ForResGameCode = append(haveUseConfigIDList_ForResGameCode, cfgIDRes)
// 		}
// 	}
// 	return haveUseConfigIDList_ForResGameCode, nil
// }

// func (m *TmpManager) DeleteNotUseTables(gameCodeRes string, haveUseConfigIDList_ForResGameCode []string) error {
// 	DeadTableKeys := []string{DBKEY_SLOTTABLE, SLOT_TABLE_KEY + ":" + gameCodeRes, DEAD_TABLE_KEY}
// 	PayTableKeys := []string{DBKEY_SLOTTABLE, SLOT_TABLE_KEY + ":" + gameCodeRes, PAY_TABLE_KEY}
// 	haveUseDeadTableKeyList, err := m.cleanNotUseTable_FromListOnRedis(DeadTableKeys, haveUseConfigIDList_ForResGameCode)
// 	if err != nil {
// 		return err
// 	}
// 	haveUsePayTableKeyList, err := m.cleanNotUseTable_FromListOnRedis(PayTableKeys, haveUseConfigIDList_ForResGameCode)
// 	if err != nil {
// 		return err
// 	}
// 	err = m.cleanNotUseTables(DeadTableKeys, haveUseDeadTableKeyList)
// 	if err != nil {
// 		return err
// 	}
// 	err = m.cleanNotUseTables(PayTableKeys, haveUsePayTableKeyList)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (m *TmpManager) DeleteNotUseRTPConfig() error {
// 	gameListKeys := []string{DBKEY_SLOTRTP, SLOT_RTP_KEY, GAME_LIST_KEY}
// 	GameCodes, err := m.TableManager.myScriptor.GetListAll(gameListKeys, []string{GAME_CODE_KEY})
// 	if err != nil {
// 		return err
// 	}
// 	// 遍歷所有GameCode 找出所有有在使用的TMP(configID)
// 	var haveUseConfigIDList []string
// 	for _, gameCodeMap := range *GameCodes {
// 		gameCodeRes := gameCodeMap.Value
// 		configIDs, err := m.TableManager.myScriptor.GetHashAll(gameListKeys, []string{gameCodeRes})
// 		if err != nil {
// 			return err
// 		}
// 		//     roomKey_bet --> tmpID
// 		for _, configIDMap := range *configIDs {
// 			cfgIDRes := configIDMap.Value
// 			if err != nil {
// 				return err
// 			}
// 			exist := isIDinList(cfgIDRes, haveUseConfigIDList)
// 			if !exist {
// 				haveUseConfigIDList = append(haveUseConfigIDList, cfgIDRes)
// 			}
// 		}

// 		rtpCfgKeys := []string{DBKEY_SLOTRTP, SLOT_RTP_KEY, RTP_CONFIG_KEY}
// 		// 遍歷現有ConfigID 移除未在haveUseConfigIDList上的config
// 		var cacheConfigIDList []string
// 		cacheConfigIDs, err := m.TableManager.myScriptor.GetListAll(rtpCfgKeys, []string{CONFIG_ID_KEY})
// 		if err != nil {
// 			return err
// 		}
// 		for _, r := range *cacheConfigIDs {
// 			cacheConfigIDList = append(cacheConfigIDList, r.Value)
// 		}
// 		for _, cacheConfigID := range cacheConfigIDList {
// 			exist := isIDinList(cacheConfigID, haveUseConfigIDList)
// 			if !exist {
// 				m.TableManager.myScriptor.DelList(rtpCfgKeys, []string{CONFIG_ID_KEY, "0", cacheConfigID})
// 				m.TableManager.myScriptor.DelHashAll(rtpCfgKeys, []string{cacheConfigID})
// 			}
// 		}
// 	}
// 	return nil
// }

// func (m *TmpManager) cleanNotUseTables(tableKeys []string, haveUseTableKeyList []string) error {
// 	var cacheTableKeyList []string
// 	cacheTableKeys, err := m.TableManager.myScriptor.GetListAll(tableKeys, []string{TABLE_LIST_KEY})
// 	if err != nil {
// 		return err
// 	}
// 	for _, r := range *cacheTableKeys {
// 		if err != nil {
// 			return err
// 		}
// 		cacheTableKeyList = append(cacheTableKeyList, r.Value)
// 	}
// 	for _, cacheTableKey := range cacheTableKeyList {
// 		exist := isIDinList(cacheTableKey, haveUseTableKeyList)
// 		if !exist {
// 			m.TableManager.myScriptor.DelList(tableKeys, []string{TABLE_LIST_KEY, "0", cacheTableKey})
// 			//                                     (slotTable+gameCode):Pay(or Dead)Table:TableKey:
// 			m.TableManager.DeleteTablesWithPattern(tableKeys[1] + ":" + tableKeys[2] + ":" + cacheTableKey + ":" + "*")
// 		}
// 	}
// 	return nil
// }

// func (m *TmpManager) cleanNotUseTable_FromListOnRedis(TableKeys []string, haveUseConfigIDList_ForResGameCode []string) ([]string, error) {
// 	var haveUseTableKeyList []string
// 	TableMap, err := m.TableManager.myScriptor.GetHashAll(TableKeys, []string{TABLE_MAP_KEY})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 遍歷出所有有在使用的Table(Key)
// 	var cacheCfgKeyList_ForTable []string
// 	for _, r := range *TableMap {
// 		cacheCfgKeyList_ForTable = append(cacheCfgKeyList_ForTable, r.Key)
// 		if !isIDinList(r.Value, haveUseTableKeyList) && isIDinList(r.Key, haveUseConfigIDList_ForResGameCode) {
// 			haveUseTableKeyList = append(haveUseTableKeyList, r.Value)
// 		}
// 	}
// 	// Table列表上移除未使用的TableKey
// 	for _, c := range cacheCfgKeyList_ForTable {
// 		exist := isIDinList(c, haveUseConfigIDList_ForResGameCode)
// 		if !exist {
// 			m.TableManager.myScriptor.DelHash(TableKeys, []string{TABLE_MAP_KEY, c})
// 		}
// 	}
// 	return haveUseTableKeyList, nil
// }

// func isIDinList[T int | string](id T, list []T) bool {
// 	for _, v := range list {
// 		if v == id {
// 			return true
// 		}
// 	}
// 	return false
// }
