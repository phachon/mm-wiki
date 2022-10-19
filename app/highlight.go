package app

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search/highlight"
	simpleFragmenter "github.com/blevesearch/bleve/v2/search/highlight/fragmenter/simple"
	simpleHighlighter "github.com/blevesearch/bleve/v2/search/highlight/highlighter/simple"
)

func HighlighterConstructor(config map[string]interface{}, cache *registry.Cache) (highlight.Highlighter, error) {

	fragmenter, err := cache.FragmenterNamed(simpleFragmenter.Name)
	if err != nil {
		return nil, fmt.Errorf("error building fragmenter: %v", err)
	}

	formatter, err := cache.FragmentFormatterNamed(Name)
	if err != nil {
		return nil, fmt.Errorf("error building fragment formatter: %v", err)
	}

	return simpleHighlighter.NewHighlighter(
			fragmenter,
			formatter,
			simpleHighlighter.DefaultSeparator),
		nil
}

func initHighlighter() {
	registry.RegisterHighlighter(Name, HighlighterConstructor)
}
