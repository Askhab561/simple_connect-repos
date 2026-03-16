package main

import (
	"fmt"
	"study/feature1"
	"study/feature2"
	simpleconnection "study/feature_postgres/simple_connection"
)

func main() {
	feature1.Feature1()
	feature2.Feature2()
	fmt.Println(" Hello i'm package main")
	simpleconnection.CheckConnection()
}
