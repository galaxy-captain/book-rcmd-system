package recall

import (
	"context"
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

var driverName = "mysql"
var user = "root"
var pwd = "a761349852."
var host = "127.0.0.1"
var port = "3306"
var dbName = "book"

var DB *sql.DB

func init() {

	// "root:a761349852.@tcp(127.0.0.1:3306)/book"
	db, err := sql.Open(driverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, host, port, dbName))
	if err != nil {
		panic(err.Error())
	} else {
		DB = db
	}

}

func ReadData(ctx context.Context) {

	limit := 100000
	offset := 0
	size := read(0, limit)
	for size > 0 {
		offset = offset + size
		size = read(offset, limit)
	}

	fmt.Println(count)

}

func read(offset, limit int) int {

	query := fmt.Sprintf("select book_id as bookId,user_id as userId from t_bk_user_book_m where book_read_time > %d", 100)
	query = query + fmt.Sprintf(" limit %d offset %d", limit, offset)

	rows, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}

	size := 0
	var bookId, userId string
	for rows.Next() {
		size = size + 1

		err := rows.Scan(&bookId, &userId)
		if err != nil {
			ErrorPrint("ReadData", err)
			continue
		}
		inputData := InputData{BookId: bookId, UserId: userId}

		// send to calculate channel
		InputDataChannel <- inputData
	}

	return size
}
