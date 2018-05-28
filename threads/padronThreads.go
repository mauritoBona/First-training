package threads

import (
	"fmt"
	"padrones/First-training/mySql"
	"padrones/First-training/padronModel"
	"strings"
	"sync"
)

var mutex sync.Mutex
var contador = 0
var max = 10000
var insertMassive strings.Builder
var contadorDeInsert = 0

func init() {
	insertMassive.WriteString(padronModel.BuildInsertInitializeQuery())
}

func BuildInsertMassivePadronFromCsv(padronCsv []string, wg *sync.WaitGroup) {
	defer wg.Done()
	padron := padronModel.Padron{}
	padron.BuildPadronFromCSV(padronCsv)
	insert := padron.BuildInsertValueQuery()
	mutex.Lock()
	insertMassive.WriteString(insert)
	contador++
	if contador == max {
		ExecuteInsertMassive()
	} else {
		insertMassive.WriteString(",")
	}
	mutex.Unlock()
}

func resetValues() {
	contadorDeInsert += contador
	fmt.Printf("Se Realizaron: %d inserts.\n", contadorDeInsert)
	insertMassive.Reset()
	insertMassive.WriteString(padronModel.BuildInsertInitializeQuery())
	contador = 0
}
func ExecuteInsertMassive() {
	insertMassive.WriteString(";")
	mySql.ExecuteQuerys(insertMassive.String())
	resetValues()
}

func RemoveLastCharacter() {
	var stringToChange string = insertMassive.String()
	insertMassive.Reset()
	insertMassive.WriteString(stringToChange[:len(stringToChange)-1])
}
