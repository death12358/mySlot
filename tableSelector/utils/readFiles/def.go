package readFiles

type DataMap map[string][]string

type CSVOrder int32

const (
	CountryName CSVOrder = iota
	PlatformName
	VendorName
	GameCode
	Bet
	TmpName
	HighTableName
	NormalTableName
	LowTableName
)
