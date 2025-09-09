package main

import "fmt"

func main() {
	movielist := []string{}
	movielist = append(movielist, "ABCD")
	movielist = append(movielist, "ABCD2")
	movielist = append(movielist, "ABCD3")
	fmt.Println("MOvielist name :")
	for i := 0; i < len(movielist); i++ {
		fmt.Println(i+1, "--", movielist[i])
	}

}
