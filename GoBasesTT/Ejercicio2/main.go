package main

import "fmt"

type usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []producto
}

type producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

func nuevoProd(nombre string, precio int) producto {
	var p producto
	p.Nombre = nombre
	p.Precio = precio

	return p
}

func (u *usuario) agregarProducto(p *producto, cantidad int) {
	u.Productos = append(u.Productos, *p)
	p.Cantidad = cantidad
}

func (u *usuario) borrarProductos() {
	u.Productos = []producto{}
}

func main() {

	var u usuario
	u.Nombre = "Joaquin"
	u.Apellido = "Ribero"
	u.Correo = "mail@mail.com"

	p := nuevoProd("Celular", 1500)

	u.agregarProducto(&p, 10)
	fmt.Println(u)
	fmt.Println(p)

}
