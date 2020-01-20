package dbreport

import (
	"fmt"
	"strconv"
	"strings"
)

// GetVersion retrun Version String
func (s *SysbenchInfo) GetVersion() (string, error) {
	return s.Version.VersionNumber, nil
}

// GetMajourVersion retrun Majour Version String
func (s *SysbenchInfo) GetMajourVersion() (string, error) {
	v := strings.Split(s.Version.VersionNumber, ".")
	if v[0] == "" {
		return "", fmt.Errorf("get majour version failed from %s", s.Version.VersionNumber)
	}

	return v[0], nil
}

// GetMajourVersionNumber retrun Majour Version Number
func (s *SysbenchInfo) GetMajourVersionNumber() (uint64, error) {
	var i uint64

	v := strings.Split(s.Version.VersionNumber, ".")
	if v[0] == "" {
		return i, fmt.Errorf("get majour version failed from %s", s.Version.VersionNumber)
	}

	return strconv.ParseUint(v[0], 10, 64)
}

// GetMinorVersion retrun Minor Version Number
func (s *SysbenchInfo) GetMinorVersion() (string, error) {
	v := strings.Split(s.Version.VersionNumber, ".")
	if len(v) < 2 || v[1] == "" {
		return "", fmt.Errorf("get minor version failed from %s", s.Version.VersionNumber)
	}

	return v[1], nil
}

// GetMinorVersionNumber retrun Minor Version Number
func (s *SysbenchInfo) GetMinorVersionNumber() (uint64, error) {
	var i uint64

	v := strings.Split(s.Version.VersionNumber, ".")
	if len(v) < 2 || v[1] == "" {
		return i, fmt.Errorf("get minor version failed from %s", s.Version.VersionNumber)
	}

	return strconv.ParseUint(v[1], 10, 64)
}

// GetVersionPatchLevel retrun Version PatchLevel String
func (s *SysbenchInfo) GetVersionPatchLevel() (string, error) {
	v := strings.Split(s.Version.VersionNumber, ".")
	if len(v) < 3 || v[2] == "" {
		return "", fmt.Errorf("get patch level failed from %s", s.Version.VersionNumber)
	}

	return v[2], nil
}

// GetVersionAddInfo retrun Version Add Information String
func (s *SysbenchInfo) GetVersionAddInfo() (string, error) {
	return strings.Join(s.Version.AddInfo, " "), nil
}
