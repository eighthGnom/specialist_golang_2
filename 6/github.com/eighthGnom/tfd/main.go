package main

import (
	"log"
	"net/http"
	"strconv"
)

var dataMap map[int]int

func init() {
	dataMap = make(map[int]int)
}

func main() {
	http.HandleFunc("/factorial", HTTPFactorial)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func HTTPFactorial(writer http.ResponseWriter, request *http.Request) {
	num := request.FormValue("num")
	intNum, err := strconv.Atoi(num)
	if err != nil {
		http.Error(writer, "cant convert", 400)
		return
	}
	intNum = factorial(intNum)
	writer.WriteHeader(200)
	writer.Write([]byte(strconv.Itoa(intNum)))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	if res, ok := dataMap[num]; ok {
		return res
	}
	fac := 1
	for i := 1; i <= num; i++ {
		fac *= i
	}
	dataMap[num] = fac
	return fac
}
