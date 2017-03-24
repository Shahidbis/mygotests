package main


import "runtime"

func main() {
	runtime.GOMAXPROCS(8)


	ch := make(chan string)
	donech := make(chan bool)
	go abcGen(ch)
	go printer(ch, donech)

	println("this comes first")

	<- donech

}go	

func abcGen(ch chan  string){
	for l := byte('a'); l <= byte('z'); l++{
		ch <- string(l)
	}
	close(ch)
}

func printer(ch chan string, donech chan bool){
	for l := range ch {
		println(l)
	}
	donech <- true
}

