package product

func CreateProductController(p ProductCreate) ProductCreate {
	var product ProductCreate
	if p.Id == "" {
		product, _ = p.CreateProductMongo(p)
	} else {
		product, _ = p.UpdateProductMongo(p)
	}
	// p.CreateProductElastic(product)
	// p.CreateProductRedis(product)
	p.CreateProductPub(product)
	return product
}

func GetProductsController() []ProductCreate {
	var product ProductCreate
	products, _ := product.GetProductsMongo()

	return products
}

func GetProductsFlashSaleController(query ProductFlashSale) []ProductCreate {
	var product ProductCreate
	var products []ProductCreate

	if query.Src == "cache" {
		products, _ = product.GetProductsFlashSaleRedis(query)
	}

	if products == nil {
		products, _ = product.GetProductsFlashSaleMongo(query)
	}

	return products
}

func GetProductDetailController(id string) ProductCreate {
	var product ProductCreate
	product, _ = product.GetProductDetailMongo(id)

	return product
}

func SearchProductController(q ProductSearch) interface{} {
	var product ProductModel
	products := product.SearchProductElastic(q)

	return products
}
