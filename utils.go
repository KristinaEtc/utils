package utils

import (
	"encoding/json"
	"io/ioutil"
	"runtime"

	"github.com/kardianos/osext"
	"github.com/ventu-io/slf"
)

var log slf.StructuredLogger

const pwdCurr string = "github.com/KristinaEtc/utils"

func init() {
	log = slf.WithContext(pwdCurr)
}

// GetGlobalConf unmarshal json-object cf
// If parsing was not successuful, function return a structure with default options
func GetFromGlobalConf(cf interface{}, defaultVal interface{}) {

	file, e := ioutil.ReadFile(GetConfigFilename())
	if e != nil {
		log.WithCaller(slf.CallerShort).Errorf("Error: %s\n", e.Error())
	}

	if err := json.Unmarshal([]byte(file), cf); err != nil {
		log.WithCaller(slf.CallerShort).Errorf("Error parsing JSON: %s. Will be used defaulf options.", err.Error())
		cf = defaultVal
	} else {
		log.Infof("Global options will be used from [%s] file", GetConfigFilename())
	}
	//log.Debugf("%v", cf)
}

// GetConfigFilename is a function fot getting a name of a binary with full path to it
func GetConfigFilename() string {
	binaryPath, err := osext.Executable()
	if err != nil {
		log.WithCaller(slf.CallerShort).Errorf("Error: could not get a path to binary file: %s\n", err.Error())
	}
	if runtime.GOOS == "windows" {
		// without ".exe"
		binaryPath = binaryPath[:len(binaryPath)-4]
		log.WithCaller(slf.CallerShort).WithField("binaryPath", binaryPath).Debug("Configfile for windows")
	}

	return binaryPath + ".config"
}
