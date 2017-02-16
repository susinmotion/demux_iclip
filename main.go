package main

import "iclip/config"
import "iclip/parse"
import "fmt"
import "time"
import "mds/utils"
import "flag"
import "os"
import "path"

var curDir, err = os.Getwd()
var outputPtr = flag.String("O", curDir, "complete path to output directory")
var configPtr = flag.String("C", path.Join(curDir, "config.json"), "complete path to config file")

func main(){
	start := time.Now()
    defer utils.PrintTime(start)
    flag.Parse()
    utils.CheckOutput(*outputPtr)
	c := config.ReadConfig(*configPtr, *outputPtr)
	fmt.Println(c)

	for i := 0; i<len(c.InputFiles); i++{
		fmt.Println("hi")
		parse.Read(c.InputFiles[i], &c)
	}
	for _,file := range(c.Outputfiles){
		file.F.Close()
	}
}
