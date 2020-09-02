package product

func CreateProductController(p ProductCreate) ProductCreate {
	product, _ := p.CreateProductMongo(p)
	// p.CreateProductElastic(product)
	// p.CreateProductRedis(product)
	p.CreateProductPub(product)
	return product
}

func GetProductsController() []ProductModel {
	var product ProductModel
	products, _ := product.GetProductsMongo()

	return products
}

func GetProductDetailController(id string) ProductModel {
	var product ProductModel
	product, _ = product.GetProductDetailMongo(id)

	return product
}

func SearchProductController(q ProductSearch) interface{} {
	var product ProductModel
	products := product.SearchProductElastic(q)

	return products
}
