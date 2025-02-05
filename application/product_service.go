package application

type ProdructService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProdructService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProdructService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err : product.IsValid()
	if err != nill {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		result &Product{}, err
	}
	return result, nil
}
