package config

import (
    "os"
    "encoding/json"
    "io/ioutil"
    "mds/utils"
    //"sync"
    "fmt"
)

type Config struct {
    Barcodes []string `json:"barcodes"`
	InputFiles []string `json:"input files"`
    BarcodeStart int `json:"barcode start"`
    BarcodeEnd int `json:"barcode end"`
    SequenceStart int `json:"sequence start"`
    SequenceEnd int `json:"sequence end"`
    Outputfiles map[string]File
}

type stringslice []string

func (s stringslice) Has(item string) bool{
    for _, i := range(s){
        if i == item{
            return true
        }
    }
    return false
}

func (s stringslice ) Index(item string) int {
    for p, v := range s {
        if (v == item) {
            return p
        }
    }
    return -1
}
type File struct {
    F *os.File
    C chan []byte
}

type ConfigParser func(data []byte) Config

func ParseJSON(data []byte) Config {
    res :=Config{Outputfiles:make(map[string]File)}
    err := json.Unmarshal(data, &res)
    utils.Checkerr(err)
    return res
}

func PopulateBarcodes(c *Config, outputdir string){
    for _, barcode := range(c.Barcodes){
        f, err :=os.OpenFile(outputdir+"/"+barcode+"_output.fastq", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
        utils.Checkerr(err)
        c.Outputfiles[utils.ReverseComplement(barcode)] = File{F:f, C:make(chan []byte)}
    }
}

func ReadConfig(filename string, outputdir string) (Config) {
    var parser ConfigParser = ParseJSON
    data, err := ioutil.ReadFile(filename)
    utils.CheckConfig(err, filename)
    parser = ParseJSON
    p := parser(data)
    PopulateBarcodes(&p, outputdir)
    return p
}
