package utility

import (
	"encoding/csv"
	"io"
	"os"
)

//CsvReadAll 读取csv
func CsvReadAll(filename string) ([]map[string]string, error) {
	records := []map[string]string{}
	file, err := os.Open(filename)
	if err != nil {
		return records, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	cols, err := csvReader.Read()
	if err == io.EOF {
		return records, err
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		row := map[string]string{}
		for index, value := range record {
			row[cols[index]] = value
		}

		records = append(records, row)
	}

	return records, nil
}

//CsvWriteLine 写入csv
func CsvWriteLine(filename string, records ...[]string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for _, record := range records {
		err = csvWriter.Write(record)
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return nil
}