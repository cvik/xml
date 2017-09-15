package xml

import (
	"bufio"
	"os"
	"testing"
)

func TestScannerShort(t *testing.T) {
	fd, err := os.Open("ex/weird.xml")
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(fd)
	scanner.Split(split)

	i := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
		i++
	}

	if i != 20 {
		t.Error("Failed to parse all tags (or too many)")
	}

	t.Logf("found %v/20 tags", i)
}

// TODO: break out scanner-test into helper func
func TestScannerLong(t *testing.T) {
	fd, err := os.Open("ex/bing_xsd1.fmt.xsd")
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(fd)
	scanner.Split(split)

	i := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
		i++
	}
	if i != 3008 {
		t.Error("Failed to parse all tags (or too many)")
	}
	t.Logf("found %v/3008 tags", i)
}

func TestScannerWsdl(t *testing.T) {
	fd, err := os.Open("ex/CampaignService.wsdl")
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(fd)
	scanner.Split(split)

	i := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
		i++
	}
	if i != 6614 {
		t.Error("Failed to parse all tags (or too many)")
	}
	t.Logf("found %v/6614 tags", i)

}
