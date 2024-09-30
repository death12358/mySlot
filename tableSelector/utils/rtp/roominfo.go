package config

import "fmt"

// GameInfo 遊戲資訊
type GameInfo struct {
	CountryID    int32  `yaml:"country_id"`    // 國家ID
	PlatformID   int32  `yaml:"platform_id"`   // 包網ID
	VendorID     int32  `yaml:"vendor_id"`     // 代理ID
	GameCode     string `yaml:"game_code"`     // 遊戲Code
	RoomType     int32  `yaml:"room_type"`     // 房間等級
	Bet          int32  `yaml:"bet"`           //投注額
	GameName     string `yaml:"game_name"`     // 遊戲名稱
	PlatformName string `yaml:"platform_name"` // 包網名稱
	VendorName   string `yaml:"vendor_name"`   // 代理名稱
	CountryName  string `yaml:"country_name"`  // 幣種名稱
	RoomTypeName string `yaml:"roomtype_name"` // 房間名稱
}

// r.CountryID_r.PlatformID_r.VendorID_r.GameCode_r.RoomType_r.Bet
func (r *GameInfo) GetRTPConfigKey() (string, error) {
	if err := r.vld(); err != nil {
		fmt.Printf("%+v", err.Error())
		return "", err
	}
	return fmt.Sprintf("%s:%d:%d:%d:%d:%d", r.GameCode, r.CountryID, r.PlatformID, r.VendorID, r.RoomType, r.Bet), nil
}

// r.CountryID_r.PlatformID_r.VendorID_r.GameCode_r.RoomType_r.Bet
func (r *GameInfo) GetKey_bet() (string, error) {
	if err := r.vld(); err != nil {
		fmt.Printf("%+v", err.Error())
		return "", err
	}
	return fmt.Sprintf("%s:%d:%d_%d_%d:%d", r.GameCode, r.RoomType, r.CountryID, r.PlatformID, r.VendorID, r.Bet), nil
}
func (r *GameInfo) GetKey() (string, error) {
	if err := r.vld(); err != nil {
		fmt.Printf("%+v", err.Error())
		return "", err
	}
	return fmt.Sprintf("%s:%d:%d_%d_%d", r.GameCode, r.RoomType, r.CountryID, r.PlatformID, r.VendorID), nil
}

func (r *GameInfo) GetGameKey() (string, error) {
	if err := r.gameVld(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d_%d_%d", r.GameCode, r.CountryID, r.PlatformID, r.VendorID), nil
}

func (r *GameInfo) vld() error {
	if r.CountryID < 0 || r.PlatformID < 0 || r.VendorID < 0 || r.GameCode == "" || r.RoomType < 0 {
		return fmt.Errorf("遊戲資訊錯誤 CountryID:[%d] PlatformID:[%d] VendorID:[%d] GameCode:[%s] RoomType:[%d]",
			r.CountryID, r.PlatformID, r.VendorID, r.GameCode, r.RoomType)
	}
	return nil
}

func (r *GameInfo) gameVld() error {
	if r.CountryID < 0 || r.PlatformID < 0 || r.VendorID < 0 || r.GameCode < "" {
		return fmt.Errorf("遊戲資訊錯誤 CountryID:[%d] PlatformID:[%d] VendorID:[%d] GameCode:[%s]",
			r.CountryID, r.PlatformID, r.VendorID, r.GameCode)
	}
	return nil
}

func GameKeyToKey(gameKey string, roomType int32) string {
	return fmt.Sprintf("%v_%v", gameKey, roomType)
}
