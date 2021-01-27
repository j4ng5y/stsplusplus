package identity

import (
	"io/ioutil"
	"os"
	"testing"
)

var tests = map[string]map[string]interface{}{
	"good": map[string]interface{}{
		"name":            "default",
		"region":          Region_USEast1,
		"output":          Output_JSON,
		"accessKeyID":     AWSAccessKeyID("AKIAQQL3AB4XXXXXZBET"),
		"secretAccessKey": AWSSecretAccessKey("/3ru4eUXXXXX7UWN0Mkwr0OKZF90eZlkcWNU4Hgi"),
	},
	"bad_noName": map[string]interface{}{
		"name":            "",
		"region":          Region_USEast1,
		"output":          Output_JSON,
		"accessKeyID":     AWSAccessKeyID("AKIAQQL3AB4XXXXXZBET"),
		"secretAccessKey": AWSSecretAccessKey("/3ru4eUXXXXX7UWN0Mkwr0OKZF90eZlkcWNU4Hgi"),
	},
	"bad_invalidAccessKeyID": map[string]interface{}{
		"name":            "default",
		"region":          Region_USEast1,
		"output":          Output_JSON,
		"accessKeyID":     AWSAccessKeyID(""),
		"secretAccessKey": AWSSecretAccessKey("/3ru4eUXXXXX7UWN0Mkwr0OKZF90eZlkcWNU4Hgi"),
	},
	"bad_invalidSecretAccessKey": map[string]interface{}{
		"name":            "default",
		"region":          Region_USEast1,
		"output":          Output_JSON,
		"accessKeyID":     AWSAccessKeyID("AKIAQQL3AB4XXXXXZBET"),
		"secretAccessKey": AWSSecretAccessKey(""),
	},
}

func Test_New(t *testing.T) {
	if _, err := New(
		tests["good"]["name"].(string),
		tests["good"]["region"].(Region),
		tests["good"]["output"].(Output),
		tests["good"]["accessKeyID"].(AWSAccessKeyID),
		tests["good"]["secretAccessKey"].(AWSSecretAccessKey)); err != nil {
		t.Logf("creating good identity struct failed due to error %v", err)
		t.Fail()
	}

	if _, err := New(
		tests["bad_noName"]["name"].(string),
		tests["bad_noName"]["region"].(Region),
		tests["bad_noName"]["output"].(Output),
		tests["bad_noName"]["accessKeyID"].(AWSAccessKeyID),
		tests["bad_noName"]["secretAccessKey"].(AWSSecretAccessKey)); err == nil {
		t.Logf("creating bad identity struct failed to err")
		t.Fail()
	}

	if _, err := New(
		tests["bad_invalidAccessKeyID"]["name"].(string),
		tests["bad_invalidAccessKeyID"]["region"].(Region),
		tests["bad_invalidAccessKeyID"]["output"].(Output),
		tests["bad_invalidAccessKeyID"]["accessKeyID"].(AWSAccessKeyID),
		tests["bad_invalidAccessKeyID"]["secretAccessKey"].(AWSSecretAccessKey)); err == nil {
		t.Logf("creating bad identity struct failed to err")
		t.Fail()
	}

	if _, err := New(
		tests["bad_invalidSecretAccessKey"]["name"].(string),
		tests["bad_invalidSecretAccessKey"]["region"].(Region),
		tests["bad_invalidSecretAccessKey"]["output"].(Output),
		tests["bad_invalidSecretAccessKey"]["accessKeyID"].(AWSAccessKeyID),
		tests["bad_invalidSecretAccessKey"]["secretAccessKey"].(AWSSecretAccessKey)); err == nil {
		t.Logf("creating bad identity struct failed to err")
		t.Fail()
	}
}

func Test_Identity_WriteCredentialFile(t *testing.T) {
	fileContents := `[default]
aws_access_key_id = AKIAQQL3AB4XXXXXZBET
aws_secret_access_key = /3ru4eUXXXXX7UWN0Mkwr0OKZF90eZlkcWNU4Hgi`
	I, err := New(
		tests["good"]["name"].(string),
		tests["good"]["region"].(Region),
		tests["good"]["output"].(Output),
		tests["good"]["accessKeyID"].(AWSAccessKeyID),
		tests["good"]["secretAccessKey"].(AWSSecretAccessKey))
	if err != nil {
		t.Logf("creating good identity struct failed due to error %v", err)
		t.FailNow()
	}
	tf, err := ioutil.TempFile("", "temp_credentials")
	if err != nil {
		t.Logf("failed to write temp file due to error: %v", err)
		t.FailNow()
	}

	defer tf.Close()
	defer os.Remove(tf.Name())

	i, err := I.WriteCredentialFile(tf)
	if err != nil {
		t.Logf("writing the credential file failed due to error: %v", err)
		t.FailNow()
	}
	if i <= 0 {
		t.Logf("writing the credential file wrote %d bytes", i)
		t.FailNow()
	}
	b, err := ioutil.ReadFile(tf.Name())
	if err != nil {
		t.Logf("failed to read temp file due to error: %v", err)
		t.FailNow()
	}
	if string(b) != fileContents {
		t.Logf("expected:\n%s\ngot:\n%s", fileContents, string(b))
		t.FailNow()
	}
}

func Test_Identity_WriteConfigFile(t *testing.T) {
	fileContents := `[default]
region = us-east-1
output = json`
	I, err := New(
		tests["good"]["name"].(string),
		tests["good"]["region"].(Region),
		tests["good"]["output"].(Output),
		tests["good"]["accessKeyID"].(AWSAccessKeyID),
		tests["good"]["secretAccessKey"].(AWSSecretAccessKey))
	if err != nil {
		t.Logf("creating good identity struct failed due to error %v", err)
		t.FailNow()
	}
	tf, err := ioutil.TempFile("", "temp_config")
	if err != nil {
		t.Logf("failed to write temp file due to error: %v", err)
		t.FailNow()
	}

	defer tf.Close()
	defer os.Remove(tf.Name())

	i, err := I.WriteConfigFile(tf)
	if err != nil {
		t.Logf("writing the credential file failed due to error: %v", err)
		t.FailNow()
	}
	if i <= 0 {
		t.Logf("writing the credential file wrote %d bytes", i)
		t.FailNow()
	}
	b, err := ioutil.ReadFile(tf.Name())
	if err != nil {
		t.Logf("failed to read temp file due to error: %v", err)
		t.FailNow()
	}
	if string(b) != fileContents {
		t.Logf("expected:\n%s\ngot:\n%s", fileContents, string(b))
		t.FailNow()
	}
}
