package main

import "fmt"

func main() {
  
  /*
  Для создания небуферизированного канала вызывается функция make() без указания емкости канала:
 */
 
	intCh := make(chan int)
	/*
Если канал пустой, то горутина-получатель блокируется, пока в канале не окажутся данные. Когда горутина-отправитель посылает данные, горутина-получатель получает эти данные и возобновляет работу.

Горутина-отправитель может отправлять данные только в пустой канал. Горутина-отправитель блокируется до тех пор, пока данные из канала не будут получены. Например:
  */

	go func() {
		fmt.Println("Go routine starts")
		
		/*
   Через небуферизированный канал intCh 
   горутина, представленная анонимно функцией, передает число 5:
   */
   
		intCh <- 5 //блокировка,пока данные не будут получены функцией main
	}()
	fmt.Println(<-intCh) //получение данных из канала
	fmt.Println("The End")
}
