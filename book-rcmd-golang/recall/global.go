package recall

import "fmt"

type InputData struct {
	BookId string
	UserId string
}

type OutputData struct {
}

var InputDataChannel = make(chan InputData, 128)

var OutputDataChannel = make(chan OutputData, 128)

func ErrorPrint(fn string, err error) {
	fmt.Errorf("[Error] func: %s, msg: %s", fn, err.Error())
}
