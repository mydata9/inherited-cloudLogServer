package modUtility

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type CGatlingConfig struct {
	kValue  map[string]string
	appPath string
}

var g_singleGatlingConfig *CGatlingConfig = &CGatlingConfig{kValue: map[string]string{}}

func GetSingleGatlingConfig() *CGatlingConfig {
	return g_singleGatlingConfig
}

func (pInst *CGatlingConfig) Initialize(appName string) error {
	if appName == "" {
		return errors.New("appname is empty")
	}
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	pInst.appPath = path[:index+1]

	pInst.listEnv()
	pInst.loadAppConfig(appName)

	return nil
}

func (pInst *CGatlingConfig) listEnv() int {
	var iCount = 0
	for i, env := range os.Environ() {
		// env is
		envPair := strings.SplitN(env, "=", 2)
		key := envPair[0]
		value := envPair[1]
		if key != "" {
			pInst.kValue[key] = value
			iCount = i
		}

	}

	return iCount
}

func (pInst *CGatlingConfig) loadAppConfig(appName string) int {
	var iCount = 0
	cfgPath := path.Join(pInst.appPath, appName+".cfg")
	f, err := os.Open(cfgPath)
	if err != nil {
		return -1
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		iRet := pInst.analyseConfig(scanner.Text())
		if iRet > 0 {
			iCount++
		}
	}

	return iCount
}

func (pInst *CGatlingConfig) analyseConfig(line string) int {
	strPair := strings.SplitN(line, "=", 2)
	if len(strPair) < 2 {
		return -1
	}
	if strPair[0] != "" {
		pInst.kValue[strPair[0]] = strPair[1]
	}
	return 1
}

func (pInst *CGatlingConfig) Get(key string) string {
	return pInst.kValue[key]
}

func (pInst *CGatlingConfig) Set(key string, value string) {
	pInst.kValue[key] = value
}
