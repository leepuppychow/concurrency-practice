package daily_walk

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name     string
	TaskDone chan bool
}

type Alarm struct {
	Delay time.Duration
	Done  chan bool
}

func randomDelay(min, max int) int {
	return rand.Intn(max-min) + min
}

func Runner() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Let's go for a walk!")

	aliceCh := make(chan bool)
	bobCh := make(chan bool)
	alice := Person{Name: "Alice", TaskDone: aliceCh}
	bob := Person{Name: "Bob", TaskDone: bobCh}

	go alice.Prepare()
	go bob.Prepare()

	<-alice.TaskDone
	<-bob.TaskDone

	alarmCh := make(chan bool)
	alarm := Alarm{Delay: 60, Done: alarmCh}
	go alarm.setAlarm()
	go alice.PutOnShoes()
	go bob.PutOnShoes()

	<-alice.TaskDone
	<-bob.TaskDone
	fmt.Println("Exiting and locking the door.")

	<-alarm.Done
}

func (p *Person) Prepare() {
	fmt.Printf("%s started getting ready.\n", p.Name)
	timeSpent := randomDelay(60, 90)
	time.Sleep(time.Duration(timeSpent) * time.Second)
	fmt.Printf("%s spent %d seconds getting ready.\n", p.Name, timeSpent)
	p.TaskDone <- true
}

func (p *Person) PutOnShoes() {
	fmt.Printf("%s started putting on shoes.\n", p.Name)
	timeSpent := randomDelay(35, 45)
	time.Sleep(time.Duration(timeSpent) * time.Second)
	fmt.Printf("%s spent %d seconds putting on shoes.\n", p.Name, timeSpent)
	p.TaskDone <- true
}

func (a *Alarm) setAlarm() {
	fmt.Printf("Arming alarm.\n")
	fmt.Printf("Alarm is counting down.\n")
	time.Sleep(a.Delay * time.Second)
	fmt.Printf("Alarm is armed.\n")
	a.Done <- true
}
