package main

import (
	"fmt"
	"red_packet/algorithm"
)


func main(){
	count,amount := int64(10),int64(100)
	list := algorithm.DoubleAverageList(count,amount)
	for _,item := range list{
		fmt.Print(item,",")
	}
	fmt.Println()
}
