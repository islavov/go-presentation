package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	wishes := make(chan string, 10)   // HL
	presents := make(chan string, 10) // HL

	go santa(wishes, presents) // HL

	elves := []string{"Алабастър", "Буши", "Пепър", "Шайни", "Шугърплъм"}
	for _, elfName := range elves {
		go elf(elfName, wishes, presents) // HL
	}
	time.Sleep(time.Duration(3) * time.Second)
}

func santa(wishes chan<- string, presents <-chan string) { // HL
	wishlist := map[string]string{"Пешко": "колело", "Гошко": "iPhone7", "Ники": "базука"}
	for _, wish := range wishlist {
		wishes <- wish // HL
	}
	for child := range wishlist {
		present := <-presents // HL
		fmt.Printf("Дядо Коледа подарява %s на %s\n", present, child)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
}

func elf(name string, wishes <-chan string, presents chan<- string) { // HL
	for wish := range wishes { // HL
		fmt.Printf("%s изработва %s\n", name, wish)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		// Present ready
		presents <- wish // HL
	}
}
