package main

import "iclip/config"
import "iclip/parse"
import "fmt"

func main(){
	c := config.ReadConfig("/Users/SusanSteinman/Documents/go/src/iclip/config.json")
	fmt.Println(c)
	for i := 0; i<len(c.InputFiles); i++{
		fmt.Println("hi")
		parse.Read(c.InputFiles[i], &c)
	}
	for _,file := range(c.Outputfiles){
		file.F.Close()
	}
}