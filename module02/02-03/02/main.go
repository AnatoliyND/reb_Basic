package main

import "fmt"

func main() {
	weekend := []string{"suturday", "sunday"}
	workDay := []string{"monday", "tuesday", "wednesday", "thursday", "friday"}
	week := []string{}
	/*weekend = append(weekend, workDay...)
	fmt.Println(weekend)*/
	week = append(week, workDay...)
	week = append(week, weekend...)
	fmt.Println(week)
}
