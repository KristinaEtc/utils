package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
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
func GetFromGlobalConf(cf interface{}, defaultVal interface{}, whatParsed string) {

	file, e := ioutil.ReadFile(GetConfigFilename())
	if e != nil {
		log.WithCaller(slf.CallerShort).Errorf("Error: %s\n", e.Error())
	}

	if err := json.Unmarshal([]byte(file), cf); err != nil {
		log.WithCaller(slf.CallerShort).Errorf("Error parsing JSON : [%s]. For [%s] will be used defaulf options.",
			whatParsed, err.Error())
		cf = defaultVal
	} else {
		log.Infof("Parsed [%s] configuration from [%s] file", whatParsed, GetConfigFilename())
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

// GetPathToDir checked a path to directory - relative or absolute - and formatted it.
// If it's relative,the function add full path of executable file to it.
func GetPathToDir(logpath string) (string, error) {

	if filepath.IsAbs(logpath) == true {
		return logpath, nil
	} else {
		filename, err := osext.Executable()
		if err != nil {
			return "", err
		}

		fpath := filepath.Dir(filename)
		fpath = filepath.Join(fpath, logpath)
		return fpath, nil
	}
}

// Exists returns whether the given file or directory exists or not.
func Exists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
