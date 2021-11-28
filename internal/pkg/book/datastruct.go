package book

type Enriched struct {
	Book
	InCartCount int
	Authors     []Author
}

type Book struct {
	ID            int64
	Name          string
	PublishedYear int
	Description   string
	Price         int
	Image         []byte
}

type Author struct {
	ID       int64
	Name     string
	Surname  string
	FullName string
}

type ListBooksQuery struct {
	Limit  int
	Offset int
	Filter *ListBooksFilter
}

type ListBooksRequest struct {
	Page   int
	Size   int
	Filter *ListBooksFilter
}

type ListBooksFilter struct {
	Name    string
	Session string
}
