package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	presentTime := time.Now()// .UnixNano() -> 1753020826131488000
	fmt.Println(presentTime) //2025-07-20 19:08:22.412752 +0530 IST m=+0.000218251%

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05")) // we are going to use this time as a stanard to format written in the documentation

	createDate := time.Date(2004, time.June, 4, 12, 45,0,0,time.Local)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}

// GOOS="windows" go build  -> build for windows 
// GOOS="darwin" go build  -> build for mac 

// In golang memory management is not your job
// memory allocation and de allocation happen by itself