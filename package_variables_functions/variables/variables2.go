package main

import "fmt"

//initialize varibales then we can ommit type

var i, j int = 10 , 20 

func inferredTypes() {
	var c , python , java = true , false , "no!" //type is inferred from the value
	fmt.Println(i, j, c, python, java)
}