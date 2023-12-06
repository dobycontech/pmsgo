// Package main: The main function waits for 5 minutes before printing "end of process".
package main

import "time"

func main() {
	println("start of process")
	<-time.After(5 * time.Minute)
	println("end of process")
}
