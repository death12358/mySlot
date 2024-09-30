package tables

// excel上的名稱
const (
	SystemWinMonthlyRTP          string = "SysWinMRTP"      // 系統贏流程- 當月系統 RTP 上限
	SystemWinMonthlySysLoss      string = "SysWinMLoss"     // 系統贏流程- 當月系統虧損上限
	SystemWinDailySysLoss        string = "SysWinDLoss"     // 系統贏流程- 當日系統虧損上限
	SystemWinDailyPlayerProfit   string = "SysWinDPrProfit" // 系統贏流程- 當日個人盈利上限
	SystemWinMonthlyPlayerProfit string = "SysWinMPrProfit" // 系統贏流程- 當月個人盈利上限
	RandomFlowProfitLimit        string = "RandomFlow"      // 隨機流程倍數上限
	ExpectedPay_RandomFlow       string = "期望倍率_隨機流程"
	ExpectedPay_SysWin           string = "期望倍率_系統贏"
)
