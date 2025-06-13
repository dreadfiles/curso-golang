package model

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       int    `json:"value"`
}

func (product *Product) NewProduct(id int, name string, description string, value int) {
	product.ID = id
	product.Name = name
	product.Description = description
	product.Value = value
}
