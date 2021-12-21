package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Insercion(vector []int) <-chan int64 {
	r := make(chan int64)
	go func() {
		inicio := time.Now().UnixMicro()
		for i := range vector {
			actual := vector[i]
			j := i
			for j > 0 && vector[j-1] > actual {
				vector[j] = vector[j-1]
				j = j - 1
			}
			vector[j] = actual
		}
		fin := time.Now().UnixMicro()
		r <- fin - inicio
	}()

	return r
}

func Burbuja(vector []int) <-chan int64 {
	r := make(chan int64)
	go func() {
		inicio := time.Now().UnixMicro()
		iteracion := 1
		permutatuion := false
		for permutatuion != true {
			for i := 0; i < len(vector)-iteracion; i++ {
				if vector[i] > vector[i+1] {
					temp := vector[i]
					vector[i] = vector[i+1]
					vector[i+1] = temp
					permutatuion = true
				}
			}
			iteracion += 1
		}
		fin := time.Now().UnixMicro()
		r <- fin - inicio
	}()

	return r
}

func seleccion(vector []int) <-chan int64 {
	r := make(chan int64)
	go func() {
		inicio := time.Now().UnixMicro()
		for i := 0; i < len(vector)-1; i++ {
			min := i
			for j := i + 1; j < len(vector); j++ {
				if vector[j] < vector[min] {
					min = j
				}
			}
			temp := vector[i]
			vector[i] = vector[min]
			vector[min] = temp
		}
		fin := time.Now().UnixMicro()
		r <- fin - inicio
	}()

	return r
}

func main() {
	i100 := rand.Perm(100)
	i1000 := rand.Perm(1000)
	i10000 := rand.Perm(10000)

	b100 := make([]int, 100)
	copy(b100, i100)
	b1000 := make([]int, 1000)
	copy(b1000, i1000)
	b10000 := make([]int, 10000)
	copy(b10000, i10000)

	s100 := make([]int, 100)
	copy(s100, i100)
	s1000 := make([]int, 1000)
	copy(s1000, i1000)
	s10000 := make([]int, 10000)
	copy(s10000, i10000)

	i, b, s := Insercion(i100), Burbuja(b100), seleccion(s100)
	fmt.Println("El tiempo de Inserción para 100 es: ", <-i)
	fmt.Println("El tiempo de Burbuja para 100 es: ", <-b)
	fmt.Println("El tiempo de Selección para 100 es: ", <-s)
	fmt.Println("------------------------------------------")
	i, b, s = Insercion(i1000), Burbuja(b1000), seleccion(s1000)
	fmt.Println("El tiempo de Inserción para 1000 es: ", <-i)
	fmt.Println("El tiempo de Burbuja para 1000 es: ", <-b)
	fmt.Println("El tiempo de Selección para 1000 es: ", <-s)
	fmt.Println("------------------------------------------")
	i, b, s = Insercion(i10000), Burbuja(b10000), seleccion(s10000)
	fmt.Println("El tiempo de Inserción para 10000 es: ", <-i)
	fmt.Println("El tiempo de Burbuja para 10000 es: ", <-b)
	fmt.Println("El tiempo de Selección para 10000 es: ", <-s)
	fmt.Println("------------------------------------------")
}
