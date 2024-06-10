package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var counter map[string]int
	var text string
	re := regexp.MustCompile(`[[:punct:]]`) // https://ru.stackoverflow.com/questions/723860/как-убрать-спецсимволы-в-тексте-golang

	fmt.Println("Введите текст")
	fmt.Fscan(os.Stdin, &text) //получение текста

	text = strings.ToLower(text)         //обработка исходного текста
	text = re.ReplaceAllString(text, "") //обработка исходного текста

	slText := strings.Split(text, " ") //разбивка текста на слова и сохранение в слайс
	var unicWords []string

	for _, word := range slText { //цикл, осуществляющий отбор уникальных слов
		for _, wordUnic := range unicWords {
			if word != wordUnic {
				unicWords = append(unicWords, word)
			}
		}
	}

	for _, word := range unicWords { //подсчёт вхождений с занесением в мапу
		counter[word] = strings.Count(text, word)
	}

	for key, amount := range counter { //вывод мапы
		fmt.Printf("Слово %q повторялось в тексте %q раз(а)", key, amount)
	}

}
