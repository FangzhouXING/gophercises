package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	filepath := "./problems.csv"
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Cannot open ", filepath)
		return
	}
	defer f.Close()

	correct, wrong := 0, 0

	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("What is %s ? ", record[0])
		var ans int
		fmt.Scanf("%d", &ans)
		if strconv.Itoa(ans) == record[1] {
			correct++
			fmt.Println("Correct")
		} else {
			wrong++
			fmt.Println("Wrong")
		}
	}

	fmt.Printf("Correct %d, Wrong %d\n", correct, wrong)

}
