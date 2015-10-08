package cobertura_test

import (
	"encoding/xml"
	"fmt"

	"github.com/stengaard/cobertura"
)

func Example() {

	xmlFile := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE coverage SYSTEM "http://cobertura.sourceforge.net/xml/coverage-04.dtd">
<coverage line-rate="0.5" branch-rate="0" lines-covered="107" lines-valid="166" branches-covered="0" branches-valid="0" complexity="0" version="0" timestamp="1444233612">
  <sources>
    <source>/source</source>
  </sources>
  <packages>
    <package name="my-package" line-rate="0.5" branch-rate="0" complexity="0">
      <classes>
        <class name="my-class" filename="myClass.rb" line-rate="0.5" branch-rate="0" complexity="0">
          <methods/>
          <lines>
            <line number="1" branch="false" hits="1"/>
          </lines>
        </class>
      </classes>
    </package>
  </packages>
</coverage>`

	cov := cobertura.Coverage{}

	_ = xml.Unmarshal([]byte(xmlFile), &cov)

	class := cov.Packages[0].Classes[0]
	fmt.Printf("%s = %0.2f\n", class.Filename, class.LineRate)

	// Output: myClass.rb = 0.50

}
