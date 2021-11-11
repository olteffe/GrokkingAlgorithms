package main

import (
	"fmt"
	"github.com/gammazero/deque"
	"strings"
)

type inputMap map[string][]string

func personIsSeller(name string) bool {
	return strings.HasSuffix(name, "m")
}

func breadthSearch(graph inputMap, name string) bool {
	var queue deque.Deque
	for _, value := range graph[name] {
		queue.PushBack(value)
	}
	searched := make(map[string]bool)
	for queue.Len() > 0 {
		person := fmt.Sprintf("%v", queue.PopFront())
		if _, ok := searched[person]; !ok {
			if personIsSeller(person) {
				fmt.Printf("%s is a mango seller!", person)
				return true
			} else {
				for _, value := range graph[person] {
					queue.PushBack(value)
				}
				searched[person] = true
			}
		}
	}
	return false
}

func main() {
	name := "you"
	graph := inputMap{"you": {"alice", "bob", "claire"}, "bob": {"anuj", "peggy"}, "alice": {"peggy"},
		"claire": {"thom", "jonny"}, "anuj": {}, "peggy": {}, "thom": {}, "jonny": {}}
	breadthSearch(graph, name)
}
