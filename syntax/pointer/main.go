package main

import (
	"fmt"
)

func main() {
	n := 100
	// %vでその型のデフォルトのフォーマットで表示する
	fmt.Printf("main: Address of n is %v\n", &n)

	// 値渡し
	// コピーされるので、元のnに変化はない
	returnValue := increment(n)
	fmt.Printf("Value of n is %d\n", n)
	fmt.Printf("Return Value of increment is %d\n", returnValue)
	// 参照渡し
	// 戻り値を受け取らなくても渡した変数が書き換わる
	incrementWithPointer(&n)
	fmt.Printf("Value of n is %d\n", n)
}

func increment(n int) int {
	fmt.Printf("increment: Address of n is %v\n", &n)
	return n + 1
}
func incrementWithPointer(n *int) {
	fmt.Printf("incrementWithPointer: Address of n is %v\n", n)
	*n++
}
