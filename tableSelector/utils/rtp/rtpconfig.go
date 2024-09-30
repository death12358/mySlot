package config

type GameConfig struct {
	GameInfo GameInfo `yaml:"game_info"` // 遊戲資訊
}
type RTPConfigs struct {
	LimitConfigs map[string]LimitConfig `yaml:"limit_config"`
}

type TemplateKeyMap map[string]map[string]interface{}

// type GameRTPConfig struct {
// 	GameConfigs    []GameConfig                      `yaml:"game_configs"`     //遊戲資訊
// 	RTPConfigs     map[string]RTPConfig              `yaml:"rtp_configs"`      //map[配置ID]RTP配置
// 	TemplateKeyMap map[string]map[string]interface{} `yaml:"template_key_map"` //map[countryName:platformName:venderName]map[roomKey_bet]配置Name
// }
