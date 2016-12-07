package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/britojr/tcc/codec"
	"github.com/britojr/tcc/dandelion"
)

func main() {
	log.Println("Welcome!")
	log.Println("This program expects input in the following format:")
	log.Println("")
	log.Println("===")
	log.Println("k")
	log.Println("Q_1")
	log.Println("...")
	log.Println("Q_k")
	log.Println("s")
	log.Println("p_1 l_1")
	log.Println("...")
	log.Println("p_s l_s")
	log.Println("===")
	log.Println("")
	log.Println("Where:")
	log.Println("- Q_i correspond to the values in Q.")
	log.Println("- (p_i, l_i) correspond to the pairs in the code S.")
	log.Println("")

	var C codec.Code

	var k int
	fmt.Scanf("%d", &k)
	C.Q = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d", &C.Q[i])
	}
	sort.Ints(C.Q)

	var s int
	fmt.Scanf("%d", &s)
	C.S = &dandelion.DandelionCode{make([]int, s), make([]int, s)}
	for i := 0; i < s; i++ {
		fmt.Scanf("%d %d", &C.S.P[i], &C.S.L[i])
	}

	Tk, err := codec.DecodingAlgorithm(&C)
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Printf("%d\n", len(Tk.Adj))
	fmt.Printf("%d\n", Tk.K)
	m := 0
	for u := 0; u < len(Tk.Adj); u++ {
		for i := 0; i < len(Tk.Adj[u]); i++ {
			v := Tk.Adj[u][i]
			if v > u {
				m++
			}
		}
	}
	fmt.Printf("%d\n", m)
	for u := 0; u < len(Tk.Adj); u++ {
		for i := 0; i < len(Tk.Adj[u]); i++ {
			v := Tk.Adj[u][i]
			if v > u {
				fmt.Printf("%d %d\n", u, v)
			}
		}
	}
}
