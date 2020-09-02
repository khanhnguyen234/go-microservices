package shop

func CreateShopController(p ShopCreate) ShopCreate {
	shop, _ := p.CreateShopMongo(p)

	p.CreateShopElastic(shop)

	return shop
}

func GetShopsController() []ShopModel {
	var shop ShopModel
	shops, _ := shop.GetShopsMongo()

	return shops
}

func GetShopDetailController(id string) ShopModel {
	var shop ShopModel
	shop, _ = shop.GetShopDetailMongo(id)

	return shop
}

func SearchShopController(q ShopSearch) interface{} {
	var shop ShopModel
	shops := shop.SearchShopElastic(q)

	return shops
}
