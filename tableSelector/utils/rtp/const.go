package config

// Todo檢查Redis的部分
const (
	CountryID    = "CountryID"
	PlatformID   = "PlatformID"
	VendorID     = "VendorID"
	GameCode     = "GameCode"
	RoomTypeID   = "RoomTypeID"
	GameName     = "GameName"
	PlatformName = "PlatformName"
	VendorName   = "VendorName"
	CountryName  = "CountryName"
	RoomTypeName = "RoomTypeName"

	SlotBaseBet                         = "SlotBaseBet"                         // 基礎投注額
	SlotSysRTPLimitEnabled              = "SlotSysRTPLimitEnabled"              // RTP上限功能
	SlotSysRTPLimit                     = "SlotSysRTPLimit"                     // 限制設定系統RTP上限（萬分比）
	SlotMonthlySysLossLimitEnabled      = "SlotMonthlySysLossLimitEnabled"      // 當日系統虧損上限功能
	SlotMonthlySysLossLimit             = "SlotMonthlySysLossLimit"             // 當日系統虧損上限（分）
	SlotDailySysLossLimitEnabled        = "SlotDailySysLossLimitEnabled"        // 當日系統虧損上限功能
	SlotDailySysLossLimit               = "SlotDailySysLossLimit"               // 當日系統虧損上限（分）
	SlotDailyPlayerProfitLimitEnabled   = "SlotDailyPlayerProfitLimitEnabled"   // 當日個人盈利上限功能
	SlotDailyPlayerProfitLimit          = "SlotDailyPlayerProfitLimit"          // 當日個人盈利上限（分）
	SlotMonthlyPlayerProfitLimitEnabled = "SlotMonthlyPlayerProfitLimitEnabled" // 當月個人盈利上限功能
	SlotMonthlyPlayerProfitLimit        = "SlotMonthlyPlayerProfitLimit"        // 當月個人盈利上限（分）
	SlotMonthlyPlayerRTPLowerLimit      = "SlotMonthlyPlayerRTPLowerLimit"      // 當月個人RTP下限（分）

)
