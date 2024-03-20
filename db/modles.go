package db

import (
	"github.com/charliego3/assistant/enums"
	"github.com/progrium/macdriver/objc"
)

var objClass objc.UserClass[ConnObj]

func init() {
	class := objc.NewClass[ConnObj]()
	objc.RegisterClass(class)
}

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

	Children []*Connection
}

type ConnObj struct {
	objc.Object `objc:"NSObject"`
}

func (obj ConnObj) Id() {

}

func NewConnObj() ConnObj {
	obj := objClass.New()
	return obj
}
