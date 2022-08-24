package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var mutex sync.Mutex
var UrlMappingHttpClient = sync.Map{}
var MapClient = make(map[int]*http.Client)

func main() {
	c := make(chan string, 2)  // buffer channel
	c1 := make(chan string, 2) // buffer channel
	c2 := make(chan string, 2) // buffer channel
	c3 := make(chan string, 2)

	fmt.Println("START", UrlMappingHttpClient)
	go AddTimeoutGoRoutine(c)
	go AddTimeoutGoRoutine2(c1)
	// go ResetTimeoutGoroutine(c3)
	// go ResetTimeoutWithMutexGoroutine(c3)
	go ResetMethodSyncMap(c3)
	go AddTimeoutGoRoutine3(c2)

	fmt.Println(<-c)
	fmt.Println(<-c3)
	fmt.Println(<-c1)
	fmt.Println(<-c2)

	fmt.Println("END", UrlMappingHttpClient)
	// value, ok := UrlMappingHttpClient.Load(1)
	// if ok {
	// 	fmt.Println(reflect.TypeOf(value.(*http.Client)))
	// 	fmt.Println(value, ok)
	// }
}

func InsertTimeout(timeoutData []int) {
	for _, val := range timeoutData {
		httpClient := &http.Client{
			Timeout: time.Second * time.Duration(val),
		}
		UrlMappingHttpClient.Store(val, httpClient)
		// MapClient[val] = httpClient
		// fmt.Println(val)
	}

}

func AddTimeoutGoRoutine(ch chan string) {
	timeoutData := []int{1, 2, 3, 4, 5, 6}
	InsertTimeout(timeoutData)
	ch <- fmt.Sprintf("Data 1")

}

func AddTimeoutGoRoutine2(ch chan string) {
	timeoutData := []int{7, 8, 9, 10}
	InsertTimeout(timeoutData)
	ch <- fmt.Sprintf("Data 2")
}

func AddTimeoutGoRoutine3(ch chan string) {
	timeoutData := []int{11, 12, 13, 14}
	InsertTimeout(timeoutData)
	ch <- fmt.Sprintf("Data 3")
}

func ResetMethodSyncMap(ch chan string) {
	eraseSyncMap(UrlMappingHttpClient)
	ch <- fmt.Sprint("Reset Timeout")

}
func ResetTimeoutGoroutine(ch chan string) {

	UrlMappingHttpClient = sync.Map{}
	ch <- fmt.Sprint("Reset Timeout")
}

func ResetTimeoutWithMutexGoroutine(ch chan string) {
	mutex.Lock()
	UrlMappingHttpClient = sync.Map{} // reset checkpoint
	mutex.Unlock()
	ch <- fmt.Sprint("Reset Timeout")
}

func eraseSyncMap(m sync.Map) {
	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Println("Deleted", key)
		m.Delete(key)
		return true
	})
}
