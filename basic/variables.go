package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	// "io"
	"net/http"
	// "strings"
)

// ======= advance variables =======
type sampleOrder struct {
	TotalCost     float64 `json:"total_cost"`
	SampleOrderId int64   `json:"sample_order_id"`
	ProductsCount int32   `json:"products_count"`
}

func main() {

	var name string = "Yogie"
	var age int = 24
	var isEmployee bool = true
	var height float64 = 170.5
	var b byte = 11

	fmt.Println("Your Name  :", name)
	fmt.Println("Your Age   :", age)
	fmt.Println("Employee   :", isEmployee)
	fmt.Println("Your Heigh :", height)
	fmt.Println("Byte Value :", b)

	// ===========
	// var url string = "http://other-team-services.consul"
	// resp, err := http.Get(url)

	// demo only
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(
			`{"total_cost":15.4, "sample_order_id":17442, "products_count":5}`,
		)),
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	// parse JSON body to get Go object here
	sampleOrder := sampleOrder{}
	err := json.Unmarshal(body, &sampleOrder)
	if err == nil {
		fmt.Println(sampleOrder.TotalCost)
		fmt.Println(sampleOrder.SampleOrderId)
		fmt.Println(sampleOrder.ProductsCount)
	}
}

// goplay.space/#A8Ad5DoutAW
