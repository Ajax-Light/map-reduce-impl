package main

import "os"
import "fmt"
import "main/master"
import "container/list"
import "strings"
import "strconv"
import "unicode"

// our simplified version of MapReduce does not supply a
// key to the Map function, as in the paper; only a value,
// which is a part of the input file contents
func Map(value string) *list.List {
	notLetter := func(r rune) bool {
		return !unicode.IsLetter(r)
	}
	fields := strings.FieldsFunc(value, notLetter)

	l := list.New()
	for _, f := range fields {
		kv := master.KeyValue{Key: f, Value: "1"}
		l.PushBack(kv)
	}
	return l
}

// iterate over list and add values
func Reduce(key string, values *list.List) string {
	count := 0
	// iterate over list
	for e := values.Front(); e != nil; e = e.Next() {
		v := e.Value.(string)
		intValue, _ := strconv.Atoi(v)
		count += intValue
	}
	return strconv.Itoa(count)
}

// Can be run in 3 ways:
// 1) Sequential (e.g., go run wc.go sequential x.txt)
// 2) Master (e.g., go run wc.go master localhost:7777)
// 3) Worker (e.g., go run wc.go worker localhost:7777 localhost:7778)
// 4) Submit (e.g., go run wc.go submit x.txt localhost:7777)
// 5) Terminate (e.g., go run wc.go terminate localhost:7777)
func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s: see usage comments in file\n", os.Args[0])
		return
	}

	switch os.Args[1] {
	case "sequential":
		job := master.Job{NMap: 5, NReduce: 3,
			InputPath: os.Args[2]}
		master.RunSequential(job, Map, Reduce)
	case "master":
		// launch a long-running master process
		master.RunMasterProcess(os.Args[2])
	case "worker":
		// launch a long-running worker process
		master.RunWorkerProcess(os.Args[2], os.Args[3],
			Map, Reduce, -1)
	case "submit":
		job := master.Job{NMap: 2, NReduce: 1,
			InputPath: os.Args[2]}
		// submit the job to master
		master.SubmitJob(job, os.Args[3])
	}
}
