package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	fmt.Println(time.Now())
	nextTime := cronexpr.MustParse("0 0 * * *").Next(time.Now())

	fmt.Println(nextTime)

	//    os.
}
