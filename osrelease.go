// Package osrelease parse /etc/os-release
// More about the os-release: https://www.linux.org/docs/man5/os-release.html
package osrelease

import (
	"fmt"
	"os"
	"strings"
)

// OSReleasePath contains the default path to the os-release file
var OSReleasePath string = "/etc/os-release"

type OSRelease struct {
	Name string
	Version string
	ID string
	IDLike string
	PrettyName string
	VersionID string
	HomeURL string
	DocumentationURL string
	SupportURL string
	BugReportURL string
	PrivacyPolicyURL string
	VersionCodename string
	UbuntuCodename string
	ANSIColor string
	CPEName string
	BuildID string
	Variant string
	VariantID string
	Logo string
}

// getLines read the OSReleasePath and return it line by line.
// Empty lines and comments (beginning with a "#") are ignored.
func getLines() ([]string, error) {

	output, err := os.ReadFile(OSReleasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %s", OSReleasePath, err)
	}

	lines := make([]string, 0)

	for _, line := range strings.Split(string(output), "\n") {

		switch true {
		case line == "":
			continue
		case []byte(line)[0] == '#':
			continue
		}

		lines = append(lines, line)
	}

	return lines, nil
}

// parseLine parse a single line.
// Return key, value, error (if any)
func parseLine(line string) (string, string, error) {

	subs := strings.SplitN(line, "=", 2)

	if len(subs) != 2 {
		return "", "", fmt.Errorf("invalid length of the substrings: %d", len(subs))
	}

	return subs[0], strings.Trim(subs[1], "\"'"), nil
}

// Parse parses the os-release file pointing to by OSReleasePath.
// Return a struct containing the known fields.
// Fields that not exist will be an empty string.
func Parse() (OSRelease, error) {

	lines, err := getLines()
	if err != nil {
		return OSRelease{}, fmt.Errorf("failed to get lines of %s: %s", OSReleasePath, err)
	}

	var release OSRelease

	for i := range lines {

		key, value, err := parseLine(lines[i])
		if err != nil {
			return OSRelease{}, fmt.Errorf("failed to parse line '%s': %s", lines[i], err)
		}

		switch key {
		case "NAME":
			release.Name = value
		case "VERSION":
			release.Version = value
		case "ID":
			release.ID = value
		case "ID_LIKE":
			release.IDLike = value
		case "PRETTY_NAME":
			release.PrettyName = value
		case "VERSION_ID":
			release.VersionID = value
		case "HOME_URL":
			release.HomeURL = value
		case "DOCUMENTATION_URL":
			release.DocumentationURL = value
		case "SUPPORT_URL":
			release.SupportURL = value
		case "BUG_REPORT_URL":
			release.BugReportURL = value
		case "PRIVACY_POLICY_URL":
			release.PrivacyPolicyURL = value
		case "VERSION_CODENAME":
			release.VersionCodename = value
		case "UBUNTU_CODENAME":
			release.UbuntuCodename = value
		case "ANSI_COLOR":
			release.ANSIColor = value
		case "CPE_NAME":
			release.CPEName = value
		case "BUILD_ID":
			release.BuildID = value
		case "VARIANT":
			release.Variant = value
		case "VARIANT_ID":
			release.VariantID = value
		case "LOGO":
			release.Logo = value
		default:
			return OSRelease{}, fmt.Errorf("unknown key found: %s", key)
		}
	}

	return release, nil
}