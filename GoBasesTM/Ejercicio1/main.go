package main
import (
	"fmt"
	"os"
)

func nuevoProd(id int, precio float64, cantidad int) {
	val := fmt.Sprintf("\n%d; %.2f; %d;", id, precio, cantidad)
	//d1 := []byte(val)

	f, err := os.OpenFile("../prueba.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(val); err != nil {
		fmt.Println(err)
	}
	
}

func main() {

	nuevoProd(1, 10, 50)
	nuevoProd(2, 45, 34)

}