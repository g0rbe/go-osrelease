package osrelease

import (
	"fmt"
	"os"
	"testing"
)

// This test is written for Ubuntu 21.04
func TestParse(t *testing.T) {

	err := Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse os-relese: %s\n", err)
		os.Exit(1)
	}

	switch true {
	case Release.Name != "Ubuntu":
		t.Errorf("Test failed on NAME: want 'Ubuntu', got '%s'\n", Release.Name)
	case Release.Version != "21.04 (Hirsute Hippo)":
		t.Errorf("Test failed on VERSION: want '21.04 (Hirsute Hippo)', got '%s'\n", Release.Version)
	case Release.ID != "ubuntu":
		t.Errorf("Test failed on ID: want 'ubuntu', got '%s'\n", Release.ID)
	case Release.IDLike != "debian":
		t.Errorf("Test failed on ID_LIKE: want 'debian', got '%s'\n", Release.IDLike)
	case Release.PrettyName != "Ubuntu 21.04":
		t.Errorf("Test failed on PRETTY_NAME: want 'Ubuntu 21.04', got '%s'\n", Release.PrettyName)
	case Release.VersionID != "21.04":
		t.Errorf("Test failed on VERSION_ID: want '21.04', got '%s'\n", Release.VersionID)
	case Release.HomeURL != "https://www.ubuntu.com/":
		t.Errorf("test failed on HOME_URL: want 'https://www.ubuntu.com/', got '%s'\n", Release.HomeURL)
	case Release.DocumentationURL != "":
		t.Errorf("Test failed on DOCUMENTATION_URL: want an empty string, got '%s'\n", Release.DocumentationURL)
	case Release.SupportURL != "https://help.ubuntu.com/":
		t.Errorf("Test failed on SUPPORT_URL: want 'https://help.ubuntu.com/', got '%s'\n", Release.SupportURL)
	case Release.BugReportURL != "https://bugs.launchpad.net/ubuntu/":
		t.Errorf("Test failed on BUG_REPORT_URL: want 'https://bugs.launchpad.net/ubuntu/', got '%s'\n", Release.BugReportURL)
	case Release.PrivacyPolicyURL != "https://www.ubuntu.com/legal/terms-and-policies/privacy-policy":
		t.Errorf("Test failed on PRIVACY_POLICY_URL: want 'https://www.ubuntu.com/legal/terms-and-policies/privacy-policy', got '%s'\n", Release.PrivacyPolicyURL)
	case Release.VersionCodename != "hirsute":
		t.Errorf("Test failed on VERSION_CODENAME: want 'hirsuite', got '%s'\n", Release.VersionCodename)
	case Release.UbuntuCodename != "hirsute":
		t.Errorf("Test failed on UBUNTU_CODENAME: want 'hirsuite', got '%s'\n", Release.UbuntuCodename)
	case Release.ANSIColor != "":
		t.Errorf("Test failed on ANSI_COLOR: want an empty string, got '%s'\n", Release.ANSIColor)
	case Release.CPEName != "":
		t.Errorf("Test failed on CPE_NAME: want an empty string, got '%s'\n", Release.CPEName)
	case Release.BuildID != "":
		t.Errorf("Test failed on BUILD_ID: want an empty string, got '%s'\n", Release.BuildID)
	case Release.Variant != "":
		t.Errorf("Test failed on VARIANT: want an empty string, got '%s'\n", Release.Variant)
	case Release.VariantID != "":
		t.Errorf("Test failed on VARIANT_ID: want an empty string, got '%s'\n", Release.VariantID)
	case Release.Logo != "":
		t.Errorf("Test failed on LOGO: want an empty string, got '%s'\n", Release.Logo)
	}
}