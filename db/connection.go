package db

import "fmt"

var connections []*Connection = nil

func Conns() []*Connection {
	if connections != nil {
		return connections
	}

	connections = make([]*Connection, 0)
	err := db.Select(&connections, "SELECT * FROM connection WHERE parent = 0")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(len(connections))
	return connections
}
