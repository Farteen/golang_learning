// package main

// import "fmt"

// func main() {
// 	// fmt.Printf("this repository is related to learning golang\n")
// 	// goku := Saiyan{}
// 	// goku = Saiyan{ Name: "Goku"}
// 	// goku.Power = 9000
// 	// goku := &Saiyan{"Goku", 9000}
// 	// Super(goku)
// 	// goku.Super()
// 	// fmt.Println(goku.Power)

// 	// goku := new(Saiyan)
// 	// goku := &Saiyan{}

// 	// goku := &Saiyan {
// 	// 	Person: &Person {"Goku"},
// 	// 	Power: 9001,
// 	// }
// 	// goku.Introduce()

// 	goku := &Saiyan{
// 		Person: &Person{"Goku"},
// 	}
// 	fmt.Println(goku.Name)
// 	// fmt.Println(goku.Person.Name)
// 	goku.Introduce()
// 	goku.Person.Introduce()
// }

// type Person struct {
// 	Name string
// }

// func (p *Person) Introduce() {
// 	fmt.Printf("Hi, I'm %s \n", p.Name)
// }

// func (s *Saiyan) Introduce() {
// 	fmt.Printf("Hi, I'm %s. Ya!\n",s.Name)
// }

// // func Super(s *Saiyan) {
// // 	s.Power += 10000
// // }

// // func (s *Saiyan) Super() {
// // 	s.Power += 10000
// // }

// // func NewSaiyan(name string, power int) *Saiyan {
// // 	return &Saiyan {
// // 		Name: name,
// // 		Power: power,
// // 	}
// // }

// // func log(message string) {

// // }

// // func add(a int, b int) int {

// // }

// // func power(name string) (int, bool) {
// // 	return 1, true
// // }

// // value, exist := power("goku")

// // if exists == false {
// // 	// 
// // }

// // type Saiyan struct {
// // 	Name string
// // 	Power int
// // }

// type Saiyan struct {
// 	*Person
// 	Power int
// }