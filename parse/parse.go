package parse 


import (
	"bufio"
	"compress/gzip"
	"os"
	"sync"
	"io"
	"fmt"
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

	if !metadata.Barcodes.Has(barcode){
		metadata.Lock.Lock()
		f, err :=os.OpenFile(utils.ReverseComplement(barcode)+"_output.fastq", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
		utils.Checkerr(err)
		metadata.Outputfiles = append (metadata.Outputfiles, config.File{F:f})
		metadata.Barcodes = append(metadata.Barcodes, barcode)
		metadata.Lock.Unlock()
	}
	index := metadata.Barcodes.Index(barcode)
	//0 based AATTTTA if we wnated ts it would be start=2 end=5
	
	lines[1] = lines[1][metadata.SequenceStart:metadata.SequenceEnd+1]
	lines[3] = lines[3][metadata.SequenceStart:metadata.SequenceEnd+1]
	metadata.Outputfiles[index].L.Lock()
	for i:=0; i<len(lines); i++{
		fmt.Println(string([]byte(lines[i]+"\n")))
		metadata.Outputfiles[index].F.Write([]byte(lines[i]+"\n"))
	}
	metadata.Outputfiles[index].L.Unlock()


}

