/*
Оператор defer откладывает выполнение функции до того момента,
как произойдет возврат из окружающей функции.

Аргументы отложенных вызовов вычисляются сразу же,
но вызов функции не происходит до того,
как произойдет возврат из окружающей функции.
*/

package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}