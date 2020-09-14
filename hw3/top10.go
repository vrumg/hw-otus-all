package top10

import (
	"sort"
	"strings"
)

//Top10 get top 10 words in text
func Top10(text string) []string {

	//split text to word array
	rawArray := strings.Split(text, " ")

	//prepare hash table
	hash := map[string]int{}

	//iterate word array and write each word as key
	for _, val := range rawArray {
		//init new key with 1 match
		if _, found := hash[val]; !found {
			hash[val] = 1
			continue
		}
		//iterate values for subsequent matches
		hash[val]++
	}

	//define and init struct for sorting
	type sortStruct struct {
		key   string
		value int
	}
	mapSortArray := []sortStruct{}

	//copy hash table to sort array
	for key, val := range hash {
		mapSortArray = append(mapSortArray, sortStruct{
			key:   key,
			value: val,
		})
	}

	//define and apply sorting condition for custom sortStruct
	sort.SliceStable(mapSortArray, func(i, j int) bool {
		return mapSortArray[i].value > mapSortArray[j].value
	})

	//init slice to return
	retSlice := []string{}

	//fill string array with first X elements of sorted array of custom struct
	for _, val := range mapSortArray[:10] {
		retSlice = append(retSlice, val.key)
	}

	//bingo
	return retSlice
}
