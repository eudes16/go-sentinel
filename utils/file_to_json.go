package utils

import (
	"io/ioutil"
	"os"
)

func LoadLanguages(file string) []byte {
	jsonFile, err := os.Open(file)

	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err.Error())
	}

	return data
}
