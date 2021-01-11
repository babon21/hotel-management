package usecase

type SortField string

const (
	ErrField       SortField = ""
	PriceField     SortField = "price"
	DateAddedField SortField = "date_added"
)

type SortOrder string

const (
	ErrOrder  SortOrder = ""
	AscOrder  SortOrder = "asc"
	DescOrder SortOrder = "desc"
)
