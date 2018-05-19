package main

import "fmt"
import "encoding/csv"
import "os"
import "log"

func main() {

	fl, err := os.Open("data/unreachable.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer fl.Close()

	reader := csv.NewReader(fl)

	reader.FieldsPerRecord = -1  // can be anything

	d, err := reader.ReadAll()

	fmt.Println(d)
	
}