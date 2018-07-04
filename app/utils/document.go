package utils

var Document = NewDocument()

const (
	Document_Dir_Page_Name = "README"
	Document_Page_Suffix = ".md"
)

func NewDocument() *document {
	return &document{}
}

type document struct {
	RootDir string
}
