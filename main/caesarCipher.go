package main

import (
	"fmt"
)

func main() {

	//Инициализация и ввод шифруемого сообщения
	var unenc string
	fmt.Print("Enter message (without spaces): ")
	fmt.Scan(&unenc)

	//Инициализация и ввод значения сдвига
	var key byte
	fmt.Print("Enter key: ")
	fmt.Scan(&key)

	//Инициализация среза для хранения зашифрованного сообщения
	enc := make([]byte, len(unenc))

	//Замена символов сообщения на символы со сдвигом по ключу
	for i := 0; i < len(unenc); i++ {
		//Проверка на принадлежность символа к нижнему регистру через ASCII код.
		if unenc[i] > 64 && unenc[i] < 91 {
			//Если зашифрованый символ выходит за границу алфавита, сдвиг продолжается в его начале и
			//в зашифрованное сообщение по индексу записывается шифруемый символ со сдвигом.
			//Если зашифрованый символ не выходит за границу алфавита, сдвиг не переносится.
			//Аналогично для символов с верхним регистром.
			if unenc[i]+key > 90 {
				enc[i] = unenc[i] + key - 26
			} else {
				enc[i] = unenc[i] + key
			}
		}
		//Проверка на принадлежность символа к верхнему регистру через ASCII код
		if unenc[i] > 96 && unenc[i] < 123 {
			if unenc[i]+key > 122 {
				enc[i] = unenc[i] + key - 26
			} else {
				enc[i] = unenc[i] + key
			}
		}
	}

	//Вывод зашифрованного сообщения
	fmt.Print("Encrypted message: ")
	fmt.Println(string(enc))
}
