package config

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	path    string
	configs map[string]string
	mu      sync.RWMutex
)

func init() {
	// path = getConfigFilePath()
	// err := updateConfigFileData()
	// if err != nil {
	// 	log.Fatal("Error initializing config file.")
	// }
	// watchConfigFile()
}

func updateConfigFileData() error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var data map[string]string
	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}
	mu.Lock()
	configs = data
	mu.Unlock()
	return nil
}

func getConfigFilePath() string {
	// constants
	key := "CONFIG_FILE"
	args := os.Args
	for _, v := range args {
		if strings.Index(v, key) == 0 && strings.Index(v, "=") != -1 {
			// this is the one
			vArr := strings.Split(v, "=")
			return vArr[1]
		}
	}
	log.Fatal("Config file path missing.")
	return ""
}

func watchConfigFile() {
	watchInitWg := sync.WaitGroup{}
	watchInitWg.Add(1)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			watchInitWg.Done()
			return
		}
		defer watcher.Close()
		eventWg := sync.WaitGroup{}
		eventWg.Add(1)
		go func() {
			defer eventWg.Done()
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					writeCreateMask := fsnotify.Write | fsnotify.Create
					if filepath.Clean(event.Name) == path &&
						event.Op&writeCreateMask != 0 {
						// file written or created newly
						go updateConfigFileData()
					} else if filepath.Clean(event.Name) == path &&
						event.Op&fsnotify.Remove != 0 {
						// file removed
						return
					}
				case <-watcher.Errors:
					return
				}
			}
		}()
		watcher.Add(filepath.Dir(path))
		watchInitWg.Done()
		eventWg.Wait()
	}()
	watchInitWg.Wait()
}

// GET is used to get the value for a config
func GET(key string) string {
	// mu.RLock()
	// defer mu.RUnlock()
	// return configs[key]
	return os.Getenv(key)
}
