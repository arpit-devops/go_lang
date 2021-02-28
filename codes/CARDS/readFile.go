package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func (d deck) toReadFile(filename string) ([]string, error) {

	dataByte, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dataString := string(dataByte)
	dataStringSlice := strings.Split(dataString, ",")
	return dataStringSlice, nil

}
