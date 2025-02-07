package minillvmtargetparser_test

import (
	"testing"

	"github.com/jcbhmr/go-minillvmtargetparser/v19"
	"github.com/jcbhmr/go-minillvmtargetparser/v19/support"
	"github.com/stretchr/testify/assert"
)

func TestBasicParsing(t *testing.T) {
	table := []struct {
		triple          string
		archName        string
		vendorName      string
		osName          string
		environmentName string
	}{
		{"", "", "", "", ""},
		{"-", "", "", "", ""},
		{"--", "", "", "", ""},
		{"---", "", "", "", ""},
		{"----", "", "", "", "-"},
		{"a", "a", "", "", ""},
		{"a-b", "a", "b", "", ""},
		{"a-b-c", "a", "b", "c", ""},
		{"a-b-c-d", "a", "b", "c", "d"},
	}
	for _, tt := range table {
		u := minillvmtargetparser.NewTriple2(tt.triple)
		assert.Equal(t, tt.archName, u.ArchName())
		assert.Equal(t, tt.vendorName, u.VendorName())
		assert.Equal(t, tt.osName, u.OSName())
		assert.Equal(t, tt.environmentName, u.EnvironmentName())
	}
}

func TestParsedIDs(t *testing.T) {

}

func TestNormalization(t *testing.T) {

}

func TestBitWidthChecks(t *testing.T) {

}

func TestBitWidthArchVariants(t *testing.T) {

}

func TestEndianArchVariants(t *testing.T) {

}

func TestXROS(t *testing.T) {
	var u *minillvmtargetparser.Triple
	var version support.VersionTuple

	u = minillvmtargetparser.NewTriple2("arm64-apple-xros")
	assert.True(t, u.IsXROS())
	assert.True(t, u.IsOSDarwin())
	assert.False(t, u.IsiOS())
	assert.False(t, u.IsMacOSX())
	assert.False(t, u.IsSimulatorEnvironment())
	assert.Equal(t, "xros", u.OSName())
	version = u.OSVersion()
	assert.True(t, support.NewVersionTuple2(0).Equal(version))

	u = minillvmtargetparser.NewTriple2("arm64-apple-visionos1.2")
	assert.True(t, u.IsXROS())
	assert.True(t, u.IsOSDarwin())
	assert.False(t, u.IsiOS())
	assert.False(t, u.IsMacOSX())
	assert.False(t, u.IsSimulatorEnvironment())
	assert.Equal(t, "visionos1.2", u.OSName())
	version = u.OSVersion()
	assert.True(t, support.NewVersionTuple3(1, 2).Equal(version))

	u = minillvmtargetparser.NewTriple2("arm64-apple-xros1-simulator")
	assert.True(t, u.IsXROS())
	assert.True(t, u.IsOSDarwin())
	assert.False(t, u.IsiOS())
	assert.False(t, u.IsMacOSX())
	assert.True(t, u.IsSimulatorEnvironment())
	version = u.OSVersion()
	assert.True(t, support.NewVersionTuple2(1).Equal(version))
	version = u.IOSVersion()
	assert.True(t, support.NewVersionTuple2(17).Equal(version))
}
