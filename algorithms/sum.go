package main

func sumSlices(a []int, b []int) []int {

	// Получаем длины срезов

	lenA := len(a)

	lenB := len(b)

	// Выбираем большую длину

	maxLen := lenA

	if lenB > maxLen {

		maxLen = lenB

	}

	// Создаем срез с максимальной длиной

	result := make([]int, maxLen+1)

	// Выполняем сложение цифр соответствующих индексов

	carry := 0 // перенос

	for i := 0; i < maxLen; i++ {

		sum := carry

		if lenA > i {

			sum += a[lenA-1-i]

		}

		if lenB > i {

			sum += b[lenB-1-i]

		}

		result[maxLen-i] = sum % 10

		carry = sum / 10

	}

	// Если остался перенос, добавляем его в начало среза

	if carry > 0 {

		result[0] = carry

	}

	if result[0] == 0 {
		return result[1:]
	}

	return result

}
