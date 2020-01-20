package dbreport

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersionsSuccess(t *testing.T) {
	tests := []struct {
		SysbenchInfo  *SysbenchInfo
		Version       string
		MajourVersion uint64
		MinorVersion  uint64
		PatchLevel    string
		AddInfo       string
	}{
		{
			SysbenchInfo: &SysbenchInfo{
				Version: &VersionInfo{
					VersionNumber: "1.2.3",
					AddInfo:       []string{""},
				},
			},
			Version:       "1.2.3",
			MajourVersion: 1,
			MinorVersion:  2,
			PatchLevel:    "3",
			AddInfo:       "",
		},
		{
			SysbenchInfo: &SysbenchInfo{
				Version: &VersionInfo{
					VersionNumber: "2.34.5678rc",
					AddInfo:       []string{"9-10-11", "additional", "informations"},
				},
			},
			Version:       "2.34.5678rc",
			MajourVersion: 2,
			MinorVersion:  34,
			PatchLevel:    "5678rc",
			AddInfo:       "9-10-11 additional informations",
		},
	}

	var (
		s   string
		i   uint64
		err error
	)

	for _, test := range tests {
		// Full Version
		s, err = test.SysbenchInfo.GetVersion()
		assert.NoError(t, err)
		assert.Equal(t, test.Version, s, "GetVersion() is failed")

		// Majour Version
		s, err = test.SysbenchInfo.GetMajourVersion()
		assert.NoError(t, err)
		assert.Equal(t, strconv.FormatUint(test.MajourVersion, 10), s, "GetMajourVersion() is failed")

		i, err = test.SysbenchInfo.GetMajourVersionNumber()
		assert.NoError(t, err)
		assert.Equal(t, test.MajourVersion, i, "GetMajourVersionNumber() is failed")

		// Minor Version
		s, err = test.SysbenchInfo.GetMinorVersion()
		assert.NoError(t, err)
		assert.Equal(t, strconv.FormatUint(test.MinorVersion, 10), s, "GetMinorVersion() is failed")

		i, err = test.SysbenchInfo.GetMinorVersionNumber()
		assert.NoError(t, err)
		assert.Equal(t, test.MinorVersion, i, "GetMinorVersionNumber() is failed")

		// PatchLevel
		s, err = test.SysbenchInfo.GetVersionPatchLevel()
		assert.NoError(t, err)
		assert.Equal(t, test.PatchLevel, s, "GetMinorVersionPatchLevel() is failed")

		// AddInfo
		s, err = test.SysbenchInfo.GetVersionAddInfo()
		assert.NoError(t, err)
		assert.Equal(t, test.AddInfo, s, "GetMinorVersionAddInfo() is failed")
	}
}

func TestGetMajourVersionFail(t *testing.T) {
	tests := []*SysbenchInfo{
		{
			Version: &VersionInfo{
				VersionNumber: "",
			},
		},
	}

	for _, test := range tests {
		_, err := test.GetMajourVersion()
		assert.Error(t, err, fmt.Sprintf("Version strings is '%s'", test.Version.VersionNumber))
	}
}

func TestGetMajourVersionNumberFail(t *testing.T) {
	tests := []*SysbenchInfo{
		{
			Version: &VersionInfo{
				VersionNumber: "",
			},
		},
		{
			Version: &VersionInfo{
				VersionNumber: "-1",
			},
		},
		{
			Version: &VersionInfo{
				VersionNumber: "3a",
			},
		},
	}

	for _, test := range tests {
		_, err := test.GetMajourVersionNumber()
		assert.Error(t, err, fmt.Sprintf("Version strings is '%s'", test.Version.VersionNumber))
	}
}

func TestGetMinorVersionFail(t *testing.T) {
	tests := []*SysbenchInfo{
		{
			Version: &VersionInfo{
				VersionNumber: "1",
			},
		},
		{
			Version: &VersionInfo{
				VersionNumber: "1.",
			},
		},
	}

	for _, test := range tests {
		_, err := test.GetMinorVersion()
		assert.Error(t, err, fmt.Sprintf("Version strings is '%s'", test.Version.VersionNumber))
	}
}

func TestGetMinorVersionNumberFail(t *testing.T) {
	tests := []*SysbenchInfo{
		{
			Version: &VersionInfo{
				VersionNumber: "1",
			},
		},
		{
			Version: &VersionInfo{
				VersionNumber: "1.",
			},
		},
		{
			Version: &VersionInfo{
				VersionNumber: "1.-10",
			},
		},
		{
			Version: &VersionInfo{
				VersionNumber: "1.20a",
			},
		},
	}

	for _, test := range tests {
		_, err := test.GetMinorVersionNumber()
		assert.Error(t, err, fmt.Sprintf("Version strings is '%s'", test.Version.VersionNumber))
	}
}

func TestGetVersionPatchLevelFail(t *testing.T) {
	tests := []*SysbenchInfo{
		{
			Version: &VersionInfo{
				VersionNumber: "1.2",
			},
		},
		{
			Version: &VersionInfo{
				VersionNumber: "1.2.",
			},
		},
	}

	for _, test := range tests {
		_, err := test.GetVersionPatchLevel()
		assert.Error(t, err, fmt.Sprintf("Version strings is '%s'", test.Version.VersionNumber))
	}
}
