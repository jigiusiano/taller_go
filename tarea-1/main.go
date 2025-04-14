package main

import "fmt"

func main() {
	var score int
	var numberBadScores int
	var numberGoodScores int
	const MAX_LIMIT int = 10
	var scores = map[int]int{}

	for i := 0; i < MAX_LIMIT; i++ {
		fmt.Println("Ingrese un puntaje del 1 al 5: ")
		fmt.Scan(&score)

		if score < 1 || score > 5 {
			fmt.Println("Puntaje inválido. Debe estar entre 1 y 5.")
			i--
		}

		countAndAssignScore(score, scores, &numberGoodScores, &numberBadScores)
	}

	for key, value := range scores {
		fmt.Printf("El puntaje %d se votó %d veces\n", key, value)
	}

	if numberGoodScores > numberBadScores {
		fmt.Println("Buen resultado!")
	} else {
		fmt.Println("Resultado mejorable")
	}
}

func countAndAssignScore(score int, scores map[int]int, numberGoodScores *int, numberBadScores *int) {
	_, existsKey := scores[score]

	if !existsKey {
		scores[score] = 1

		if score == 4 || score == 5 {
			*numberGoodScores = 1
		} else if score == 1 || score == 2 {
			*numberBadScores = 1
		}
	} else {
		scores[score]++

		if score == 4 || score == 5 {
			*numberGoodScores++
		} else if score == 1 || score == 2 {
			*numberBadScores++
		}
	}
}
