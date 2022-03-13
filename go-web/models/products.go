package models

import "github.com/formacao-go/go-web/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func GetAllProducts() []Product {
	db := db.ConnectionDB()

	selectProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func Create(name, description string, price float64, quantity int) {
	db := db.ConnectionDB()

	insertDataInDB, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertDataInDB.Exec(name, description, price, quantity)

	defer db.Close()
}

func Delete(id string) {
	db := db.ConnectionDB()

	delete, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func Edit(id string) Product {
	db := db.ConnectionDB()

	productByID, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdated := Product{}

	for productByID.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productByID.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		productToUpdated.Id = id
		productToUpdated.Name = name
		productToUpdated.Description = description
		productToUpdated.Price = price
		productToUpdated.Quantity = quantity
	}

	defer db.Close()

	return productToUpdated
}

func Update(id int, name string, description string, price float64, quantity int) {
	db := db.ConnectionDB()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
