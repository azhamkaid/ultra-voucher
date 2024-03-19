package main

import (
	"fmt"
)

func main() {
	name := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := Anagrams(name)
	fmt.Println(result)
	fmt.Println(result)
	for _, group := range result {
		fmt.Println(group)
	}

}

func Anagrams(name []string) [][]string {
	//membuat map dengan isi string array untuk anagram yang akan di kelompokan
	anagramMap := make(map[string][]string)
	//melakukan loop untuk mencari kata anagram dengan menggunakan rune
	for _, str := range name {
		//string menjadi array berisi angka menggunakan rune
		runes := []rune(str)
		//mengurutkan angka dalam runes menggunakan func sort manual
		sortRune := sort(runes)
		fmt.Println(sortRune)
		//merubah string angka menjadi string kembali
		sortedstring := string(sortRune)
		//memasukan string dengan runes yang sama kedalam anagramMap
		//setiap terjadi looping jika ada kata dengan sorted rune yang sama maka kata tersebut akan ditambahkan
		anagramMap[sortedstring] = append(anagramMap[sortedstring], str)

	}
	//array 2d
	//conversi anagram dalam bentuk mep menjadi bentuk array 2d yang lebih sederhana
	result := make([][]string, 0, len(anagramMap))
	for _, groupedname := range anagramMap {
		result = append(result, groupedname)
	}
	return result
}
func sort(numbers []int32) []int32 {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}
