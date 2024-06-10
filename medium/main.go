package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	names := []string{"Маша", "Саша", "Боря", "Миша", "Настя", "Коля", "Серёжа"}
	var letter string
	flag := false

	fmt.Println("Введите букву")
	fmt.Fscan(os.Stdin, &letter)

	letter = strings.ToUpper(letter)

	for _, name := range names {
		if strings.Contains(name, letter) {
			flag = true
			fmt.Println(name)
		}
	}

	if !flag {
		fmt.Printf("В списке нет имени, начинающегося с буквы %s.", letter)
	}
}
