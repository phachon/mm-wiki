package global

import (
	"github.com/blevesearch/bleve/v2"
)

var SearchMap = bleve.NewIndexMapping()
var SearchIndex bleve.Index
