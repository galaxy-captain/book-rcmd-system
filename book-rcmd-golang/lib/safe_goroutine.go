package lib

import "fmt"

func Go(f func()) {
	go func(f func()) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Error - a panic happened: %v\n", r)
			}
		}()
		f() // execute the real-function
	}(f)
}
