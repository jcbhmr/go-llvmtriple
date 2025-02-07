package minillvmtargetparser

type ARMArchExtKind uint64

const (
	ARMAEK_INVALID    = 0
	ARMAEK_NONE       = 1
	ARMAEK_CRC        = 1 << 1
	ARMAEK_CRYPTO     = 1 << 2
	ARMAEK_FP         = 1 << 3
	ARMAEK_HWDIVTHUMB = 1 << 4
	ARMAEK_HWDIVARM   = 1 << 5
	ARMAEK_MP         = 1 << 6
	ARMAEK_SIMD       = 1 << 7
	ARMAEK_SEC        = 1 << 8
	ARMAEK_VIRT       = 1 << 9
	ARMAEK_DSP        = 1 << 10
	ARMAEK_FP16       = 1 << 11
	ARMAEK_RAS        = 1 << 12
	ARMAEK_DOTPROD    = 1 << 13
	ARMAEK_SHA2       = 1 << 14
	ARMAEK_AES        = 1 << 15
	ARMAEK_FP16FML    = 1 << 16
	ARMAEK_SB         = 1 << 17
	ARMAEK_FP_DP      = 1 << 18
	ARMAEK_LOB        = 1 << 19
	ARMAEK_BF16       = 1 << 20
	ARMAEK_I8MM       = 1 << 21
	ARMAEK_CDECP0     = 1 << 22
	ARMAEK_CDECP1     = 1 << 23
	ARMAEK_CDECP2     = 1 << 24
	ARMAEK_CDECP3     = 1 << 25
	ARMAEK_CDECP4     = 1 << 26
	ARMAEK_CDECP5     = 1 << 27
	ARMAEK_CDECP6     = 1 << 28
	ARMAEK_CDECP7     = 1 << 29
	ARMAEK_PACBTI     = 1 << 30
	// Unsupported extensions.
	ARMAEK_OS       = 1 << 59
	ARMAEK_IWMMXT   = 1 << 60
	ARMAEK_IWMMXT2  = 1 << 61
	ARMAEK_MAVERICK = 1 << 62
	ARMAEK_XSCALE   = 1 << 63
)

// List of Arch Extension names.
type ARMExtName struct {
	Name       string
	ID         uint64
	Feature    string
	NegFeature string
}

func ARMARCHExtNames() []ARMExtName {
	return []ARMExtName{
		{"invalid", ARMAEK_INVALID, "", ""},
		{"none", ARMAEK_NONE, "", ""},
		{"crc", ARMAEK_CRC, "+crc", "-crc"},
		{"crypto", ARMAEK_CRYPTO, "+crypto", "-crypto"},
		{"sha2", ARMAEK_SHA2, "+sha2", "-sha2"},
		{"aes", ARMAEK_AES, "+aes", "-aes"},
		{"dotprod", ARMAEK_DOTPROD, "+dotprod", "-dotprod"},
		{"dsp", ARMAEK_DSP, "+dsp", "-dsp"},
		{"fp", ARMAEK_FP, "", ""},
		{"fp.dp", ARMAEK_FP_DP, "", ""},
		{"mve", (ARMAEK_DSP | ARMAEK_SIMD), "+mve", "-mve"},
		{"mve.fp", (ARMAEK_DSP | ARMAEK_SIMD | ARMAEK_FP), "+mve.fp", "-mve.fp"},
		{"idiv", (ARMAEK_HWDIVARM | ARMAEK_HWDIVTHUMB), "", ""},
		{"mp", ARMAEK_MP, "", ""},
		{"simd", ARMAEK_SIMD, "", ""},
		{"sec", ARMAEK_SEC, "", ""},
		{"virt", ARMAEK_VIRT, "", ""},
		{"fp16", ARMAEK_FP16, "+fullfp16", "-fullfp16"},
		{"ras", ARMAEK_RAS, "+ras", "-ras"},
		{"os", ARMAEK_OS, "", ""},
		{"iwmmxt", ARMAEK_IWMMXT, "", ""},
		{"iwmmxt2", ARMAEK_IWMMXT2, "", ""},
		{"maverick", ARMAEK_MAVERICK, "", ""},
		{"xscale", ARMAEK_XSCALE, "", ""},
		{"fp16fml", ARMAEK_FP16FML, "+fp16fml", "-fp16fml"},
		{"bf16", ARMAEK_BF16, "+bf16", "-bf16"},
		{"sb", ARMAEK_SB, "+sb", "-sb"},
		{"i8mm", ARMAEK_I8MM, "+i8mm", "-i8mm"},
		{"lob", ARMAEK_LOB, "+lob", "-lob"},
		{"cdecp0", ARMAEK_CDECP0, "+cdecp0", "-cdecp0"},
		{"cdecp1", ARMAEK_CDECP1, "+cdecp1", "-cdecp1"},
		{"cdecp2", ARMAEK_CDECP2, "+cdecp2", "-cdecp2"},
		{"cdecp3", ARMAEK_CDECP3, "+cdecp3", "-cdecp3"},
		{"cdecp4", ARMAEK_CDECP4, "+cdecp4", "-cdecp4"},
		{"cdecp5", ARMAEK_CDECP5, "+cdecp5", "-cdecp5"},
		{"cdecp6", ARMAEK_CDECP6, "+cdecp6", "-cdecp6"},
		{"cdecp7", ARMAEK_CDECP7, "+cdecp7", "-cdecp7"},
		{"pacbti", ARMAEK_PACBTI, "+pacbti", "-pacbti"},
	}
}

// List of HWDiv names (use getHWDivSynonym) and which architectural
// features they correspond to (use getHWDivFeatures).
func ARMHWDivNames() []struct{Name string; ID uint64} {
	return []struct{Name string; ID uint64}{
		{"invalid", ARMAEK_INVALID},
		{"none", ARMAEK_NONE},
		{"thumb", ARMAEK_HWDIVTHUMB},
		{"arm", ARMAEK_HWDIVARM},
		{"arm,thumb", (ARMAEK_HWDIVARM | ARMAEK_HWDIVTHUMB)},
	}
}

// Arch names.
type ARMArchKind int
const (
	ARMArchKindINVALID ARMArchKind = iota
	ARMArchKindARMV4
	ARMArchKindARMV4T
	ARMArchKindARMV5T
	ARMArchKindARMV5TE
	ARMArchKindARMV5TEJ
	ARMArchKindARMV6
	ARMArchKindARMV6K
	ARMArchKindARMV6T2
	ARMArchKindARMV6KZ
	ARMArchKindARMV6M
	ARMArchKindARMV7A
	ARMArchKindARMV7VE
	ARMArchKindARMV7R
	ARMArchKindARMV7M
	ARMArchKindARMV7EM
	ARMArchKindARMV8A
	ARMArchKindARMV8_1A
	ARMArchKindARMV8_2A
	ARMArchKindARMV8_3A
	ARMArchKindARMV8_4A
	ARMArchKindARMV8_5A
	ARMArchKindARMV8_6A
	ARMArchKindARMV8_7A
	ARMArchKindARMV8_8A
	ARMArchKindARMV8_9A
	ARMArchKindARMV9A
	ARMArchKindARMV9_1A
	ARMArchKindARMV9_2A
	ARMArchKindARMV9_3A
	ARMArchKindARMV9_4A
	ARMArchKindARMV9_5A
	ARMArchKindARMV8R
	ARMArchKindARMV8MBaseline
	ARMArchKindARMV8MMainline
	ARMArchKindARMV8_1MMainline
	// Non-standard Arch names.
	ARMArchKindIWMMXT
	ARMArchKindIWMMXT2
	ARMArchKindXSCALE
	ARMArchKindARMV7S
	ARMArchKindARMV7K
)

// List of CPU names and their arches.
// The same CPU can have multiple arches and can be default on multiple arches.
// When finding the Arch for a CPU, first-found prevails. Sort them accordingly.
// When this becomes table-generated, we'd probably need two tables.
type ARMCpuNames struct {
	Name string
	ArchID ARMArchKind
	Default bool // is $Name the default CPU for $ArchID ?
	DefaultExtensions uint64
}

func ARMCPUNames() []ARMCpuNames {
	return []ARMCpuNames{
		{"invalid", ARMArchKindINVALID, "", "+", ARMBuildAttrs::CPUArch::Pre_v4, FK_NONE, ARMAEK_NONE}
		{"armv4", ARMArchKindARMV4, "4", "+v4", ARMBuildAttrs::CPUArch::v4, FK_NONE, ARMAEK_NONE}
		{"armv4t", ARMArchKindARMV4T, "4T", "+v4t", ARMBuildAttrs::CPUArch::v4T, FK_NONE, ARMAEK_NONE}
		{"armv5t", ARMArchKindARMV5T, "5T", "+v5", ARMBuildAttrs::CPUArch::v5T, FK_NONE, ARMAEK_NONE}
		{"armv5te", ARMArchKindARMV5TE, "5TE", "+v5e", ARMBuildAttrs::CPUArch::v5TE, FK_NONE, ARMAEK_DSP}
		{"armv5tej", ARMArchKindARMV5TEJ, "5TEJ", "+v5e", ARMBuildAttrs::CPUArch::v5TEJ, FK_NONE, ARMAEK_DSP}
		{"armv6", ARMArchKindARMV6, "6", "+v6", ARMBuildAttrs::CPUArch::v6, FK_VFPV2, ARMAEK_DSP}
		{"armv6k", ARMArchKindARMV6K, "6K", "+v6k", ARMBuildAttrs::CPUArch::v6K, FK_VFPV2, ARMAEK_DSP}
		{"armv6t2", ARMArchKindARMV6T2, "6T2", "+v6t2", ARMBuildAttrs::CPUArch::v6T2, FK_NONE, ARMAEK_DSP}
		{"armv6kz", ARMArchKindARMV6KZ, "6KZ", "+v6kz", ARMBuildAttrs::CPUArch::v6KZ, FK_VFPV2, (ARMAEK_SEC | ARMAEK_DSP)}
		{"armv6-m", ARMArchKindARMV6M, "6-M", "+v6m", ARMBuildAttrs::CPUArch::v6_M, FK_NONE, ARMAEK_NONE}
		{"armv7-a", ARMArchKindARMV7A, "7-A", "+v7", ARMBuildAttrs::CPUArch::v7, FK_NEON, ARMAEK_DSP}
		{"armv7ve", ARMArchKindARMV7VE, "7VE", "+v7ve", ARMBuildAttrs::CPUArch::v7, FK_NEON, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP)}
		{"armv7-r", ARMArchKindARMV7R, "7-R", "+v7r", ARMBuildAttrs::CPUArch::v7, FK_NONE, (ARMAEK_HWDIVTHUMB | ARMAEK_DSP)}
		{"armv7-m", ARMArchKindARMV7M, "7-M", "+v7m", ARMBuildAttrs::CPUArch::v7, FK_NONE, ARMAEK_HWDIVTHUMB}
		{"armv7e-m", ARMArchKindARMV7EM, "7E-M", "+v7em", ARMBuildAttrs::CPUArch::v7E_M, FK_NONE, (ARMAEK_HWDIVTHUMB | ARMAEK_DSP)}
		{"armv8-a", ARMArchKindARMV8A, "8-A", "+v8a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC)}
		{"armv8.1-a", ARMArchKindARMV8_1A, "8.1-A", "+v8.1a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC)}
		{"armv8.2-a", ARMArchKindARMV8_2A, "8.2-A", "+v8.2a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS)}
		{"armv8.3-a", ARMArchKindARMV8_3A, "8.3-A", "+v8.3a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS)}
		{"armv8.4-a", ARMArchKindARMV8_4A, "8.4-A", "+v8.4a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD)}
		{"armv8.5-a", ARMArchKindARMV8_5A, "8.5-A", "+v8.5a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD)}
		{"armv8.6-a", ARMArchKindARMV8_6A, "8.6-A", "+v8.6a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_I8MM)}
		{"armv8.7-a", ARMArchKindARMV8_7A, "8.7-A", "+v8.7a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_I8MM)}
		{"armv8.8-a", ARMArchKindARMV8_8A, "8.8-A", "+v8.8a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_SHA2 | ARMAEK_AES |  ARMAEK_I8MM)}
		{"armv8.9-a", ARMArchKindARMV8_9A, "8.9-A", "+v8.9a", ARMBuildAttrs::CPUArch::v8_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_SHA2 | ARMAEK_AES |  ARMAEK_I8MM)}
		{"armv9-a", ARMArchKindARMV9A, "9-A", "+v9a", ARMBuildAttrs::CPUArch::v9_A, FK_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD)}
		{"armv9.1-a", ARMArchKindARMV9_1A, "9.1-A", "+v9.1a", ARMBuildAttrs::CPUArch::v9_A, FK_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_I8MM)}
		{"armv9.2-a", ARMArchKindARMV9_2A, "9.2-A", "+v9.2a", ARMBuildAttrs::CPUArch::v9_A, FK_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_I8MM)}
		{"armv9.3-a", ARMArchKindARMV9_3A, "9.3-A", "+v9.3a", ARMBuildAttrs::CPUArch::v9_A, FK_CRYPTO_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_I8MM)}
		{"armv9.4-a", ARMArchKindARMV9_4A, "9.4-A", "+v9.4a", ARMBuildAttrs::CPUArch::v9_A, FK_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_I8MM)}
		{"armv9.5-a", ARMArchKindARMV9_5A, "9.5-A", "+v9.5a", ARMBuildAttrs::CPUArch::v9_A, FK_NEON_FP_ARMV8, (ARMAEK_SEC | ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC | ARMAEK_RAS |  ARMAEK_DOTPROD | ARMAEK_BF16 | ARMAEK_I8MM)}
		{"armv8-r", ARMArchKindARMV8R, "8-R", "+v8r", ARMBuildAttrs::CPUArch::v8_R, FK_FPV5_SP_D16, (ARMAEK_MP | ARMAEK_VIRT | ARMAEK_HWDIVARM |  ARMAEK_HWDIVTHUMB | ARMAEK_DSP | ARMAEK_CRC)}
		{"armv8-m.base", ARMArchKindARMV8MBaseline, "8-M.Baseline", "+v8m.base", ARMBuildAttrs::CPUArch::v8_M_Base, FK_NONE, ARMAEK_HWDIVTHUMB}
		{"armv8-m.main", ARMArchKindARMV8MMainline, "8-M.Mainline", "+v8m.main", ARMBuildAttrs::CPUArch::v8_M_Main, FK_FPV5_D16, ARMAEK_HWDIVTHUMB}
		{"armv8.1-m.main", ARMArchKindARMV8_1MMainline, "8.1-M.Mainline", "+v8.1m.main", ARMBuildAttrs::CPUArch::v8_1_M_Main, FK_FP_ARMV8_FULLFP16_SP_D16, ARMAEK_HWDIVTHUMB | ARMAEK_RAS | ARMAEK_LOB}
		// Non-standard Arch names.
		{"iwmmxt", ARMArchKindIWMMXT, "iwmmxt", "+", ARMBuildAttrs::CPUArch::v5TE, FK_NONE, ARMAEK_NONE}
		{"iwmmxt2", ARMArchKindIWMMXT2, "iwmmxt2", "+", ARMBuildAttrs::CPUArch::v5TE, FK_NONE, ARMAEK_NONE}
		{"xscale", ARMArchKindXSCALE, "xscale", "+v5e", ARMBuildAttrs::CPUArch::v5TE, FK_NONE, ARMAEK_NONE}
		{"armv7s", ARMArchKindARMV7S, "7-S", "+v7s", ARMBuildAttrs::CPUArch::v7, FK_NEON_VFPV4, ARMAEK_DSP}
		{"armv7k", ARMArchKindARMV7K, "7-K", "+v7k", ARMBuildAttrs::CPUArch::v7, FK_NONE, ARMAEK_DSP}
	}
}