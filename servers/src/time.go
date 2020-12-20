package main

import (
	"fmt"
	"time"
)

func main() {

	server_local := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	fmt.Printf("server_local = %v\n", server_local)

}
