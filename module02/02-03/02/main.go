/* Задание:
нужно объединить слайс с выходными днями и слайс с рабочими в один слайс. Выведите на экран итоговый слайс с днями.
*/
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
