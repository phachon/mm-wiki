package app

import (
	"html"

	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search/highlight"
)

const Name = "mm-wiki"

//更换颜色代码，可参考：https://html-color-codes.info/chinese/
const defaultHTMLHighlightBefore = "<span style=\"color:#F03F3F\">"
const defaultHTMLHighlightAfter = "</span>"

type FragmentFormatter struct {
	before string
	after  string
}

func NewFragmentFormatter(before, after string) *FragmentFormatter {
	return &FragmentFormatter{
		before: before,
		after:  after,
	}
}

func (a *FragmentFormatter) Format(f *highlight.Fragment, orderedTermLocations highlight.TermLocations) string {
	rv := ""
	curr := f.Start
	for _, termLocation := range orderedTermLocations {
		if termLocation == nil {
			continue
		}
		// make sure the array positions match
		if !termLocation.ArrayPositions.Equals(f.ArrayPositions) {
			continue
		}
		if termLocation.Start < curr {
			continue
		}
		if termLocation.End > f.End {
			break
		}
		// add the stuff before this location
		rv += html.EscapeString(string(f.Orig[curr:termLocation.Start]))
		// start the <mark> tag
		rv += a.before
		// add the term itself
		rv += html.EscapeString(string(f.Orig[termLocation.Start:termLocation.End]))
		// end the <mark> tag
		rv += a.after
		// update current
		curr = termLocation.End
	}
	// add any remaining text after the last token
	rv += html.EscapeString(string(f.Orig[curr:f.End]))

	return rv
}

func Constructor(config map[string]interface{}, cache *registry.Cache) (highlight.FragmentFormatter, error) {
	before := defaultHTMLHighlightBefore
	beforeVal, ok := config["before"].(string)
	if ok {
		before = beforeVal
	}
	after := defaultHTMLHighlightAfter
	afterVal, ok := config["after"].(string)
	if ok {
		after = afterVal
	}
	return NewFragmentFormatter(before, after), nil
}

func initFragmentFormatter() {
	registry.RegisterFragmentFormatter(Name, Constructor)
}
