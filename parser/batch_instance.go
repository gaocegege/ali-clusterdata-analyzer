package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

const (
	batchInstanceFileName = "batch_instance.csv"
)

type BatchInstanceParser struct {
	DirName  string
	FileName string
}

func NewBatchInstanceParser(dn string) *BatchInstanceParser {
	return &BatchInstanceParser{
		DirName:  dn,
		FileName: batchInstanceFileName,
	}
}

func (b BatchInstanceParser) GetFileName() string {
	return path.Join(b.DirName, b.FileName)
}

func (b BatchInstanceParser) ParseFromFile() (string, error) {
	file, err := os.Open(b.GetFileName())
	if err != nil {
		return "", err
	}
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
	return "", nil
}
