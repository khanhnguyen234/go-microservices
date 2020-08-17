package procedure

import (
	"fmt"
	"khanhnguyen234/sql/_postgres"
)

func ProductSearch() {
	db := _postgres.GetPostgres()
	db.Exec(`
		CREATE OR REPLACE PROCEDURE pro_get_products_by_name(
			INOUT product_name text DEFAULT null, 
			INOUT product_price int DEFAULT null
		)
		  LANGUAGE plpgsql AS
		$$
		BEGIN
		   SELECT name, price FROM product_models WHERE id = 1
		   INTO product_name, product_price;
		END
		$$;
	`)
}

type ResultProductSearch struct {
	Product_Name  string
	Product_Price int
}

func ExecProductSearch(startWith string) ResultProductSearch {
	var result ResultProductSearch

	db := _postgres.GetPostgres()
	query := fmt.Sprintf(`CALL pro_get_products_by_name('%s%s');`, startWith, "%")
	db.Raw(query).Scan(&result)

	fmt.Println("pro_get_products_by_name", result)
	return result
}
