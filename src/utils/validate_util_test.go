package utils

import (
	"testing"
)

func TestGetRandomString(t *testing.T) {
	if ValidateScriptName("name") != true {
		t.Error("ValidateScriptName failed")
	}

	if ValidateScriptName(".name") != false {
		t.Error("ValidateScriptName failed")
	}

	if ValidateScriptName("../name") != false {
		t.Error("ValidateScriptName failed")
	}
	if ValidateScriptName("8name") != true {
		t.Error("ValidateScriptName failed")
	}

	if ValidateScriptName("") != false {
		t.Error("ValidateScriptName failed")
	}

	if ValidateScriptName("_") != true {
		t.Error("ValidateScriptName failed")
	}

	if ValidateScriptName("-") != true {
		t.Error("ValidateScriptName failed")
	}
}
