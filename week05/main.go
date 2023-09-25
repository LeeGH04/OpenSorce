package main

import (
	"fmt"
	"time"
)

func main(){
	var now time.Time = time.Now()
	year := now.Year()
	// var year int = now.Year()
	var month string = now.month().string()
	fmt.Println(year,month,now.Hour(),now.Minute(), now.Second())
}

*
func main() {
	HotSpurs := "hm ? j madi?"
	replacePlayer := strings.NewReplacer("?", "son")
	Player := replacePlayer.Replace(HotSpurs)
	fmt.Println(player)
}


