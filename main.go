package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ked/learngo/accounts"

	"github.com/ked/learngo/mydict"
)

func main() {
	// basicPratice()
	// methodPractice()
	// mapPractice()

}

func mapPractice() {
	dictionary := mydict.Dictionary{"first": "First word"}
	defunition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(defunition)
	}

	dictionary.Add("second", "hello")

	newWord := "third"
	_, err2 := dictionary.Search(newWord)
	if err2 != nil {
		dictionary.Add(newWord, "hello")
	} else {
		dictionary.Update(newWord, "hi")
	}

	fmt.Println(dictionary)

	result, err3 := dictionary.Delete(newWord)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(result)
	}

	fmt.Println(dictionary)
}

func methodPractice() {
	account := accounts.NewAccount("ked")
	account.Deposit(100)
	err := account.WithDraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.ShowDeposit(), account.Owner())

	fmt.Println(account)
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

// defer 은 func이 다 실행되고 나서 실행
func lenAndString(name string) (int, string) {
	defer fmt.Println("Hola!!")
	fmt.Println(name)
	return len(name), strings.ToUpper(name)
}

// for index, item
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

// account 와 다르게 같은 패키지 내에서는 소문자
type person struct {
	name    string
	age     int
	favFood []string
}
