package config

import (
	"fmt"

	"gitlab.baifu-tech.net/v3/slot-table/utils/rtp/vld"
)

type LimitConfig struct {
	BaseBet                              int64 `yaml:"base_bet"`                                  // 基礎投注額
	SysRTPLimitEnabled                   bool  `yaml:"sys_rtp_limit_enabled"`                     // 系統RTP上限功能
	SysRTPLimit                          int32 `yaml:"sys_rtp_limit"`                             // 系統RTP上限（萬分比)
	MonthlySysLossLimitEnabled           bool  `yaml:"monthly_sys_loss_limit_enabled"`            // 當月系統虧損上限功能
	MonthlySysLossLimit                  int64 `yaml:"monthly_sys_loss_limit"`                    // 當月系統虧損上限（分）
	DailySysLossLimitEnabled             bool  `yaml:"daily_sys_loss_limit_enabled"`              // 當日系統虧損上限功能
	DailySysLossLimit                    int64 `yaml:"daily_sys_loss_limit"`                      // 當日系統虧損上限（分）
	DailyPlayerProfitLimitEnabled        bool  `yaml:"daily_player_profit_limit_enabled"`         // 當日個人盈利上限功能
	DailyPlayerProfitLimit               int64 `yaml:"daily_player_profit_limit"`                 // 當日個人盈利上限（分）
	MonthlyPlayerProfitLimitEnabled      bool  `yaml:"monthly_player_profit_limit_enabled"`       // 當月個人盈利上限功能
	MonthlyPlayerProfitLimit             int64 `yaml:"monthly_player_profit_limit"`               // 當月個人盈利上限（分）
	MonthlyPlayerProfitLowerLimitEnabled bool  `yaml:"monthly_player_profit_lower_limit_enabled"` // 當月個人盈利上限功能
	MonthlyPlayerRTPLowerLimit           int64 `yaml:"monthly_player_rtp_lower_limit"`            // 當月個人RTP下限（分）
}

// 限制配置數值驗證
func (l *LimitConfig) ValueVLD() error {
	if !vld.PositiveIntVLD(l.BaseBet) {
		return fmt.Errorf("基礎投注額須為正整數. BaseBet[%v]", l.BaseBet)
	}
	if !vld.PositiveIntVLD(l.SysRTPLimit) && l.SysRTPLimitEnabled {
		return fmt.Errorf("系統RTP上限須為正整數. SysRTPLimit[%v]", l.SysRTPLimit)
	}
	if !vld.PositiveIntVLD(l.MonthlySysLossLimit) && l.MonthlySysLossLimitEnabled {
		return fmt.Errorf("當月系統虧損上限須為正整數. MonthlySysLossLimit[%v]", l.MonthlySysLossLimit)
	}
	if !vld.PositiveIntVLD(l.DailySysLossLimit) && l.DailySysLossLimitEnabled {
		return fmt.Errorf("當日系統虧損上限須為正整數. DailySysLossLimit[%v]", l.DailySysLossLimit)
	}

	if !vld.PositiveIntVLD(l.DailyPlayerProfitLimit) && l.DailyPlayerProfitLimitEnabled {
		return fmt.Errorf("當日個人盈利上限須為正整數. DailyPlayerProfitLimit[%v]", l.DailyPlayerProfitLimit)
	}
	if !vld.PositiveIntVLD(l.MonthlyPlayerProfitLimit) && l.MonthlyPlayerProfitLimitEnabled {
		return fmt.Errorf("當月個人盈利上限須為正整數. MonthlyPlayerProfitLimit[%v]", l.MonthlyPlayerProfitLimit)
	}
	return nil
}
