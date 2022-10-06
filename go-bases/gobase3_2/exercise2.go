
package main

import "fmt"

type User struct {
	name string
	lastName string
	email string
	products []Product
}

type Product struct {
	name string
	price float64
	quantity int
}

func newProduct(name string, price float64) (Product){
	product := Product {
		name,
		price,
		0,
	}

	return product
}

func addProduct(user *User, product *Product, quantity int) ([]Product){
	(*product).quantity = quantity
	s := append((*user).products, *product)
	return s
}

func deleteProduct(user *User) {
	(*user).products = nil
}

func main() {
	var testProducts []Product
	testUser := User {
		"Adriana",
		"Echeverri Romero",
		"adrianaer98@hotmail.com",
		testProducts,
	}

	product := newProduct("Buds", 890)

	fmt.Println(product)

	addProduct(&testUser, &product, 5)

	fmt.Println(testUser)
	
	deleteProduct(&testUser)
	
	fmt.Println(testUser)
}



// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
// Se necesitan las estructuras:
// Usuario: Nombre, Apellido, Correo, Productos (array de productos).
// Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
// Nuevo producto: recibe nombre y precio, y retorna un producto.
// Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
// Borrar productos: recibe un usuario, borra los productos del usuario.
