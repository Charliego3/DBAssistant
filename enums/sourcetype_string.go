// Code generated by "enumer -output=sourcetype_string.go -type DataSourceType -sql -json -values -trimprefix DataSourceType"; DO NOT EDIT.

package enums

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _DataSourceTypeName = "MySQLPostgreRedisMongoFolder"

var _DataSourceTypeIndex = [...]uint8{0, 5, 12, 17, 22, 28}

const _DataSourceTypeLowerName = "mysqlpostgreredismongofolder"

func (i DataSourceType) String() string {
	if i >= DataSourceType(len(_DataSourceTypeIndex)-1) {
		return fmt.Sprintf("DataSourceType(%d)", i)
	}
	return _DataSourceTypeName[_DataSourceTypeIndex[i]:_DataSourceTypeIndex[i+1]]
}

func (DataSourceType) Values() []string {
	return DataSourceTypeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DataSourceTypeNoOp() {
	var x [1]struct{}
	_ = x[DataSourceTypeMySQL-(0)]
	_ = x[DataSourceTypePostgre-(1)]
	_ = x[DataSourceTypeRedis-(2)]
	_ = x[DataSourceTypeMongo-(3)]
	_ = x[DataSourceTypeFolder-(4)]
}

var _DataSourceTypeValues = []DataSourceType{DataSourceTypeMySQL, DataSourceTypePostgre, DataSourceTypeRedis, DataSourceTypeMongo, DataSourceTypeFolder}

var _DataSourceTypeNameToValueMap = map[string]DataSourceType{
	_DataSourceTypeName[0:5]:        DataSourceTypeMySQL,
	_DataSourceTypeLowerName[0:5]:   DataSourceTypeMySQL,
	_DataSourceTypeName[5:12]:       DataSourceTypePostgre,
	_DataSourceTypeLowerName[5:12]:  DataSourceTypePostgre,
	_DataSourceTypeName[12:17]:      DataSourceTypeRedis,
	_DataSourceTypeLowerName[12:17]: DataSourceTypeRedis,
	_DataSourceTypeName[17:22]:      DataSourceTypeMongo,
	_DataSourceTypeLowerName[17:22]: DataSourceTypeMongo,
	_DataSourceTypeName[22:28]:      DataSourceTypeFolder,
	_DataSourceTypeLowerName[22:28]: DataSourceTypeFolder,
}

var _DataSourceTypeNames = []string{
	_DataSourceTypeName[0:5],
	_DataSourceTypeName[5:12],
	_DataSourceTypeName[12:17],
	_DataSourceTypeName[17:22],
	_DataSourceTypeName[22:28],
}

// DataSourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DataSourceTypeString(s string) (DataSourceType, error) {
	if val, ok := _DataSourceTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _DataSourceTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DataSourceType values", s)
}

// DataSourceTypeValues returns all values of the enum
func DataSourceTypeValues() []DataSourceType {
	return _DataSourceTypeValues
}

// DataSourceTypeStrings returns a slice of all String values of the enum
func DataSourceTypeStrings() []string {
	strs := make([]string, len(_DataSourceTypeNames))
	copy(strs, _DataSourceTypeNames)
	return strs
}

// IsADataSourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DataSourceType) IsADataSourceType() bool {
	for _, v := range _DataSourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for DataSourceType
func (i DataSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for DataSourceType
func (i *DataSourceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("DataSourceType should be a string, got %s", data)
	}

	var err error
	*i, err = DataSourceTypeString(s)
	return err
}

func (i DataSourceType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *DataSourceType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of DataSourceType: %[1]T(%[1]v)", value)
	}

	val, err := DataSourceTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
