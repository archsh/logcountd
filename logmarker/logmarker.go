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
	Filters []FilterWorker
	Processors []ProcessWorker
	Outputs map[string]OutputWorker
}


func Run(option *Configuration) error {

	return nil
}
