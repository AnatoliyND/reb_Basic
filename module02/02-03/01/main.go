package main

import "fmt"

func main() {
	week := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "suturday", "sunday"}
	workDay := make([]string, 5)
	copy(workDay, week)
	fmt.Println("Работчие дни недели:\n", workDay)
	week = append(week[5:])
	fmt.Println("Выходные дни недели:\n", week)
}
