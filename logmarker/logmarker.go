package logmarker

type InputWorker interface {

}

type OutputWorker interface {

}

type ProcessWorker interface {

}

type FilterWorker interface {

}

type LogMarker struct {
	Inputs map[string]InputWorker
	Filters map[string]FilterWorker
	Processors map[string]ProcessWorker
	Outputs map[string]OutputWorker
}


func Run(option *Configuration) error {

	return nil
}
