package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Wish struct {
	childName string
	wish      string
}

type Present string

func main() {
	wishes := make(chan Wish, 1)      // HL
	presents := make(chan Present, 1) // HL

	go santa(wishes, presents) // HL

	elves := []string{"Алабастър", "Буши", "Пепър", "Шайни", "Шугърплъм"}
	for _, elfName := range elves {
		go elf(elfName, wishes, presents) // HL
	}
	time.Sleep(time.Duration(3) * time.Second)
}

func santa(wishes chan<- Wish, presents <-chan Present) { // HL
	wishlist := []Wish{
		Wish{childName: "Пешко", wish: "колело"},
		Wish{childName: "Гошко", wish: "iPhone7"},
		Wish{childName: "Ники", wish: "базука"},
	}
	go func() {
		for _, wish := range wishlist {
			wishes <- wish // HL
		}
	}()
	for _, wish := range wishlist {
		present := <-presents // HL
		fmt.Printf("Дядо Коледа подари %s на %s\n", present, wish.childName)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
}

func elf(name string, wishes <-chan Wish, presents chan<- Present) { // HL
	for wish := range wishes { // HL
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("%s изработи %s\n", name, wish.wish)
		// Present ready
		presents <- Present(wish.wish) // HL
	}
}
