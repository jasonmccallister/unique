package unique

import (
	"os"
	"testing"
)

func TestCSV(t *testing.T) {
	testCases := []struct {
		desc           string
		file           string
		hasHeader      bool
		column         int
		expectedUnique int
		expectedTotal  int
	}{
		{
			desc:           "default column results to the first",
			file:           "../testdata/emails.csv",
			hasHeader:      true,
			expectedUnique: 2,
			expectedTotal:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// open the file
			f, err := os.Open(tC.file)
			if err != nil {
				t.Fatal(err)
			}

			opts := Options{HasHeaders: false, Column: 0}

			// create the reader
			actualUnique, actualTotal, err := CSV(f, opts)
			if err != nil {
				t.Error(err)
			}

			if actualUnique != tC.expectedUnique {
				t.Errorf("expected count to be %v, got %v instead", tC.expectedUnique, actualUnique)
			}

			if actualTotal != tC.expectedTotal {
				t.Errorf("expected count to be %v, got %v instead", tC.expectedUnique, actualTotal)
			}
		})
	}
}
