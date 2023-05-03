package harvester

import (
	"io/fs"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	defaultFilePerm = 0644
)

// Generate Temp file with YAML content, return the file name
func GenerateYAMLTempFile(obj interface{}, prefix string) (string, error) {
	return GenerateYAMLTempFileWithPerm(obj, prefix, defaultFilePerm)
}

// Generate Temp file with YAML content and permission, return the file name
func GenerateYAMLTempFileWithPerm(obj interface{}, prefix string, perm fs.FileMode) (string, error) {
	tempFile, err := os.CreateTemp("/tmp", prefix)
	if err != nil {
		logrus.Errorf("Create temp file failed. err: %v", err)
		return "", err
	}
	defer tempFile.Close()

	if err = tempFile.Chmod(perm); err != nil {
		logrus.Errorf("Chmod temp file failed. err: %v", err)
		return "", err
	}

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

// Generate Temp file with buffer, return the file name
func GenerateTempFile(buf []byte, prefix string) (string, error) {
	return GenerateYAMLTempFileWithPerm(buf, prefix, defaultFilePerm)
}

// Generate Temp file with buffer and permission, return the file name
func GenerateTempFileWithPerm(buf []byte, prefix string, perm fs.FileMode) (string, error) {
	tempFile, err := os.CreateTemp("/tmp", prefix)
	if err != nil {
		logrus.Errorf("Create temp file failed. err: %v", err)
		return "", err
	}
	defer tempFile.Close()

	if err = tempFile.Chmod(perm); err != nil {
		logrus.Errorf("Chmod temp file failed. err: %v", err)
		return "", err
	}

	if _, err := tempFile.Write(buf); err != nil {
		logrus.Errorf("Write YAML content to file failed. err: %v", err)
		return "", err
	}

	logrus.Debugf("Content of %s: %s", tempFile.Name(), string(buf))

	return tempFile.Name(), nil
}
