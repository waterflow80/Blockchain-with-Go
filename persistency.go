package main

import (
	//"os"
	"encoding/json"
	"errors"
	"fmt"

	//"fmt"
	"os"
	//"io/ioutil"
)

var ErrInexistantBC = errors.New("no existing Blockchain found - Create one first")

// true if the file is existant
func bcFileExists(file string) bool {
	_, err := os.ReadFile(file)
	return err == nil
}

func LoadBlockchain(file string) (*Blockchain,error){
	bcJsonByte, err := os.ReadFile(file)
	if err == nil {
	var bc Blockchain
	json.Unmarshal(bcJsonByte, &bc)
	return &bc, nil
	} else {
		return nil, err
	}
}

func SaveBlockchain(bc *Blockchain, file string) error {
	bcMap := BlockchainToMap(bc)
	bcJson, err := json.Marshal(bcMap)
	//fmt.Printf("%#v", bcMap)
	if err == nil {
		err = os.WriteFile(file, bcJson, 0644)
	} else {
		fmt.Println("persistency.go: Failed to json.Marshal")
	}
	fmt.Println("err=", err)
	return err
}