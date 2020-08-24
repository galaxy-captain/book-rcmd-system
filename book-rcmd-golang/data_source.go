package rcmd

import "sync"

type DataSource interface {
	Read(count int) []interface{}
}

type MySqlDataSource struct {
	lock *sync.Mutex
}

func (m *MySqlDataSource) init() {
	m.lock = new(sync.Mutex)
}

func (m *MySqlDataSource) Read(count int) (result []interface{}) {

	result = make([]interface{}, 0, count)

	return
}
