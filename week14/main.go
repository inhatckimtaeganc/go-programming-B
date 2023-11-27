package main

import "fmt"

func main() {
	//games := map[int]string{456: "성기훈", 218: "박해수", 067: "강새벽", 001: "오일남", 199: "알리", 101: "아이오아이"}
	games := map[int]string{
		456: "성기훈",
		218: "박해수",
		067: "강새벽",
		001: "오일남",
		199: "알리",
		101: "아이오아이",
	}

	//fmt.Println(games[067])
	for _, v := range games {
		fmt.Println(v)
	}
	// update
	games[101] = "장덕수"
	// delete
	delete(games, 199)
	for k, v := range games {
		fmt.Println(k, v)
	}
}
