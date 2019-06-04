package tapas

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Person struct {
	Name string
	Done chan *Person
}

type Dish struct {
	Name    string
	Morsels int
	mu      sync.Mutex
}

func randomNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Note: return slice of Dish pointers because you should not copy the sync.Mutex
func OrderDishes() []*Dish {
	d1 := Dish{Name: "chorizo", Morsels: randomNum(5, 11)}
	d2 := Dish{Name: "chopitos", Morsels: randomNum(5, 11)}
	d3 := Dish{Name: "pimientos de padrón", Morsels: randomNum(5, 11)}
	d4 := Dish{Name: "croquetas", Morsels: randomNum(5, 11)}
	d5 := Dish{Name: "patatas bravas", Morsels: randomNum(5, 11)}
	return []*Dish{&d1, &d2, &d3, &d4, &d5}
}

func Runner() {
	dishes := OrderDishes()
	done := make(chan *Person)
	foodLeft := true
	foodCh := make(chan bool)

	alice := Person{Name: "Alice", Done: done}
	bob := Person{Name: "Bob", Done: done}
	charlie := Person{Name: "Charlie", Done: done}
	dave := Person{Name: "Dave", Done: done}

	fmt.Println("Bon appétit!")
	printDishes(dishes)

	go alice.EatMorsel(dishes)
	go bob.EatMorsel(dishes)
	go charlie.EatMorsel(dishes)
	go dave.EatMorsel(dishes)

	for foodLeft {
		select {
		case person := <-done:
			go person.EatMorsel(dishes)
			go checkFoodLeft(dishes, foodCh)
		case foodLeft = <-foodCh:
		}
	}
	fmt.Println("That was delicious!")
	printDishes(dishes)
}

func checkFoodLeft(dishes []*Dish, ch chan bool) {
	for _, d := range dishes {
		if d.Morsels > 0 {
			ch <- true
			return
		}
	}
	ch <- false
}

func (p *Person) EatMorsel(dishes []*Dish) {
	// Pick random dish:
	dish := &Dish{Name: "", Morsels: 0}
	for dish.Morsels <= 0 {
		i := randomNum(0, len(dishes))
		dish = dishes[i]
	}

	// Decrement Morsels, protect shared var with mutex
	dish.mu.Lock()
	dish.Morsels--
	dish.mu.Unlock()

	fmt.Printf("\t%s is enjoying some %s (morsels = %d).\n", p.Name, dish.Name, dish.Morsels)
	// timeSpent := randomNum(30, 180)
	timeSpent := randomNum(3, 5)
	time.Sleep(time.Duration(timeSpent) * time.Second)

	p.Done <- p
}

// For debugging purposes
func printDishes(dishes []*Dish) {
	for _, d := range dishes {
		fmt.Println(d)
	}
}
