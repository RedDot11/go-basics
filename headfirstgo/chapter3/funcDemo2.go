/*
Программа,определяющая сдал ли студент экзамен.
Проходной балл равен 60,если студент получает за экзамен 60 и более баллов,то он прошел экзамен,о чем сообщит программа,если же набрано менее 60 баллов,то экзамен считается проваленным.
*/
package main

import "fmt"

func status(grade float64) string { //Функция, определяющая результат экзамена(объявлена переменная grade типа float64,являющая собой оценку за экзамен).Возвращаемое значение имеет  тип string
	if grade < 60.0 { //если grade меньше 60...
		return "failing" //...функция возвращает failing(провал)
	}
	return "passing" //в остальных случаях функция возращает passing(успех)
}

func main() {
	fmt.Println("You're", status(65.0)) //Вывод на экран сообщения о результате сдачи экзамена(в данном случае экзамен пройден,т.к.  получено 65 баллов,что выше минимального порога в 60 баллов)
	fmt.Println("You're", status(59))   //Вывод на экран сообщения о результате сдачи экзамена(в данном случае экзамен не пройден,т.к. получено 59 баллов,что ниже минимального порога в 60 баллов)
}
