package template

import (
	"sync"
)

func (m TmpManager) SetTableMapToRedis(gameCode string, tMap map[string]TableMap, wg *sync.WaitGroup) error {
	for k, v := range tMap {
		err := m.SetTableMap(gameCode, v, k, wg)
		if err != nil {
			return err
		}
	}

	return nil
}
