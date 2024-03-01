package methods

import (
	"fmt"
	"time"
)

type MyErrors struct {
	When time.Time
	What string
}

func (e *MyErrors) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyErrors{
		time.Now(),
		"tidak berkerja!",
	}
}

func Errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
