package db

// var connections []*Connection = nil

func Conns() []*Connection {
	// if connections != nil {
	// 	return connections
	// }
	//
	// connections = make([]*Connection, 0)
	// err := db.Select(&connections, "SELECT * FROM connection WHERE parent = 0")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(len(connections))
	// return connections
	return nil
}

var (
	topConnIds  []int
	Connections map[int]*Connection
	Childrens   map[int][]int
)

func TopConnLength() int {
	return len(topConnIds)
}

func TopConnection(idx int) int {
	return topConnIds[idx]
}

func InitializeConnections() error {
	topConnIds = make([]int, 0)
	Connections = make(map[int]*Connection)
	Childrens = make(map[int][]int)
	_, err := fetch(0)
	return err
}

func fetch(parent int) ([]*Connection, error) {
	var connections []*Connection
	err := db.Select(&connections, "SELECT * FROM connection WHERE parent = ?", parent)
	if err != nil {
		return connections, err
	}

	for _, conn := range connections {
		Connections[conn.Id] = conn
		if parent == 0 {
			topConnIds = append(topConnIds, conn.Id)
		} else {
			Childrens[parent] = append(Childrens[parent], conn.Id)
		}
	}
	return connections, nil
}

func FetchChildrenLength(parent int) int {
	if children, ok := Childrens[parent]; ok {
		return len(children)
	}

	conns, err := fetch(parent)
	if err != nil {
		return 0
	}

	return len(conns)
}
