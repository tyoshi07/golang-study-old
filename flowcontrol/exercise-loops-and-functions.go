package main

import (
		"fmt"
		"math"
)

/*
	ニュートン法による平方根の近似
	xの値を変えながら計算式を10回だけ繰り返す
*/
func Sqrt(x float64) float64 {
	z := 1.0
	for count := 0; count < 10; count++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	} 
	return z
}

/*
	ループ時、直前に求めたzの値の変化がない(少ない)場合にループから抜けるよう変更
*/
func Sqrt2(x float64) float64 {
	var z2 float64	// 最新の近似値を格納する変数
	diff := 1.0		// 近似値の差分
	count := 0		// forループの回数計測用変数
	for z := 1.0; math.Abs(diff) > 1.0e-15; diff = z2 - z {
		if count != 0 {
			z = z2		// 初回のループ以外、zの値をz2に更新
		}
		z2 = z - ((math.Pow(z, 2) - x) / (2 * z))	// ニュートン法による近似値計算
		count++
	}
	fmt.Println("ループ回数:", count)		// 計算回数の出力
	return z2		// 計算結果(ニュートン法による近似値を返す)
}

func main() {
	x := 8.0
	fmt.Println(math.Sqrt(x))	// mathパッケージのSqrt利用時
	fmt.Println(Sqrt(x))		// ニュートン法による計算利用時1
	fmt.Println(Sqrt2(x))		// ニュートン法による計算利用時2
}