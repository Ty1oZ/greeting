// count подсчитывает количество вхождений
// каждой строки в файле.

package main

import (
	"fmt"
	"log"

	"github.com/ty1oz/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt") // Читает файл «votes.txt» и возвращает сегмент содержащий все строки из файла
	if err != nil {
		log.Fatal(err)
	}

	var names []string // переменная для хранения сегмента с именами кандидатов
	var counts []int   // сегмент с количеством вхождения каждого имени
	for _, line := range lines {
		matched := false
		for i, name := range names { // перебор всех значений из сегмента names
			if name == line { // если эта строка совпадает с текщим именем
				counts[i]++    // увеличивем соответсвующий счетчик
				matched = true // устанавливает признак обнаружения совпадения
			}
		}
		if matched == false { // если совпадений не найдено
			names = append(names, line) // добавить его как новое имя
			counts = append(counts, 1)  // добавить новый счетчик(текущая строка станет первым вхождением)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i]) // вывести каждый элемент из сегмента names и соответсующий элемент из сегмента counts
	}
}
