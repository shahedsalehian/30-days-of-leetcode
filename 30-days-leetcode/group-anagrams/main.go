package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	temp := groupAnagrams(strs)
	for _, val := range temp {
		fmt.Println(val)
	}
}

func groupAnagrams(strs []string) [][]string {
	temp := [][]string{}
	tempMap := make(map[string][]string)
	for _, str := range strs {
		sortedString := SortString(str)
		tempMap[sortedString] = append(tempMap[sortedString], str)
	}

	for _, value := range tempMap {
		temp = append(temp, value)
	}
	
	return temp
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
