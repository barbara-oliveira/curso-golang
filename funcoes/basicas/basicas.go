package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println(p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "R1", "R2"
}

func main() {
	f1()
	f2("Julie", "Luke")
	fmt.Println(f3())
	fmt.Println(f4("Oieee", "Tchaaau"))
	fmt.Println(f5())

	r3, r4 := f3(), f4("P1", "P2")
	fmt.Println(r3, r4)

	r51, r52 := f5()
	fmt.Println(r51, r52)

	r512, _ := f5() // ignorar um dos retornos da função
	fmt.Println(r512)
}
