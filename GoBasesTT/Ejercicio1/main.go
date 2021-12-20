package main

import "fmt"

type usuario struct {
	Nombre string
	Apellido string
	Edad int
	Correo string
	Contrasena string
}

func (u *usuario) setNombre(nombre string) {
	u.Nombre = nombre
}

func (u *usuario) setApellido(apellido string) {
	u.Apellido = apellido
}

func (u *usuario) setEdad(edad int) {
	u.Edad = edad
}

func (u *usuario) setContrasena(contrasena string) {
	u.Contrasena = contrasena
}

func (u *usuario) setCorreo(correo string) {
	u.Correo = correo
}

func main() {

	var u usuario
	u.Nombre = "Joaquin"
	u.Apellido = "Ribero"
	u.Edad = 25
	u.Correo = "mail@mail.com"
	u.Contrasena = "contraseña"

	fmt.Println(u)
	u.setNombre("Tony")
	u.setApellido("Stark")

	fmt.Println(u)


}