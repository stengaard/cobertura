// Package cobertura contains types for marshalling and unmarshalling the cobertura
// coverage format. (Based on http://cobertura.sourceforge.net/xml/coverage-04.dtd)
package cobertura

import "encoding/xml"

type Condition struct {
	XMLName xml.Name `xml:"condition"`

	Number   int    `xml:"number,attr"`
	Type     string `xml:"type,attr"`
	Coverage string `xml:"coverage,attr"`
}

type Line struct {
	XMLName xml.Name `xml:"line"`

	Number            int    `xml:"number,attr"`
	Hits              int    `xml:"hits,attr"`
	Branch            string `xml:"branch,attr"`
	ConditionCoverage string `xml:"condition-coverage,attr"`

	Conditions []Condition `xml:"conditions>condition"`
}

type Method struct {
	XMLName xml.Name `xml:"method"`

	Signature  string  `xml:"signature,attr"`
	Name       string  `xml:"name,attr"`
	LineRate   float64 `xml:"line-rate,attr"`
	BranchRate float64 `xml:"branch-rate,attr"`

	Lines []Line `xml:"lines>line"`
}

type Class struct {
	XMLName xml.Name `xml:"class"`

	Filename   string  `xml:"filename,attr"`
	Name       string  `xml:"name,attr"`
	LineRate   float64 `xml:"line-rate,attr"`
	BranchRate float64 `xml:"branch-rate,attr"`
	Complexity float64 `xml:"complexity,attr"`

	Lines   []Line   `xml:"lines>line"`
	Methods []Method `xml:"methods>method"`
}

type Package struct {
	XMLName xml.Name `xml:"package"`

	Name       string  `xml:"name,attr"`
	LineRate   float64 `xml:"line-rate,attr"`
	BranchRate float64 `xml:"branch-rate,attr"`
	Complexity float64 `xml:"complexity,attr"`

	Classes []Class `xml:"classes>class"`
}

type Coverage struct {
	XMLName xml.Name `xml:"coverage"`

	Name            string  `xml:"name,attr"`
	LineRate        float64 `xml:"line-rate,attr"`
	LinesCovered    int     `xml:"lines-covered,attr"`
	LinesValid      int     `xml:"lines-valid,attr"`
	BranchRate      float64 `xml:"branch-rate,attr"`
	BranchesCovered int     `xml:"branches-covered,attr"`
	BranchesValid   int     `xml:"branches-valid,attr"`
	Complexity      float64 `xml:"complexity,attr"`
	Version         int     `xml:"version,attr"`
	Timestamp       int     `xml:"timestamp,attr"`

	Packages []Package `xml:"packages>package"`
	Sources  []string  `xml:"source>source"`
}
