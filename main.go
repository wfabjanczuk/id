package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/wfabjanczuk/id/unique"
)

func main() {
	alphanumericCharList := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	generator, err := unique.NewGenerator(math.MaxInt, 128, alphanumericCharList)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	ids, err := generator.ToChannel(context.Background())
	if err != nil {
		log.Fatal(err, ids)
	}

	total, mod10k := 0, 0
	for id := range ids {
		total++
		if mod10k++; mod10k == 10000 {
			fmt.Println(string(id), "| total:", total, "| duration:", time.Since(now))
			mod10k = 0
		}
	}
}
