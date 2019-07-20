package main

import "log"

const layoutStr = "2006-01-02 15:04:05"

func checkErr(err error)  {
	if err != nil {
		log.Fatal("eeeee", err)
	}
}
