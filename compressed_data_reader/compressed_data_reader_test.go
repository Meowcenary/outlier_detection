package compressed_data_reader

import (
	"log"
  "os"
	"testing"
)

// Remove the decompressed file so future tests do not erroneously pass
func teardownDecompressTest(filepath string) {
	if err := os.Remove(filepath); err != nil {
		log.Println(err)
	}
}

func TestDecompress(t *testing.T) {
	filepath := "../testdata/simple.txt.gz"
	expectedDecompressedFilepath := "../testdata/simple.txt"
	defer teardownDecompressTest(expectedDecompressedFilepath)
	Decompress(filepath)
	_, err := os.Stat(expectedDecompressedFilepath)

	if err != nil {
		t.Errorf("Expected to find file: %s, but found nothing", expectedDecompressedFilepath)
	}
}
