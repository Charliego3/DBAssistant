package enums

//go:generate enumer -output=sourcetype_string.go -type DataSourceType -json -values -trimprefix Fee
type DataSourceType uint

const (
	DataSourceTypeMySQL DataSourceType = iota
	DataSourceTypePostgre
	DataSourceTypeRedis
	DataSourceTypeMongo
)
