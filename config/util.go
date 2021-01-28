package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func getConfigDir() string {
	dirPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dirPath, "tmeow")
}

func configFileExists(dir string) bool {
	list := []string{
		dir,
		filepath.Join(dir, crdFile),
		filepath.Join(dir, optFile),
		filepath.Join(dir, colFile),
	}

	for _, path := range list {
		if _, err := os.Stat(path); err != nil {
			return false
		}
	}

	return true
}

func saveYaml(dir, filename string, in interface{}) {
	// 変換
	buf, err := yaml.Marshal(in)
	if err != nil {
		log.Fatal(err)
	}
	// 保存
	path := filepath.Join(dir, filename)
	err = ioutil.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func loadYaml(dir, filename string, out interface{}) {
	// 読込
	path := filepath.Join(dir, filename)
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// 構造体にマッピング
	err = yaml.Unmarshal(buf, out)
	if err != nil {
		log.Fatal(err)
	}
}
