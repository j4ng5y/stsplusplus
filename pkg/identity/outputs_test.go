package identity

import "testing"

func TestAWSAccessKeyID_IsValid(t *testing.T) {
	goodKey := AWSAccessKeyID("AKIAQQL3AXXXXGANZBET")
	badKeyShort := AWSAccessKeyID("QQL3AB4XXXXNZBET")
	badKeyLong := AWSAccessKeyID("AKIAQQL3AB4JXXXXZBETAKIA")

	if !goodKey.IsValid() {
		t.Logf("%s is not a valid key, but should be", goodKey)
		t.Fail()
	}
	if badKeyShort.IsValid() {
		t.Logf("%s is a valid key, but should not be", badKeyShort)
		t.Fail()
	}
	if badKeyLong.IsValid() {
		t.Logf("%s is a valid key, but should not be", badKeyLong)
		t.Fail()
	}
}

func TestAWSSecretAccessKey_IsValid(t *testing.T) {
	goodKey := AWSSecretAccessKey("/3ru4eUDIXXXXXXXXMkwr0OKZF90eZlkcWNU4Hgi")
	badKeyShort := AWSSecretAccessKey("/3ru4XXXXXXXXUWN0Mkwr0OKZF90eZlkcWNU")
	badKeyLong := AWSSecretAccessKey("/3ru4eUXXXXXXXXN0Mkwr0OKZF90eZlkcWNU4Hgi/3ru")

	if !goodKey.IsValid() {
		t.Logf("%s is not a valid key, but should be", goodKey)
		t.Fail()
	}
	if badKeyShort.IsValid() {
		t.Logf("%s is a valid key, but should not be", badKeyShort)
		t.Fail()
	}
	if badKeyLong.IsValid() {
		t.Logf("%s is a valid key, but should not be", badKeyLong)
		t.Fail()
	}
}
