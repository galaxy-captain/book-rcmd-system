package db

import (
	"database/sql"
	"sync"
)
import _ "github.com/go-sql-driver/mysql"

type MySqlDataSource struct {
	lock   *sync.Mutex
	db     *sql.DB
	size   int64
	offset int64
}

func (m *MySqlDataSource) Init() {

	db, err := sql.Open("mysql", "root:a761349852.@tcp(127.0.0.1:3306)/book")
	if err != nil {
		panic(err.Error())
	}

	m.db = db
	m.offset = 0
	m.size = m.total()
}

func (m *MySqlDataSource) total() int64 {
	var total int64
	r := m.db.QueryRow("select count(1) as total from t_bk_user_book_m")
	err := r.Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	return total
}

func (m *MySqlDataSource) occupy(count int64) (start, end int64) {
	defer m.lock.Unlock()
	m.lock.Lock()

	if m.offset >= m.size {
		return -1, -1
	}

	start = m.offset
	end = m.offset + count
	if end > m.size {
		end = m.size
	}
	m.offset = end

	return
}

func (m *MySqlDataSource) Total() int64 {
	return m.size
}

func (m *MySqlDataSource) Read(count int64) (result []interface{}) {

	start, end := m.occupy(count)
	if start == -1 || end == -1 {
		return nil
	}

	result = make([]interface{}, 0, count)

	return
}
