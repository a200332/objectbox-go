package main

import (
	"flag"
	"log"
	"os"

	"github.com/objectbox/objectbox-go/test/performance/perf"
)

func main() {
	var dbName = flag.String("db", "db", "database directory")
	var count = flag.Int("count", 100000, "number of objects")
	var runs = flag.Int("runs", 30, "number of times the tests should be executed")
	flag.Parse()

	log.Printf("running the test %d times with %d objects", *runs, *count)

	// remove old database in case it already exists (and remove it after the test as well)
	os.RemoveAll(*dbName)
	defer os.RemoveAll(*dbName)

	executor := perf.CreateExecutor(*dbName)
	defer executor.Close()

	inserts := executor.PrepareData(*count)

	for i := 0; i < *runs; i++ {
		executor.PutAll(inserts)
		items := executor.ReadAll(*count)
		executor.UpdateAll(items)
		executor.RemoveAll()
		log.Printf("%d/%d finished", i+1, *runs)
	}

	executor.PrintTimes([]string{
		"PutAll",
		"ReadAll",
		"UpdateAll",
		"RemoveAll",
	})
}