// specify this one once for each key / value / implementation triple
//go:generate genny -pkg=structures -in=../src/templates/priority_queue.go -out=./generated/int_pq.go     gen "GenericItem=int"
//go:generate genny -pkg=structures -in=../src/templates/priority_queue.go -out=./generated/custom_pq.go  gen "GenericItem=examples.MyCustomType"

package examples

type MyCustomType struct{}
