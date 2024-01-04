package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	buckets := make(map[int][]float64)

	for _, temperature := range temperatures {
		bucket := int(temperature/10) * 10                     // определяем к какой группе будет принадлежать число
		buckets[bucket] = append(buckets[bucket], temperature) // добавляем число в группу
	}

	for k, v := range buckets {
		fmt.Println(k, ":", v)
	}
}
