package main

import (
	"log"
	"os"
	"ttraveller7/anita"
	"ttraveller7/anita/processor"
	"ttraveller7/anita/state"
)

var logs *log.Logger = log.New(os.Stdout, "", 0)

func main() {
	builder := anita.NewTopology()
	if err := builder.AddStateStore(state.NewRedisKeyValueStateStoreBuilder("counts", false, "")); err != nil {
		logs.Printf("add state store failed: %v", err)
		return
	}
	builder.AddProcessor(processor.NewProcessorNodeBuilder("processorOne", nil))
}
