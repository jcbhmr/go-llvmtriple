package minillvmtargetparser

type ARMISAKind int
const (
	ARMISAKindINVALID ARMISAKind = iota
	ARMISAKindARM
	ARMISAKindTHUMB
	ARMISAKindAARCH64
)

type ARMEndianKind int
const (
	ARMEndianKindINVALID ARMEndianKind = iota
	ARMEndianKindLITTLE
	ARMEndianKindBIG
)

// Converts e.g. "armv8" -> "armv8-a"
func ARMGetArchSynonym(arch string) string {
	panic("not implemented")
}

// MArch is expected to be of the form (arm|thumb)?(eb)?(v.+)?(eb)?, but
// (iwmmxt|xscale)(eb)? is also permitted. If the former, return
// "v.+", if the latter, return unmodified string, minus 'eb'.
// If invalid, return empty string.
func ARMGetCanonicalArchName(arch string) string {
	panic("not implemented")
}

// ARM, Thumb, AArch64
func ARMParseArchISA(arch string) ARMISAKind {
	panic("not implemented")
}

// Little/Big endian
func ARMParseArchEndian(arch string) ARMEndianKind {
	panic("not implemented")
}
