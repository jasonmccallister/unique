package unique

import (
	"encoding/csv"
	"io"
	"os"
)

// Options stores the options for the
type Options struct {
	HasHeaders bool
	Column     int
}

// CSV takes a file as a Reader and returns an int of the unique count.
func CSV(f *os.File, opts Options) (int, error) {
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
			col := line[opts.Column]
			if d, ok := data[col]; ok {
				data[col] = d + 1
			} else {
				data[col] = 1
			}
		}
	}

	return len(data), nil
}
