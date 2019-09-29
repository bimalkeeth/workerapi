package errorhandle

import "fmt"
//----------------------------------------------
//function to handle error with panic
//----------------------------------------------
func Error(message string, err error) {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
//-----------------------------------------------
// function to manage panic
//-----------------------------------------------
func Panicking() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
}

