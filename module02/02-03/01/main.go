/* Задание:
•	из слайса с днями недели надо скопировать в новый слайс рабочие дни, а из исходного слайса удалить скопированные, чтобы остались только выходные дни.
•	Выведите на экран слайсы с выходными и рабочими днями недели.
*/
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
