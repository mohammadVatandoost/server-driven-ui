package pages

const (
	TitleType       = "Title"
	DescriptionType = "Description"
	ImageType       = "Image"
	TextType        = "Text"
	TableType       = "Table"
)

type Widget struct {
	Type string      `json:"Type"`
	Data interface{} `json:"Data,omitempty"`
}

type TitleWidget struct {
	Title string `json:"Title"`
}

type DescriptionWidget struct {
	Description string `json:"Description"`
}

type ImageWidget struct {
	Source string `json:"Source"`
}

type TextWidget struct {
	Text string `json:"Text"`
}

type TableWidget struct {
	Titles []string   `json:"Titles"`
	Rows   [][]string `json:"Rows"`
}
