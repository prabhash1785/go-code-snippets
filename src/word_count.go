// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// @author Prabhash Rathore

package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	
	word_map := make(map[string]int)
	
	for i := 0; i < len(words); i++ {
		v := word_map[words[i]]
		
		if v != 0 {
			word_map[words[i]] = v + 1			
		} else {
			word_map[words[i]] = 1	
		}	
	}
	
	return word_map
}

func main() {
	fmt.Println(WordCount("A man a plan a canal panama."))

	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}

