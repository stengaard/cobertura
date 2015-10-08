package cobertura

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

// simple test to check that we don't err out while unmarshalling xml
func TestUnmarshal(t *testing.T) {
	buf, err := ioutil.ReadFile("coverage.xml")
	if err != nil {
		t.Fatal(err)
	}

	cov := Coverage{}

	err = xml.Unmarshal(buf, &cov)
	if err != nil {
		t.Fatal(err)
	}

}
