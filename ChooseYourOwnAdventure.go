package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Scene struct {
	Text    string
	Choices map[string]int
}

func main() {
	scenes := []Scene{
		{Text: "You wake up in a forest. Go north or south?", Choices: map[string]int{"north": 1, "south": 2}},
		{Text: "You find a river. Swim or return?", Choices: map[string]int{"swim": 3, "return": 0}},
		{Text: "You meet a wolf! Run or fight?", Choices: map[string]int{"run": 0, "fight": 4}},
		{Text: "You swim and find treasure! THE END.", Choices: map[string]int{}},
		{Text: "The wolf is strong! You are defeated. THE END.", Choices: map[string]int{}},
	}
	cur := 0
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(scenes[cur].Text)
		if len(scenes[cur].Choices) == 0 {
			break
		}
		fmt.Print("Choice? ")
		ch, _ := r.ReadString('\n')
		ch = strings.TrimSpace(strings.ToLower(ch))
		if idx, ok := scenes[cur].Choices[ch]; ok {
			cur = idx
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}