package main

import "fmt"

func main() {
	tasks := [5]string{"wakeup", "Study", "Excersize", "Learning new skill", "sleep"}
	fmt.Println("Lucky's TO-DO LIST")
	for i := 0; i < len(tasks); i++ {
		fmt.Println(i+1, "-", tasks[i])
	}
}
