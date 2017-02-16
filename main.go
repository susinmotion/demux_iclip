package main

import "iclip/config"
import "iclip/parse"
import "fmt"
import "time"
import "mds/utils"

func main(){
	start := time.Now()
    defer utils.PrintTime(start)
	c := config.ReadConfig("config.json")
	fmt.Println(c)
	for i := 0; i<len(c.InputFiles); i++{
		fmt.Println("hi")
		parse.Read(c.InputFiles[i], &c)
	}
	for _,file := range(c.Outputfiles){
		file.F.Close()
	}
}
