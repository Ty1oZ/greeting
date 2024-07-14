// count подсчитывает количество вхождений
// каждой строки в файле.

package main

import (
	"fmt"
	"log"

	"github.com/ty1oz/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int) // объявляем карту, у которой ключами являются имена кандидатов, а значениями - счетчики голосов
	for _, line := range lines {
		counts[line]++ // увеличивает счетчик голосов для текущего кандидата
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count) // вывод ключа(имя кандидата) и значения (количество голосов)
	}
}
