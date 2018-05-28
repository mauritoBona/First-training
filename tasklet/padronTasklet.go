package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"padrones/First-training/file"
	"padrones/First-training/padronModel"
	"padrones/First-training/threads"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	padronModel.CompleteMaps()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var path = "D:\\GO\\src\\padrones\\First-training\\file\\"
	var fileName = "PadronRGSRet.txt"
	file.ReplaceCharacterInFile(path, fileName, ",", ".")
	f, err := os.Open(path + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))

	processRecord(r)
}

func processRecord(r *csv.Reader) {
	contador := 0
	start := time.Now()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		for value := range record {
			wg.Add(1)
			go threads.BuildInsertMassivePadronFromCsv(strings.Split(record[value], ";"), &wg)
			if contador%10000 == 0 {
				wg.Wait()
			}
		}
		contador++
	}
	wg.Wait()
	threads.RemoveLastCharacter()
	threads.ExecuteInsertMassive()
	elapsed := time.Since(start)
	fmt.Print("TARDO: ")
	fmt.Print(elapsed)
}
