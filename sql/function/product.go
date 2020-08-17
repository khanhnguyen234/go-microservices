package function

import (
	"fmt"
	"khanhnguyen234/sql/_postgres"
)

func ProductSearch() {
	db := _postgres.GetPostgres()
	db.Exec(`
		create or replace function func_get_products_by_name (
		  p_pattern varchar
		) 
			returns table (
				product_name text,
				product_price int
			) 
			language plpgsql
		as $$
		begin
			return query 
				select
					name,
					price
				from
					product_models
				where
					name ilike p_pattern;
		end;$$
	`)
}

type ResultProductSearch []struct {
	Product_Name  string
	Product_Price int
}

func ExecProductSearch(startWith string) ResultProductSearch {
	var result ResultProductSearch

	db := _postgres.GetPostgres()
	query := fmt.Sprintf(`SELECT * FROM func_get_products_by_name ('%s%s');`, startWith, "%")
	db.Raw(query).Scan(&result)

	fmt.Println("func_get_products_by_name", result)
	return result
}
