// average2 вычисляет среднее значение

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64 = 0                    // Объявление переменной для хранения суммы чисел
	for _, arargument := range arguments { // обрабатываем каждый элемент командной строки
		number, err := strconv.ParseFloat(arargument, 64) // строка преобразуется в float64
		if err != nil {                                   // если при преобразовании строки произошла ошибка, программа выводит сообщение и завершается
			log.Fatal(err)
		}
		sum += number // число прибавляется к сумме
	}
	sampleCount := float64(len(arguments))          // длина сегмента argemnts используется как количество значений данных
	fmt.Printf("Average: %0.2f\n", sum/sampleCount) // вычесление среднего значения и его вывод
}
