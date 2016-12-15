package main
import (
	"github.com/BurntSushi/toml"
	"io"
)

type Configuration struct {
	Input InputConfiguration
	Filter FilterConfiguration
	Process ProcessConfiguration
	Output OutputConfiguration
}

type InputConfiguration struct {

}

type OutputConfiguration struct {

}

type FilterConfiguration struct {

}

type ProcessConfiguration struct {

}

func CheckConfiguration(option *Configuration, output io.Writer) {
	var _print = func (s string) {
		io.WriteString(output, s)
	}
	_print("Checking options ...\n")
	if nil == option {
		_print("Invalid configuration!!!\n")
		return
	}
	_print("\n")
	toml.NewEncoder(output).Encode(option)
	_print("\n")
	_print("Configration validated!\n")
}

func LoadConfiguration(filename string, option *Configuration) (e error) {
	_, e = toml.DecodeFile(filename, option)
	return e
}