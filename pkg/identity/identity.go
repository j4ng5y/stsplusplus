package identity

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	configFileContents = `[%s]
region = %s
output = %s`
	credentialsFileContents = `[%s]
aws_access_key_id = %s
aws_secret_access_key = %s`
)

type Identity struct {
	Name               string
	Region             Region
	Output             Output
	AWSAccessKeyID     AWSAccessKeyID
	AWSSecretAccessKey AWSSecretAccessKey
}

func New(name string, region Region, output Output, accessKeyID AWSAccessKeyID, secretAccessKey AWSSecretAccessKey) (*Identity, error) {
	I := new(Identity)
	switch {
	case name == "":
		return nil, fmt.Errorf("name can not be blank")
	case !accessKeyID.IsValid():
		return nil, fmt.Errorf("%s is not a valid AWS Access Key ID", accessKeyID)
	case !secretAccessKey.IsValid():
		return nil, fmt.Errorf("%s is not a valid AWS Secret Access Key", secretAccessKey)
	default:
		I.Name = name
		I.Region = region
		I.Output = output
		I.AWSAccessKeyID = accessKeyID
		I.AWSSecretAccessKey = secretAccessKey
		return I, nil
	}
}

func (I *Identity) WriteCredentialFile(fileOverride *os.File) (int, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return 0, err
	}

	if fileOverride != nil {
		return fileOverride.WriteString(fmt.Sprintf(credentialsFileContents, I.Name, I.AWSAccessKeyID, I.AWSSecretAccessKey))
	}

	if err := os.MkdirAll(filepath.Join(home, ".aws"), os.ModeDir); err != nil {
		return 0, err
	}
	credentialsFile := filepath.Join(home, ".aws", "credentials")

	credf, err := os.Open(credentialsFile)
	if err != nil {
		return 0, err
	}
	defer credf.Close()

	return credf.WriteString(fmt.Sprintf(credentialsFileContents, I.Name, I.AWSAccessKeyID, I.AWSSecretAccessKey))
}

func (I *Identity) WriteConfigFile(fileOverride *os.File) (int, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return 0, err
	}

	if fileOverride != nil {
		return fileOverride.WriteString(fmt.Sprintf(configFileContents, I.Name, I.Region, I.Output))
	}

	if err := os.MkdirAll(filepath.Join(home, ".aws"), os.ModeDir); err != nil {
		return 0, err
	}
	configFile := filepath.Join(home, ".aws", "config")

	conff, err := os.Open(configFile)
	if err != nil {
		return 0, err
	}
	defer conff.Close()

	return conff.WriteString(fmt.Sprintf(configFileContents, I.Name, I.Region, I.Output))
}
