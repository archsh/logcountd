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