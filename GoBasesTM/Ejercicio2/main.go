package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) [][]string {
    f, _ := os.Open(filePath)
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, _ := csvReader.ReadAll()
    return records
}


func main() {

	sum := 0.0
	val := ""
	records := readCsvFile("../prueba.csv")
    for i, record := range records {
		res := strings.Split(record[0], ";")
		if i > 0 {
			s, _:= strconv.ParseFloat(strings.Trim(res[1]," "),64)
			sum += s
			val += fmt.Sprintf("%s\t %s\t %s \n", res[0],res[1],res[2])
		} else {
			val += fmt.Sprintf("%s\t %s %s \n", res[0],res[1],res[2])
		}
		
	}
	val += fmt.Sprintf("%s\t  %.2f\t" ,"", sum)
	fmt.Println(val)


}