package processor

type Processor interface {
	init(processorContext *ProcessorContext)
	process(record *Record)
	close()
}
