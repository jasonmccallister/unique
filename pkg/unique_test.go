package unique

import (
	"os"
	"testing"
)

func TestCSV(t *testing.T) {
	testCases := []struct {
		desc      string
		file      string
		hasHeader bool
		column    int
		expected  int
	}{
		{
			desc:      "default column results to the first",
			file:      "testdata/emails.csv",
			hasHeader: true,
			expected:  2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// open the file
			f, err := os.Open(tC.file)
			if err != nil {
				t.Fatal(err)
			}

			opts := CSVOptions{HasHeaders: false, Column: 0}

			// create the reader
			actual, err := CSV(f, opts)
			if err != nil {
				t.Error(err)
			}

			if actual != tC.expected {
				t.Errorf("expected count to be %v, got %v instead", tC.expected, actual)
			}
		})
	}
}
