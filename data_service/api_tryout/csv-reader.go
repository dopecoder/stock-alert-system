package api_tryout

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

func CSVToMap(reader io.Reader) []map[string][]string {
	r := csv.NewReader(reader)
	rows := []map[string][]string{}
	header := make([]string, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string][]string{}
			// for i := range header {
			if dict[header[1]] != nil {
				fmt.Printf("Duplicate token found! - %s\n", record)
			}
			dict[header[1]] = record
			// }
			rows = append(rows, dict)
		}
	}
	return rows
}
