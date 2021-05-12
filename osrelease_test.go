package osrelease

import (
	"fmt"
	"os"
	"testing"
)

// This test is written for Ubuntu 21.04
func TestParse(t *testing.T) {

	release, err := Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse os-relese: %s\n", err)
		os.Exit(1)
	}

	switch true {
	case release.Name != "Ubuntu":
		t.Errorf("Test failed on NAME: want 'Ubuntu', got '%s'\n", release.Name)
	case release.Version != "21.04 (Hirsute Hippo)":
		t.Errorf("Test failed on VERSION: want '21.04 (Hirsute Hippo)', got '%s'\n", release.Version)
	case release.ID != "ubuntu":
		t.Errorf("Test failed on ID: want 'ubuntu', got '%s'\n", release.ID)
	case release.IDLike != "debian":
		t.Errorf("Test failed on ID_LIKE: want 'debian', got '%s'\n", release.IDLike)
	case release.PrettyName != "Ubuntu 21.04":
		t.Errorf("Test failed on PRETTY_NAME: want 'Ubuntu 21.04', got '%s'\n", release.PrettyName)
	case release.VersionID != "21.04":
		t.Errorf("Test failed on VERSION_ID: want '21.04', got '%s'\n", release.VersionID)
	case release.HomeURL != "https://www.ubuntu.com/":
		t.Errorf("test failed on HOME_URL: want 'https://www.ubuntu.com/', got '%s'\n", release.HomeURL)
	case release.DocumentationURL != "":
		t.Errorf("Test failed on DOCUMENTATION_URL: want an empty string, got '%s'\n", release.DocumentationURL)
	case release.SupportURL != "https://help.ubuntu.com/":
		t.Errorf("Test failed on SUPPORT_URL: want 'https://help.ubuntu.com/', got '%s'\n", release.SupportURL)
	case release.BugReportURL != "https://bugs.launchpad.net/ubuntu/":
		t.Errorf("Test failed on BUG_REPORT_URL: want 'https://bugs.launchpad.net/ubuntu/', got '%s'\n", release.BugReportURL)
	case release.PrivacyPolicyURL != "https://www.ubuntu.com/legal/terms-and-policies/privacy-policy":
		t.Errorf("Test failed on PRIVACY_POLICY_URL: want 'https://www.ubuntu.com/legal/terms-and-policies/privacy-policy', got '%s'\n", release.PrivacyPolicyURL)
	case release.VersionCodename != "hirsute":
		t.Errorf("Test failed on VERSION_CODENAME: want 'hirsuite', got '%s'\n", release.VersionCodename)
	case release.UbuntuCodename != "hirsute":
		t.Errorf("Test failed on UBUNTU_CODENAME: want 'hirsuite', got '%s'\n", release.UbuntuCodename)
	case release.ANSIColor != "":
		t.Errorf("Test failed on ANSI_COLOR: want an empty string, got '%s'\n", release.ANSIColor)
	case release.CPEName != "":
		t.Errorf("Test failed on CPE_NAME: want an empty string, got '%s'\n", release.CPEName)
	case release.BuildID != "":
		t.Errorf("Test failed on BUILD_ID: want an empty string, got '%s'\n", release.BuildID)
	case release.Variant != "":
		t.Errorf("Test failed on VARIANT: want an empty string, got '%s'\n", release.Variant)
	case release.VariantID != "":
		t.Errorf("Test failed on VARIANT_ID: want an empty string, got '%s'\n", release.VariantID)
	case release.Logo != "":
		t.Errorf("Test failed on LOGO: want an empty string, got '%s'\n", release.Logo)
	}
}