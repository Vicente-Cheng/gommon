package harvester

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Generate Temp file with YAML content, return the file name
func GenerateYAMLTempFile(obj interface{}, prefix string) (string, error) {
	tempFile, err := os.CreateTemp("/tmp", prefix)
	if err != nil {
		logrus.Errorf("Create temp file failed. err: %v", err)
		return "", err
	}
	defer tempFile.Close()

	bytes, err := yaml.Marshal(obj)
	if err != nil {
		logrus.Errorf("Generate YAML content failed. err: %v", err)
		return "", err
	}
	if _, err := tempFile.Write(bytes); err != nil {
		logrus.Errorf("Write YAML content to file failed. err: %v", err)
		return "", err
	}

	logrus.Debugf("Content of %s: %s", tempFile.Name(), string(bytes))

	return tempFile.Name(), nil
}
