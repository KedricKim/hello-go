package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ked/learngo/accounts"
)

func main() {
	// basicPratice()

	account := accounts.NewAccount("ked")
	account.Deposit(1000)
	err := account.WithDraw(9900)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.ShowDeposit())
}

func basicPratice() {
	fmt.Println(multiple(3, 7))

	leng, name := lenAndString("kedric")
	fmt.Println(leng, name)

	total := forLoop(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	fmt.Println(canIDrink(20))

	a := 10
	b := &a
	*b = 20
	fmt.Println(a)

	favFood := []string{"ramen", "beef"}
	whoAmI := person{name: "kedric", age: 34, favFood: favFood}
	fmt.Println(whoAmI)
	fmt.Println(whoAmI.name)
}

func multiple(a int, b int) int {
	return a * b
}

func lenAndString(name string) (int, string) {
	defer fmt.Println("Hola!!")
	fmt.Println(name)
	return len(name), strings.ToUpper(name)
}

func forLoop(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) string {
	if koreanAge := age - 2; koreanAge <= 18 {
		return "no u cant"
	}
	return "of cource drink it !"
}

type person struct {
	name    string
	age     int
	favFood []string
}
