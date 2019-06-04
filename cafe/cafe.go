package cafe

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Computer struct {
	ID       int
	Users    []*Tourist
	Occupied bool
}

type Tourist struct {
	Name     string
	Computer *Computer
}

func randomNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func getTouristQueue(total int) []*Tourist {
	tourists := []*Tourist{}
	for i := 1; i <= total; i++ {
		tourists = append(tourists, &Tourist{
			Name:     fmt.Sprintf("Tourist %d", i),
			Computer: nil,
		})
	}
	return tourists
}

func getComputers(total int) []*Computer {
	computers := []*Computer{}
	for i := 1; i <= total; i++ {
		computers = append(computers, &Computer{
			ID:       i,
			Users:    []*Tourist{},
			Occupied: false,
		})
	}
	return computers
}

func Runner() {
	tourists := getTouristQueue(25)
	computers := getComputers(8)
	computerFree := make(chan *Computer)

	// Seat the first users of the day
	for _, c := range computers {
		tourists = c.seatUser(tourists, computerFree)
	}

	for len(tourists) > 0 {
		select {
		case computer := <-computerFree:
			tourists = computer.seatUser(tourists, computerFree)
		}
	}

	time.Sleep(2 * time.Second) // TODO: fix this code smell, wait until all computers are unoccupied then exit
	fmt.Println("All done for the day")
	for _, c := range computers {
		c.PrintUsersForTheDay()
	}
}

func (t *Tourist) startBrowsing(compFree chan *Computer) {
	timeSpent := randomNum(15, 120)
	fmt.Printf("%s is online.\n", t.Name)
	time.Sleep(time.Duration(timeSpent) * time.Millisecond) // Change to Minutes for real simulation
	fmt.Printf("%s is done having spent %d minutes online.\n", t.Name, timeSpent)
	t.Computer.Occupied = false
	compFree <- t.Computer
}

func (c *Computer) seatUser(tourists []*Tourist, compFree chan *Computer) []*Tourist {
	ch := make(chan []*Tourist)

	go func() {
		c.Occupied = true
		c.Users = append(c.Users, tourists[0])
		tourists[0].Computer = c
		go tourists[0].startBrowsing(compFree)
		ch <- tourists[1:]
	}()

	return <-ch
}

func (c *Computer) PrintUsersForTheDay() {
	usersNames := []string{}
	for _, u := range c.Users {
		usersNames = append(usersNames, u.Name)
	}
	fmt.Printf("  Computer %d's users for the day: %s\n", c.ID, strings.Join(usersNames, ", "))
}
