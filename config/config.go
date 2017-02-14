package config

import (
    "os"
    "encoding/json"
    "io/ioutil"
    "mds/utils"
    "sync"
    //"fmt"
)

type Config struct {
    Lock sync.Mutex
	InputFiles []string `json:"input files"`
    BarcodeStart int `json:"barcode start"`
    BarcodeEnd int `json:"barcode end"`
    SequenceStart int `json:"sequence start"`
    SequenceEnd int `json:"sequence end"`
    Outputfiles []File
    Barcodes stringslice //this is absolutely the worst! I actually just want to be able to access the lock within a map that links barcode to file+lock, but I can't because of restrictions talking to objects from a pointer to a map
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
    L sync.Mutex
}

type ConfigParser func(data []byte) Config

func ParseJSON(data []byte) Config {
    res :=Config{}
    err := json.Unmarshal(data, &res)
    utils.Checkerr(err)
    return res
}

func ReadConfig(filename string) (Config) {
    var parser ConfigParser = ParseJSON
    data, err := ioutil.ReadFile(filename)
    utils.CheckConfig(err, filename)
    parser = ParseJSON
    return parser(data)
}