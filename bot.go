package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"./treasures"
)

// No error handling.
// No pretty bells and whistles.
// Only code doin' shit.

func main() {
	client := treasures.NewClient()

	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "usage: %s [ID] [AuthKey] [Coins]\n", os.Args[0])

		os.Exit(-1)
	}

	id := os.Args[1]
	auth := os.Args[2]
	coins, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Введите числовое значение желаемого количества монет в игре\n")

		os.Exit(-1)
	}

	err = client.Authorize(id, auth)
	if err != nil {
		panic(err)
	}

	if client.Episode == 1 && client.Level == 128 {
		fmt.Fprintf(os.Stderr, "Аккаунт заблокирован\n")

		os.Exit(-1)
	}

	start_level := (client.Episode-1)*20 + client.Level - 1
	fmt.Println("Уровень персонажа:", (client.Episode-1)*20+client.Level-1)

	for l := start_level; l <= 2119; l++ {
		level := l%20 + 1
		episode := l/20 + 1
		userLevel := l
		score := 0xDEAD*10 + rand.Int()%0xFAC

		_ = client.FinishLevel(episode, level, userLevel, score)
		fmt.Println("Прошёл уровень", level, "эпизода", episode)

		if level == 20 && episode != 1 && l < 2119 {
			_ = client.BuyKeys(episode, coins+rand.Int()%1337)

			fmt.Println("Купил ключи для перехода на следующий эпизод!")
		}
	}

	fmt.Println("Готово! Вы достигли максимального на данный момент уровня.")
}
