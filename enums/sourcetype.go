package enums

//go:generate enumer -output=sourcetype_string.go -type DataSourceType -sql -json -values -trimprefix DataSourceType
type DataSourceType uint

const (
	DataSourceTypeMySQL DataSourceType = iota
	DataSourceTypePostgre
	DataSourceTypeRedis
	DataSourceTypeMongo

	DataSourceTypeFolder
)
