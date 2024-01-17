package main

import (
	//"os"
	"errors"
	//"encoding/json" to save/load the blockchain using json encoding  
        //"io/ioutil"
)

var ErrInexistantBC = errors.New("no existing Blockchain found - Create one first")

// true if the file is existant
func bcFileExists(file string) bool {
	return true
}

func LoadBlockchain(file string) (*Blockchain,error){
	return nil,nil
}

func SaveBlockchain(bc *Blockchain, file string) error {
	return nil
}
