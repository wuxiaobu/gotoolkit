package utility

import (
	"testing"
)

func TestCsv(t *testing.T) {
	filename := "../test/sample.csv"

	record := []string{"id", "name"}
	err := CsvWriteLine(filename, record)
	if err != nil {
		t.Error(err)
	}

	_, err = CsvReadAll(filename)
	if err != nil {
		t.Error(err)
	}

	//os.Remove(filename)
}
