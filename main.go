// average2 вычисляет среднее значение

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	var numbers []float64                  // В этом сегменте будут храниться числа, для которых вычисляется среднее
	for _, arargument := range arguments { // обрабатываем каждый элемент командной строки
		number, err := strconv.ParseFloat(arargument, 64) // строка преобразуется в float64
		if err != nil {                                   // если при преобразовании строки произошла ошибка, программа выводит сообщение и завершается
			log.Fatal(err)
		}
		numbers = append(numbers, number)

	}
	fmt.Printf("Average: %0.2f\n", average(numbers...)) // вычесление среднего значения и его вывод
}
