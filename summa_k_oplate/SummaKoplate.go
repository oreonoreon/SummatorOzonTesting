package summa_k_oplate

import (
	"bufio"
	"fmt"
	"os"
)

//В магазине акция: «купи три одинаковых товара и заплати только за два». Конечно, каждый купленный товар может участвовать лишь в одной акции. Акцию можно использовать многократно.
//
//Например, если будут куплены 7 товаров одного вида по цене 2 за штуку и 5 товаров другого вида по цене 3 за штуку, то вместо 7⋅2+5⋅3 надо будет оплатить 5⋅2+4⋅3=22.
//
//Считая, что одинаковые цены имеют только одинаковые товары, найдите сумму к оплате.
//
//Неполные решения этой задачи (например, недостаточно эффективные) могут быть оценены частичным баллом.
//
//Входные данные
//В первой строке записано целое число t (1≤t≤104) — количество наборов входных данных.
//
//Далее записаны наборы входных данных. Каждый начинается строкой, которая содержит n (1≤n≤2⋅105) — количество купленных товаров. Следующая строка содержит их цены p1,p2,…,pn (1≤pi≤104). Если цены двух товаров одинаковые, то надо считать, что это один и тот товар.
//
//Гарантируется, что сумма значений n по всем тестам не превосходит 2⋅105.
//
//Выходные данные
//Выведите t целых чисел — суммы к оплате для каждого из наборов входных данных.

func Start() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)
	var numberofgoods int

	for i := 0; i < testCount; i++ {
		mapa := make(map[int]int)
		fmt.Fscan(in, &numberofgoods)
		for j := 0; j < numberofgoods; j++ {
			var a int
			fmt.Fscan(in, &a)
			if v, ok := mapa[a]; !ok {
				mapa[a] = 1
			} else {
				v++
				mapa[a] = v
			}

		}
		vgixx := magic(mapa)

		fmt.Fprintln(out, vgixx)
	}
}
func magic(mapa map[int]int) int {
	var sum int
	for k, v := range mapa {
		sum += k*(v%3) + 2*k*(v/3)
	}
	return sum
}

//6
//12
//2 2 2 2 2 2 2 3 3 3 3 3
//12
//2 3 2 3 2 2 3 2 3 2 2 3
//1
//10000
//9
//1 2 3 1 2 3 1 2 3
//6
//10000 10000 10000 10000 10000 10000
//6
//300 100 200 300 200 300