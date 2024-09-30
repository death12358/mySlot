package template

import (
	"fmt"
	"sync"

	"github.com/adimax2953/go-redis/src"
	LogTool "github.com/adimax2953/log-tool"
	rtp "gitlab.baifu-tech.net/v3/slot-table/utils/rtp"
)

func (m *TmpManager) SetLimitConfigs(gameCode string, RTPConfigs map[string]rtp.LimitConfig, wg *sync.WaitGroup) error {
	for cfgName, rtpCfg := range RTPConfigs {
		wg.Add(1)
		go func(code string, name string, rtpCfg rtp.LimitConfig, waitGruup *sync.WaitGroup) error {
			defer wg.Done()
			rtpCfg_map := configToMap(rtpCfg)
			res, err := m.updateHashBatch(RTP_CONFIG_KEY+":"+code, name, rtpCfg_map)
			if err != nil {
				return err
			}
			if len(*res) != len(rtpCfg_map) {
				return fmt.Errorf("rtpCfg上傳成功的項目與yaml上的數量不符\n 分別為%v,%v", len(*res), len(rtpCfg_map))
			}
			wg.Add(1)
			go func(n string, waitGruup *sync.WaitGroup) error {
				err = m.SetConfigList(n, wg)
				if err != nil {
					return err
				}
				return nil
			}(name, wg)
			// LogTool.LogInfof("", "上傳:%#v", name)
			return nil
		}(gameCode, cfgName, rtpCfg, wg)
	}
	return nil
}

func (m *TmpManager) SetConfigList(configName string, wg *sync.WaitGroup) error {
	defer wg.Done()
	keys := []string{DBKEY_SLOTRTP, SLOT_RTP_KEY, RTP_CONFIG_KEY}
	var configIDs []string
	res, err := m.myScriptor.GetListAll(keys, []string{CONFIG_ID_KEY})
	if err != nil {
		return err
	}
	for _, r := range *res {
		configIDs = append(configIDs, r.Value)
	}

	exist := false
	for _, r := range *res {
		if r.Value == configName {
			exist = true
			break
		}
	}

	if !exist {
		_, err = m.myScriptor.NewList(keys, []string{CONFIG_ID_KEY, "R", configName})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *TmpManager) SetGameList(gameCode string, templateKeyMap map[string]map[string]interface{}, wg *sync.WaitGroup) error {
	defer wg.Done()

	// 確認目前有哪些模板
	keys := []string{DBKEY_SLOTRTP, SLOT_RTP_KEY, RTP_CONFIG_KEY}
	res, err := m.myScriptor.GetListAll(keys, []string{CONFIG_ID_KEY})
	if err != nil {
		return err
	}
	existTmp := make(map[string]bool)
	for _, r := range *res {
		existTmp[r.Value] = true
	}

	// "gameCode:CountryName:PlatformName:VendorName"   templateKeyMap->room:bet
	number_of_gameListUpdated := 0
	for vendorKey, tmpMap := range templateKeyMap {
		for _, tmpName := range tmpMap {
			if !existTmp[tmpName.(string)] && tmpName.(string) != "TmpName" {
				return fmt.Errorf("tmpName:%#v 不存在", tmpName)
			}
		}
		// go func(vendorKey string, tmpMap map[string]interface{}) error {
		results, err := m.updateHashBatch(GAME_LIST_KEY, gameCode+":"+vendorKey, tmpMap)
		if err != nil {
			LogTool.LogErrorf("configManage.SetTmpMap", "%#v", err)
			return err
		}
		number_of_gameListUpdated += len(*results)
	}
	number_of_gameListFromCSV := 0
	for _, m := range templateKeyMap {
		number_of_gameListFromCSV += len(m)
	}

	if number_of_gameListUpdated != number_of_gameListFromCSV {
		return fmt.Errorf("上傳的GameList數量與CSV上的數量不符\n 分別為%v,%v", number_of_gameListUpdated, number_of_gameListFromCSV)
	}
	return nil
}

func (m *TmpManager) SetDesignatedList(gameCode string, designatedKeyMap map[string]map[string]interface{}, wg *sync.WaitGroup) error {
	defer wg.Done()
	for designatedKey, table := range designatedKeyMap {
		_, err := m.updateHashBatch(DESIGNATED_KEY, designatedKey, table)
		if err != nil {
			LogTool.LogErrorf("configManage.SetTmpMap", "%#v", err)
			return err
		}
	}

	return nil
}

func (m *TmpManager) SetGameCodeList(gameCode string) error {
	keys := []string{DBKEY_SLOTRTP, SLOT_RTP_KEY, GAME_LIST_KEY}
	res, err := m.myScriptor.GetListAll(keys, []string{GAME_CODE_KEY})
	if err != nil {
		return err
	}
	exist := false
	for _, r := range *res {
		if r.Value == gameCode {
			exist = true
			break
		}
	}
	if !exist {
		_, err = m.myScriptor.NewList(keys, []string{GAME_CODE_KEY, "R", gameCode})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *TmpManager) updateHashBatch(tagKey, mainKey string, sysArgs map[string]interface{}) (*[]src.RedisResult, error) {
	keys := []string{DBKEY_SLOTRTP, SLOT_RTP_KEY, tagKey, mainKey}
	results, err := m.myScriptor.UpdateHashBatch(keys, sysArgs)
	return results, err
}

// 於SlotRTP中, 刪除"SlotRTP:tagKey"前墜的所有hashMap
func (m *TmpManager) DelHashAll(tagKey string) {
	keys := []string{DBKEY_SLOTRTP, SLOT_RTP_KEY + ":" + tagKey}
	res, err := m.myScriptor.ScanMatchKeys(keys, []string{})
	if err != nil {
		LogTool.LogErrorf("", "%#v", err)
	}
	keys = []string{DBKEY_SLOTRTP, SLOT_RTP_KEY, tagKey}

	for _, v := range *res {
		m.myScriptor.DelHashAll(keys, []string{v.Key[len(SLOT_RTP_KEY+":"+tagKey+":"):]})
	}
}
func configToMap(lc rtp.LimitConfig) map[string]interface{} {
	data := make(map[string]interface{})
	data[rtp.SlotBaseBet] = lc.BaseBet
	data[rtp.SlotSysRTPLimitEnabled] = lc.SysRTPLimitEnabled
	data[rtp.SlotSysRTPLimit] = lc.SysRTPLimit
	data[rtp.SlotMonthlySysLossLimitEnabled] = lc.MonthlySysLossLimitEnabled
	data[rtp.SlotMonthlySysLossLimit] = lc.MonthlySysLossLimit
	data[rtp.SlotDailySysLossLimitEnabled] = lc.DailySysLossLimitEnabled
	data[rtp.SlotDailySysLossLimit] = lc.DailySysLossLimit
	data[rtp.SlotDailyPlayerProfitLimitEnabled] = lc.DailyPlayerProfitLimitEnabled
	data[rtp.SlotDailyPlayerProfitLimit] = lc.DailyPlayerProfitLimit
	data[rtp.SlotMonthlyPlayerProfitLimitEnabled] = lc.MonthlyPlayerProfitLimitEnabled
	data[rtp.SlotMonthlyPlayerProfitLimit] = lc.MonthlyPlayerProfitLimit
	data[rtp.SlotMonthlyPlayerRTPLowerLimit] = lc.MonthlyPlayerRTPLowerLimit

	return data
}
