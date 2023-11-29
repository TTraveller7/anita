package main

import (
	"context"
	"log"
	"os"
	"ttraveller7/anita"
	"ttraveller7/anita/config"
	"ttraveller7/anita/processor"
	"ttraveller7/anita/state"
)

var logs *log.Logger = log.New(os.Stdout, "", 0)

type WordCountProcessor struct {
	processor.Processor
}

func (p *WordCountProcessor) init(processorContext *processor.ProcessorContext) {

}

func (p *WordCountProcessor) process(record *processor.Record) {

}

func (p *WordCountProcessor) close() {}

func main() {
	conf := &config.AnitaConfig{}
	builder := anita.NewTopology(conf)
	if err := builder.AddStateStore(state.NewRedisKeyValueStateStoreBuilder("counts", false, "")); err != nil {
		logs.Printf("add state store failed: %v", err)
		return
	}
	builder.AddProcessor(processor.NewProcessorNodeBuilder("processorOne", WordCountProcessor{}))
	kstream := anita.NewKafkaStreams(builder)
	if err := kstream.Start(context.Background()); err != nil {
		logs.Printf("start kstream failed: %v", err)
		return
	}
}
