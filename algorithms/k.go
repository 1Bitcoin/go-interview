package main

import "sort"

func k(arr []int, k int) []int {

	// Создаём карту для подсчёта количества вхождений каждого элемента

	counts := make(map[int]int)

	// Подсчитываем количество вхождений каждого элемента

	for _, num := range arr {

		counts[num]++

	}

	// Создаём срез для хранения уникальных значений

	unique := make([]int, len(counts))

	// Добавляем уникальные значения в срез

	for num := range counts {

		unique = append(unique, num)

	}

	// Сортируем уникальные значения в порядке убывания исходя из количества вхождений

	sort.Slice(unique, func(i, j int) bool {

		return counts[unique[i]] > counts[unique[j]]

	})

	return unique[:k]

}
