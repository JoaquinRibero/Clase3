package main

import "fmt"

type Productos struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Servicios struct {
	Nombre  string
	Precio  int
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio int
}

func sumarProductos(prod []Productos) <-chan int {
	r := make(chan int)
	go func() {
		sum := 0
		for _, p := range prod {
			sum += p.Precio * p.Cantidad
		}

		r <- sum
	}()

	return r
}

func sumarServicios(servicios []Servicios) <-chan int {
	r := make(chan int)
	go func() {
		sum := 0
		for _, s := range servicios {
			sum += s.Precio * (s.Minutos / 30)
			if s.Minutos%30 > 0 {
				sum += s.Precio
			}
		}
		r <- sum
	}()

	return r
}

func sumarMantenimiento(mantenimiento []Mantenimiento) <-chan int {
	r := make(chan int)
	go func() {
		sum := 0
		for _, m := range mantenimiento {
			sum += m.Precio
		}
		r <- sum
	}()

	return r
}

func main() {

	var p1 Productos = Productos{"Celular", 1500, 10}
	var p2 Productos = Productos{"Tablet", 1000, 5}
	var p3 Productos = Productos{"PC", 5000, 8}
	var p4 Productos = Productos{"Notebook", 5500, 7}

	var prod []Productos = []Productos{p1, p2, p3, p4}

	var s1 Servicios = Servicios{"Limpieza", 1000, 45}
	var s2 Servicios = Servicios{"Formateo", 500, 60}
	var s3 Servicios = Servicios{"Instalacion", 3000, 90}

	var serv []Servicios = []Servicios{s1, s2, s3}

	var m1 Mantenimiento = Mantenimiento{"Mantenimiento", 3000}
	var m2 Mantenimiento = Mantenimiento{"Proteccion", 2500}

	var mant []Mantenimiento = []Mantenimiento{m1, m2}

	p, s, m := sumarProductos(prod), sumarServicios(serv), sumarMantenimiento(mant)

	sum := <-p + <-s + <-m

	fmt.Println("La suma total es: ", sum)

}
