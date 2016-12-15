package logmarker
import (
	"io"
	"logmarked/input"
	"logmarked/filter"
	"logmarked/output"
	"logmarked/process"
	"gopkg.in/yaml.v2"
	"os"
	//log "github.com/Sirupsen/logrus"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	LogMarked *GlobalConfig
	Inputs []*input.InputConfig
	Filters []*filter.FilterConfig
	Processes []*process.ProcessConfig
	Outputs []*output.OutputConfig
}

type GlobalConfig struct {
	Log_File string
	Log_Level string
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
	if b, e := yaml.Marshal(option); nil == e {
		output.Write(b)
	}else{
		output.Write([]byte(fmt.Sprintf("Marshal conifiguration failed:> %s", e)))
	}
	_print("\n")
	_print("Configration validated!\n")
}

func LoadConfiguration(filename string, option *Configuration) (e error) {
	if f, e := os.Open(filename); nil != e {
		return e
	}else{
		defer f.Close()
		if buf, e := ioutil.ReadAll(f); nil != e {
			return e
		}else{
			//log.Debugf("Read %d bytes from file '%s'.", n, filename)
			//fmt.Printf("Read %d bytes from file '%s'.\n", len(buf), filename)
			e = yaml.Unmarshal(buf, &option)
		}
	}
	return e
}