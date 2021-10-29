package test

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
)

func GetMatrix() [][]string{
	file, err := os.ReadFile(".././resource/matrix.csv")
	if err != nil {
		log.Fatal("unable to read file")
	}
	reader := bytes.NewReader(file)
	records, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		log.Fatal("unable to read content of file")
	}
	return records
}