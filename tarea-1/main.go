package main

import "fmt"

const (
	MAX_LIMIT      int = 10 // Constante para el número máximo de puntuaciones
	MIN_SCORE      int = 1  // Constante para la puntuación mínima
	MAX_SCORE      int = 5  // Constante para la puntuación máxima
	MIN_GOOD_SCORE int = 4  // Constante para la puntuación mínima de buena puntuación
	MIN_BAD_SCORE  int = 1  // Constante para la puntuación mínima de mala puntuación
	MAX_GOOD_SCORE int = 5  // Constante para la puntuación máxima de buena puntuación
	MAX_BAD_SCORE  int = 2  // Constante para la puntuación máxima de mala puntuación
)

func main() {
	var score int
	var numberBadScores int
	var numberGoodScores int
	var scores = map[int]int{}

	for i := 0; i < MAX_LIMIT; i++ {
		fmt.Println("Ingrese un puntaje del 1 al 5: ")
		fmt.Scan(&score)

		if score < MIN_SCORE || score > MAX_SCORE {
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

		if score == MIN_GOOD_SCORE || score == MAX_GOOD_SCORE {
			*numberGoodScores = 1
		} else if score == MIN_BAD_SCORE || score == MAX_BAD_SCORE {
			*numberBadScores = 1
		}
	} else {
		scores[score]++

		if score == MIN_GOOD_SCORE || score == MAX_GOOD_SCORE {
			*numberGoodScores++
		} else if score == MIN_BAD_SCORE || score == MAX_BAD_SCORE {
			*numberBadScores++
		}
	}
}
