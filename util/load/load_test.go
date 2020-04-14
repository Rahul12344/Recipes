package load

import "testing"

func LoadTest(t *testing.T) {
	dataset := DownloadDataset("test.csv")
	if len(dataset) == 0 {
		t.Errorf("Failed to read data")
	}
}
