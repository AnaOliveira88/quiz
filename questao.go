package main

import (
	"fmt"
)

func questao(Questao q) {
	var resp int
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(q[0])
		fmt.Println("1", q[1])
		fmt.Println("2", q[2])
		fmt.Println("3", q[3])
		fmt.Scan("%d", &resp)
		if resp == q[4] {
			fmt.Println("correcto")
		} else {
			fmt.Println("errado")
		}
	}

}
