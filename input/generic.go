package input

type InputConfig struct {
	Source string
	File *FileConfig
	Http *HttpConfig
}

type FileConfig struct {
	Follow bool
	Offset string
}

type HttpConfig struct {
	Method string
	Auth string
	Username string
	Password string
	Response ResponseStatus
}

type ResponseStatus struct {
	Error uint   // Default 400
	Failed uint  // Default 500
	Success uint // Default 201
}


type InputCollector interface {
	Initialize()
	Run()
	Destroy()
}

// Initialize
func init() {

}


func CreateInput(cfg *InputConfig) (InputCollector, error) {
	return nil, nil
}