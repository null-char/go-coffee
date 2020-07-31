package data

// Product defines the structure of a product in our API.
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	Discount    float32
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}
