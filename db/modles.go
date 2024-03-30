package db

import (
	"github.com/charliego3/assistant/enums"
)

type Connection struct {
	Id       int
	Type     enums.DataSourceType
	Name     string `db:"name"`
	Host     string
	Username string `db:"user"`
	Password string `db:"pass"`
	Database string
	Timezone string
	Parent   int `db:"parent"`
}
