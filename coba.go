package main

import (
	"fmt"
)

type Hewan interface {
	Suara()
	Makan()
	SetSuara(suara string)
}

type Sapi struct {
	suara   string
	makanan string
}

func (s Sapi) Suara() {
	fmt.Println("Suara sapi: ", s.suara)
}

func (s *Sapi) SetSuara(suara string) {
	s.suara = suara
}

func (s Sapi) Makan() {

	fmt.Println("Sapi memakan: ", s.makanan)
}

// func main() {
// 	// var sapi Hewan = &Sapi{suara: "Moo", makanan: "Daging"}
// 	// sapi1 := sapi

// 	// sapi.Suara()
// 	// sapi.SetSuara("Mooo")
// 	// sapi1.Suara()
// 	text := "this is a test string this is a string"
// 	words := strings.Fields(text) // Splits the string into a slice of words

// 	wordCount := map[string]int{
// 		"asu": 0,
// 	}

// 	for _, word := range words {
// 		wordCount[word]++ // Increment the count for each word
// 	}

// 	for word, count := range wordCount {
// 		fmt.Printf("Word: %s, Count: %d\n", word, count)
// 	}

// }

func bunyiKan[T Hewan](object T) T {
	object.Suara()

	return object
}

type Kucing[T any] struct {
	attribute []T
}

type AttrKucing struct {
	nama    string
	umur    int
	pemilik string
}

func main() {
	// sapi := &Sapi{suara: "Moo", makanan: "Daging"}

	attr := []AttrKucing{
		{
			nama:    "kucing 1",
			umur:    2,
			pemilik: "budi",
		},
		{
			nama:    "kucing 2",
			umur:    2,
			pemilik: "budi",
		},
	}

	fmt.Println(len(attr), cap(attr))

	kucing := &Kucing[AttrKucing]{attr}

	for i, v := range kucing.attribute {
		fmt.Println("=======================")
		fmt.Println("name: ", v.nama)
		fmt.Println("umur: ", v.umur)
		fmt.Println("pemilik: ", v.pemilik)

		if i == len(kucing.attribute)-1 {
			fmt.Println("=======================")
		}
	}

}
