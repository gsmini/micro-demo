package main

import (
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/store"
)

func main() {
	srv := service.New(service.Name("storage-operate"))
	srv.Init()

	records, err := store.Read("name", func(r *store.ReadOptions) {
		//读取micro表中的title key
		r.Table = "micro"
	})
	if err != nil {
		fmt.Println("Error reading from store: ", err)
	}

	if len(records) == 0 {
		fmt.Println("No records")
	}
	for _, record := range records {
		fmt.Printf("key: %v, value: %v\n", record.Key, string(record.Value))
	}

	time.Sleep(1 * time.Hour)
}
