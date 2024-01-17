package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// IntToHex converts an int64 to a byte array of length 8
//
func IntToHex(num int64) []byte {
	bytesArr := make([]byte, 8) 
	hexadecimalValue := fmt.Sprintf("%x", num)
	bytesTempArr, _ := hex.DecodeString(hexadecimalValue)
	copy(bytesArr, bytesTempArr)
	return bytesArr
}

/**
Convert a byte array into a string of bytes
eg: 0x0 -> 00; 0x0a -> 0a, etc*/
func ByteArrToString(arr []byte) string {
	return fmt.Sprintf("%x", arr)
}

/**
Convert a byte array into a string of its corresponding bits sequence
Eg: Input: 1234567895
		Output: "00000000 00000000 00000000 00000000 01001001 10010110 00000010 11010111"	*/
func byteArrToStringBits(byteArr []byte) string {
	bitsStr := fmt.Sprintf("%08b", byteArr) // [00000000 00101011 ...]
	bitsStr = strings.ReplaceAll(bitsStr," ", "") // [0000000000101011 ...] (removed empty spaces)
	bitsStr = strings.ReplaceAll(bitsStr,"[", "") // 0000000000101011 ...] (removed left barcket)
	bitsStr = strings.ReplaceAll(bitsStr,"]", "") // 0000000000101011 ... (removed right barcket)
	return bitsStr
}

// true if the hash starts with zeroBits zeros, note that the hash is
// a slice of *bytes* but we want zeroBits *bits* (a byte has 8 bits)
func StartsWithXZeros(hash []byte, zeroBits int) bool { 
	hashBitsStr := byteArrToStringBits(hash)
	count := 0
	i := 0 // iterate over the hash bits
	for count < zeroBits && i < len(hashBitsStr) {
		if hashBitsStr[i] != '0' {
			return false
		} 
		i ++
		count ++
	}

	return count == zeroBits;
}

func EqualSlices(a, b []byte) bool {
	if (len(a) != len(b)) {
		return false
	}
	for i := 0; i< len(a); i++ {
		if (a[i] != b[i]) {
			return false
		}
	}
	return true
}

func EqualMaps(a, b map[string]int) bool {
  if (len(a) != len(b)) {
		return false
	}
	for key, val := range(a) {
		_, ok := b[key]
		if (!ok) {
			return false
		} 
		if b[key] != val {
			return false
		}
	}
	return true
}

func EqualTransactions(a,b Transaction) bool{
	return EqualSlices(a.Hash,b.Hash)
}

func EqualBlocks(a,b Block) bool{
	return EqualSlices(a.Hash,b.Hash)
}

/**
Return the number of all elements in the given 2 dimensional array
*/
func numElements2DArr(arr2D [][]byte) int {
	numElts := 0
	for _, arr := range arr2D {
		numElts += len(arr)
	}
	return numElts
}

// Serializes a slice of byte slices by converting it to a byte slice so
// needed to easily hash data
func Serialize(input [][]byte )[]byte {
	numElts := numElements2DArr(input)
	inputArr := make([]byte, numElts) // The final []byte array
	it := 0 // iterator for inputArr
	for _, arr := range input {
		for i:=0; i<len(arr); i++ {
			inputArr[it] = arr[i]
			it ++;
		}
	}
	return inputArr;
}

