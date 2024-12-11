package sorting

type SortDirection string

const (
	ASC  SortDirection = "asc"
	DESC SortDirection = "desc"
)

type SortField struct {
	Field     string        `json:"field"`
	Direction SortDirection `json:"direction"`
}

type Sortable struct {
	SortFields []SortField `json:"sort_fields,omitempty"`
	Enabled    bool        `json:"-" swaggerignore:"true"`
}
