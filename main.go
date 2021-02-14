package main

import "fmt"

func main() {

	user := NewUser("Sondre", 20, true, []string{"english", "german"})
	job, ok:= EvaluateApplicant(user)
	if ok {
		fmt.Println(job)
		fmt.Println(user)
		return
	}
	fmt.Println("No job found")

}
