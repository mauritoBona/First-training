package file

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func ReplaceCharacterInFile(path, fileName, characterToReplace, newCharacter string) {
	input, err := ioutil.ReadFile(path + fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(characterToReplace), []byte(newCharacter), -1)

	if err = ioutil.WriteFile(path+fileName, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
