package parse 


import (
	"bufio"
	"compress/gzip"
	"os"
	"sync"
	"io"
	//"fmt"
	"strings"
	"iclip/config"
	"mds/utils"
)

func Read(filename string, config *config.Config) {
	file, err := os.Open(filename)
	utils.Checkerr(err)
	fz, err := gzip.NewReader(file)
	utils.Checkerr(err)
	r := bufio.NewReader(fz)
	var wg2 sync.WaitGroup
	count := 0
	lines := make([]string, 0, 4)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF{
			break
		}else{
			utils.Checkerr(err)
		}
		lines = append(lines, string(line))
		if len(lines) == 4{
			wg2.Add(1)
			var l [4]string
			copy(l[:], lines[0:4])
			go Parse(&wg2, l, config)
			lines = lines[:0]
		}
		count++
	}
	wg2.Wait()
}


func Parse(wg *sync.WaitGroup, lines [4]string, metadata *config.Config){
	defer wg.Done()
	barcode := lines[1] [metadata.BarcodeStart:metadata.BarcodeEnd+1]
	_, ok := metadata.Outputfiles[barcode]
	if ok{
		//0 based AATTTTA if we wnated ts it would be start=2 end=5
		if len(lines[1])>metadata.SequenceEnd{
				lines[1] = lines[1][metadata.SequenceStart:metadata.SequenceEnd+1]
				lines[3] = lines[3][metadata.SequenceStart:metadata.SequenceEnd+1]
				go func(){
					metadata.Outputfiles[barcode].C <- []byte(strings.Join(lines[:], "\n")+"\n")
					}()
				metadata.Outputfiles[barcode].F.Write(<-metadata.Outputfiles[barcode].C)
		}
	}
}

