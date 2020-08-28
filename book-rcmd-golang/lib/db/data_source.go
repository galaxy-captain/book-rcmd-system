package db

type DataSource interface {
	Total() int
	Read(count int) []interface{}
}
