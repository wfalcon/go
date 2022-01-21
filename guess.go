//giess - игра, в которой игрок должен угадать задуманное число
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	success := false
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я загадал случайное число от 1 до 100.")
	fmt.Println("Сможете его угадать?")
	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Попыток:", 10-guesses)
		fmt.Print("Ваш вариант числа?: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Нужно было ввести число. Попробуйте снова.")
			guesses -= 1
			continue
		}

		if guess < target {
			fmt.Println("Ваше число меньше загаданного")
		} else if guess > target {
			fmt.Println("Ваше число больше загаданного")
		} else if guess == target {
			success = true
			fmt.Println("Вы угадали! Загаданное число:", target)
			break
		}

	}

	if !success {
		fmt.Println("Вы проиграли :( , загаданное число:", target)
	}
}
