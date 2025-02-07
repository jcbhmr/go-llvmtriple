package support

import (
	"cmp"
	"strconv"
)

// Represents a version number in the form major[.minor[.subminor[.build]]].
type VersionTuple struct {
	major       uint
	minor       uint
	hasMinor    bool
	subMinor    uint
	hasSubMinor bool
	build       uint
	hasBuild    bool
}

func NewVersionTuple() VersionTuple {
	return VersionTuple{
		major:       0,
		minor:       0,
		hasMinor:    false,
		subMinor:    0,
		hasSubMinor: false,
		build:       0,
		hasBuild:    false,
	}
}

func NewVersionTuple2(major uint) VersionTuple {
	return VersionTuple{
		major:       major,
		minor:       0,
		hasMinor:    false,
		subMinor:    0,
		hasSubMinor: false,
		build:       0,
		hasBuild:    false,
	}
}

func NewVersionTuple3(major uint, minor uint) VersionTuple {
	return VersionTuple{
		major:       major,
		minor:       minor,
		hasMinor:    true,
		subMinor:    0,
		hasSubMinor: false,
		build:       0,
		hasBuild:    false,
	}
}

func NewVersionTuple4(major uint, minor uint, subMinor uint) VersionTuple {
	return VersionTuple{
		major:       major,
		minor:       minor,
		hasMinor:    true,
		subMinor:    subMinor,
		hasSubMinor: true,
		build:       0,
		hasBuild:    false,
	}
}

func NewVersionTuple5(major uint, minor uint, subMinor uint, build uint) VersionTuple {
	return VersionTuple{
		major:       major,
		minor:       minor,
		hasMinor:    true,
		subMinor:    subMinor,
		hasSubMinor: true,
		build:       build,
		hasBuild:    true,
	}
}

// Determine whether this version information is empty
// (e.g., all version components are zero).
func (v VersionTuple) Empty() bool {
	return v.major == 0 && v.minor == 0 && v.subMinor == 0 && v.build == 0
}

// Retrieve the major version number.
func (v VersionTuple) Major() uint {
	return v.major
}

// Retrieve the minor version number, if provided.
func (v VersionTuple) Minor() (uint, bool) {
	if !v.hasMinor {
		return 0, false
	}
	return v.minor, true
}

// Retrieve the subminor version number, if provided.
func (v VersionTuple) SubMinor() (uint, bool) {
	if !v.hasSubMinor {
		return 0, false
	}
	return v.subMinor, true
}

// Retrieve the build version number, if provided.
func (v VersionTuple) Build() (uint, bool) {
	if !v.hasBuild {
		return 0, false
	}
	return v.build, true
}

// snip

// Determine if two version numbers are equivalent. If not
// provided, minor and subminor version numbers are considered to be zero.
func (v VersionTuple) Equal(other VersionTuple) bool {
	return v.major == other.major && v.minor == other.minor && v.subMinor == other.subMinor && v.build == other.build
}

func (v VersionTuple) Cmp(other VersionTuple) int {
	return cmp.Or(cmp.Compare(v.major, other.major), cmp.Or(cmp.Compare(v.minor, other.minor), cmp.Or(cmp.Compare(v.subMinor, other.subMinor), cmp.Compare(v.build, other.build))))
}

func (v VersionTuple) String() string {
	var result string
	result = strconv.FormatUint(uint64(v.major), 10)
	if v.hasMinor {
		result += "." + strconv.FormatUint(uint64(v.minor), 10)
		if v.hasSubMinor {
			result += "." + strconv.FormatUint(uint64(v.subMinor), 10)
			if v.hasBuild {
				result += "." + strconv.FormatUint(uint64(v.build), 10)
			}
		}
	}
	return result
}
