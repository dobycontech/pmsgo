package main

import "time"

func main(){
	println("start of process")
	<-time.After(5*time.Minute)
	println("end of process")
}