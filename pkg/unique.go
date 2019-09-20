package unique

import (
	"encoding/csv"
	"io"
	"os"
)

// CSVOptions stores the options for the
type CSVOptions struct {
	HasHeaders bool
	Column     int
}

// CSV takes a file as a Reader and returns an int of the unique count.
func CSV(f *os.File, opts CSVOptions) (int, error) {
	data := make(map[string]int)

	rdr := csv.NewReader(f)

	for {
		line, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		if line != nil {
			col := line[0]
			if d, ok := data[col]; ok {
				data[col] = d + 1
			} else {
				data[col] = 1
			}
		}
	}

	return len(data), nil
}
