package inventory

// TODO: contract repo implementation
type Repository interface {
	GetAll(page int, limit int) ([]Inventory, error)
	GetByCode(code string) (Inventory, error)
}
