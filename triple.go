package llvmtriple

import "github.com/jcbhmr/go-llvmtriple/internal/support"

// Triple - Helper class for working with autoconf configuration names. For
// historical reasons, we also call these 'triples' (they used to contain
// exactly three fields).
//
// Configuration names are strings in the canonical form:
//   ARCHITECTURE-VENDOR-OPERATING_SYSTEM
// or
//   ARCHITECTURE-VENDOR-OPERATING_SYSTEM-ENVIRONMENT
//
// This class is used for clients which want to support arbitrary
// configuration names, but also want to implement certain special
// behavior for particular configurations. This class isolates the mapping
// from the components of the configuration name to well known IDs.
//
// At its core the Triple class is designed to be a wrapper for a triple
// string; the constructor does not change or normalize the triple string.
// Clients that need to handle the non-canonical triples that users often
// specify should use the normalize method.
//
// See autoconf/config.guess for a glimpse into what configuration names
// look like in practice.
type Triple struct {
	data string
	// The parsed arch type.
	arch TripleArchType
	// The parsed subarchitecture type.
	subArch TripleSubArchType
	// The parsed vendor type.
	vendor TripleVendorType
	// The parsed OS type.
	os TripleOSType
	// The parsed environment type.
	environment TripleEnvironmentType
	// The object format type.
	objectFormat TripleObjectFormatType
}

type TripleArchType int

const (
	TripleUnknownArch    TripleArchType = iota
	TripleArm                           // ARM (little endian): arm, armv.*, xscale
	TripleArmeb                         // ARM (big endian): armeb
	TripleAarch64                       // AArch64 (little endian): aarch64
	TripleAarch64_be                    // AArch64 (big endian): aarch64_be
	TripleAarch64_32                    // AArch64 (little endian) ILP32: aarch64_32
	TripleArc                           // ARC: Synopsys ARC
	TripleAvr                           // AVR: Atmel AVR microcontroller
	TripleBpfel                         // eBPF or extended BPF or 64-bit BPF (little endian)
	TripleBpfeb                         // eBPF or extended BPF or 64-bit BPF (big endian)
	TripleCsky                          // CSKY: csky
	TripleDxil                          // DXIL 32-bit DirectX bytecode
	TripleHexagon                       // Hexagon: hexagon
	TripleLoongarch32                   // LoongArch (32-bit): loongarch32
	TripleLoongarch64                   // LoongArch (64-bit): loongarch64
	TripleM68k                          // M68k: Motorola 680x0 family
	TripleMips                          // MIPS: mips, mipsallegrex, mipsr6
	TripleMipsel                        // MIPSEL: mipsel, mipsallegrexe, mipsr6el
	TripleMips64                        // MIPS64: mips64, mips64r6, mipsn32, mipsn32r6
	TripleMips64el                      // MIPS64EL: mips64el, mips64r6el, mipsn32el, mipsn32r6el
	TripleMsp430                        // MSP430: msp430
	TriplePpc                           // PPC: powerpc
	TriplePpcle                         // PPCLE: powerpc (little endian)
	TriplePpc64                         // PPC64: powerpc64, ppu
	TriplePpc64le                       // PPC64LE: powerpc64le
	TripleR600                          // R600: AMD GPUs HD2XXX - HD6XXX
	TripleAmdgcn                        // AMDGCN: AMD GCN GPUs
	TripleRiscv32                       // RISC-V (32-bit): riscv32
	TripleRiscv64                       // RISC-V (64-bit): riscv64
	TripleSparc                         // Sparc: sparc
	TripleSparcv9                       // Sparcv9: Sparcv9
	TripleSparcel                       // Sparc: (endianness = little). NB: 'Sparcle' is a CPU variant
	TripleSystemz                       // SystemZ: s390x
	TripleTce                           // TCE (http://tce.cs.tut.fi/): tce
	TripleTcele                         // TCE little endian (http://tce.cs.tut.fi/): tcele
	TripleThumb                         // Thumb (little endian): thumb, thumbv.*
	TripleThumbeb                       // Thumb (big endian): thumbeb
	TripleX86                           // X86: i[3-9]86
	TripleX86_64                        // X86-64: amd64, x86_64
	TripleXcore                         // XCore: xcore
	TripleXtensa                        // Tensilica: Xtensa
	TripleNvptx                         // NVPTX: 32-bit
	TripleNvptx64                       // NVPTX: 64-bit
	TripleLe32                          // le32: generic little-endian 32-bit CPU (PNaCl)
	TripleLe64                          // le64: generic little-endian 64-bit CPU (PNaCl)
	TripleAmdil                         // AMDIL
	TripleAmdil64                       // AMDIL with 64-bit pointers
	TripleHsail                         // AMD HSAIL
	TripleHsail64                       // AMD HSAIL with 64-bit pointers
	TripleSpir                          // SPIR: standard portable IR for OpenCL 32-bit version
	TripleSpir64                        // SPIR: standard portable IR for OpenCL 64-bit version
	TripleSpirv                         // SPIR-V with logical memory layout.
	TripleSpirv32                       // SPIR-V with 32-bit pointers
	TripleSpirv64                       // SPIR-V with 64-bit pointers
	TripleKalimba                       // Kalimba: generic kalimba
	TripleShave                         // SHAVE: Movidius vector VLIW processors
	TripleLanai                         // Lanai: Lanai 32-bit
	TripleWasm32                        // WebAssembly with 32-bit pointers
	TripleWasm64                        // WebAssembly with 64-bit pointers
	TripleRenderscript32                // 32-bit RenderScript
	TripleRenderscript64                // 64-bit RenderScript
	TripleVe                            // NEC SX-Aurora Vector Engine
	TripleLastArchType   = TripleVe
)

type TripleSubArchType int

const (
	TripleNoSubArch TripleSubArchType = iota

	TripleARMSubArch_v9_5a
	TripleARMSubArch_v9_4a
	TripleARMSubArch_v9_3a
	TripleARMSubArch_v9_2a
	TripleARMSubArch_v9_1a
	TripleARMSubArch_v9
	TripleARMSubArch_v8_9a
	TripleARMSubArch_v8_8a
	TripleARMSubArch_v8_7a
	TripleARMSubArch_v8_6a
	TripleARMSubArch_v8_5a
	TripleARMSubArch_v8_4a
	TripleARMSubArch_v8_3a
	TripleARMSubArch_v8_2a
	TripleARMSubArch_v8_1a
	TripleARMSubArch_v8
	TripleARMSubArch_v8r
	TripleARMSubArch_v8m_baseline
	TripleARMSubArch_v8m_mainline
	TripleARMSubArch_v8_1m_mainline
	TripleARMSubArch_v7
	TripleARMSubArch_v7em
	TripleARMSubArch_v7m
	TripleARMSubArch_v7s
	TripleARMSubArch_v7k
	TripleARMSubArch_v7ve
	TripleARMSubArch_v6
	TripleARMSubArch_v6m
	TripleARMSubArch_v6k
	TripleARMSubArch_v6t2
	TripleARMSubArch_v5
	TripleARMSubArch_v5te
	TripleARMSubArch_v4t

	TripleAArch64SubArch_arm64e
	TripleAArch64SubArch_arm64ec

	TripleKalimbaSubArch_v3
	TripleKalimbaSubArch_v4
	TripleKalimbaSubArch_v5

	TripleMipsSubArch_r6

	TriplePPCSubArch_spe

	// SPIR-V sub-arch corresponds to its version.
	TripleSPIRVSubArch_v10
	TripleSPIRVSubArch_v11
	TripleSPIRVSubArch_v12
	TripleSPIRVSubArch_v13
	TripleSPIRVSubArch_v14
	TripleSPIRVSubArch_v15
	TripleSPIRVSubArch_v16

	// DXIL sub-arch corresponds to its version.
	TripleDXILSubArch_v1_0
	TripleDXILSubArch_v1_1
	TripleDXILSubArch_v1_2
	TripleDXILSubArch_v1_3
	TripleDXILSubArch_v1_4
	TripleDXILSubArch_v1_5
	TripleDXILSubArch_v1_6
	TripleDXILSubArch_v1_7
	TripleDXILSubArch_v1_8
	TripleLatestDXILSubArch = TripleDXILSubArch_v1_8
)

type TripleVendorType int

const (
	TripleUnknownVendor TripleVendorType = iota

	TripleApple
	TriplePC
	TripleSCEI
	TripleFreescale
	TripleIBM
	TripleImaginationTechnologies
	TripleMipsTechnologies
	TripleNVIDIA
	TripleCSR
	TripleAMD
	TripleMesa
	TripleSUSE
	TripleOpenEmbedded
	TripleLastVendorType = TripleOpenEmbedded
)

type TripleOSType int

const (
	TripleUnknownOS TripleOSType = iota

	TripleDarwin
	TripleDragonFly
	TripleFreeBSD
	TripleFuchsia
	TripleIOS
	TripleKFreeBSD
	TripleLinux
	TripleLv2 // PS3
	TripleMacOSX
	TripleNetBSD
	TripleOpenBSD
	TripleSolaris
	TripleUEFI
	TripleWin32
	TripleZOS
	TripleHaiku
	TripleRTEMS
	TripleNaCl // Native Client
	TripleAIX
	TripleCUDA   // NVIDIA CUDA
	TripleNVCL   // NVIDIA OpenCL
	TripleAMDHSA // AMD HSA Runtime
	TriplePS4
	TriplePS5
	TripleELFIAMCU
	TripleTvOS      // Apple tvOS
	TripleWatchOS   // Apple watchOS
	TripleBridgeOS  // Apple bridgeOS
	TripleDriverKit // Apple DriverKit
	TripleXROS      // Apple XROS
	TripleMesa3D
	TripleAMDPAL     // AMD PAL Runtime
	TripleHermitCore // HermitCore Unikernel/Multikernel
	TripleHurd       // GNU/Hurd
	TripleWASI       // Experimental WebAssembly OS
	TripleEmscripten
	TripleShaderModel // DirectX ShaderModel
	TripleLiteOS
	TripleSerenity
	TripleVulkan     // Vulkan SPIR-V
	TripleLastOSType = TripleVulkan
)

type TripleEnvironmentType int

const (
	TripleUnknownEnvironment TripleEnvironmentType = iota

	TripleGNU
	TripleGNUABIN32
	TripleGNUABI64
	TripleGNUEABI
	TripleGNUEABIHF
	TripleGNUF32
	TripleGNUF64
	TripleGNUSF
	TripleGNUX32
	TripleGNUILP32
	TripleCODE16
	TripleEABI
	TripleEABIHF
	TripleAndroid
	TripleMusl
	TripleMuslEABI
	TripleMuslEABIHF
	TripleMuslX32

	TripleMSVC
	TripleItanium
	TripleCygnus
	TripleCoreCLR
	TripleSimulator // Simulator variants of other systems, e.g., Apple's iOS
	TripleMacABI    // Mac Catalyst variant of Apple's iOS deployment target.

	// Shader Stages
	// The order of these values matters, and must be kept in sync with the
	// language options enum in Clang. The ordering is enforced in
	// static_asserts in Triple.cpp and in Clang.
	TriplePixel
	TripleVertex
	TripleGeometry
	TripleHull
	TripleDomain
	TripleCompute
	TripleLibrary
	TripleRayGeneration
	TripleIntersection
	TripleAnyHit
	TripleClosestHit
	TripleMiss
	TripleCallable
	TripleMesh
	TripleAmplification
	TripleOpenCL
	TripleOpenHOS

	TriplePAuthTest

	TripleGNUT64
	TripleGNUEABIT64
	TripleGNUEABIHFT64

	TripleLastEnvironmentType = TripleGNUEABIHFT64
)

type TripleObjectFormatType int

const (
	TripleUnknownObjectFormat TripleObjectFormatType = iota

	TripleCOFF
	TripleDXContainer
	TripleELF
	TripleGOFF
	TripleMachO
	TripleSPIRV
	TripleWasm
	TripleXCOFF
)

// Default constructor is the same as an empty string and leaves all
// triple fields unknown.
func NewTriple() *Triple {
	panic("not implemented")
}

func NewTriple2(str string) *Triple {
	panic("not implemented")
}

func NewTriple3(archStr string, vendorStr string, osStr string) *Triple {
	panic("not implemented")
}

func NewTriple4(archStr string, vendorStr string, osStr string, environmentStr string) *Triple {
	panic("not implemented")
}

func (t *Triple) Equal(other *Triple) bool {
	return t.arch == other.arch && t.subArch == other.subArch && t.vendor == other.vendor && t.os == other.os && t.environment == other.environment && t.objectFormat == other.objectFormat
}

// Turn an arbitrary machine specification into the canonical triple form (or
// something sensible that the Triple class understands if nothing better can
// reasonably be done).  In particular, it handles the common case in which
// otherwise valid components are in the wrong order.
func TripleNormalize(str string) string {
	panic("not implemented")
}

// Return the normalized form of this triple's string.
func (t *Triple) Normalize() string {
	return TripleNormalize(t.data)
}

// Get the parsed architecture type of this triple.
func (t *Triple) Arch() TripleArchType {
	return t.arch
}

// get the parsed subarchitecture type for this triple.
func (t *Triple) SubArch() TripleSubArchType {
	return t.subArch
}

// Get the parsed vendor type of this triple.
func (t *Triple) Vendor() TripleVendorType {
	return t.vendor
}

// Does this triple have the optional environment (fourth) component?
func (t *Triple) HasEnvironment() bool {
	return t.EnvironmentName() != ""
}

// Get the parsed environment type of this triple.
func (t *Triple) Environment() TripleEnvironmentType {
	return t.environment
}

// Parse the version number from the OS name component of the
// triple, if present.
//
// For example, "fooos1.2.3" would return (1, 2, 3).
func (t *Triple) EnvironmentVersion() support.VersionTuple {
	panic("not implemented")
}

// Get the object format for this triple.
func (t *Triple) ObjectFormat() TripleObjectFormatType {
	return t.objectFormat
}

// Parse the version number from the OS name component of the triple, if
// present.
//
// For example, "fooos1.2.3" would return (1, 2, 3).
func (t *Triple) OSVersion() support.VersionTuple {
	panic("not implemented")
}

// Return just the major version number, this is specialized because it is a
// common query.
func (t *Triple) OSMajorVersion() uint32 {
	panic("not implemented")
}

// Parse the version number as with getOSVersion and then translate generic
// "darwin" versions to the corresponding OS X versions.  This may also be
// called with IOS triples but the OS X version number is just set to a
// constant 10.4.0 in that case.  Returns true if successful.
func (t *Triple) MacOSXVersion() (support.VersionTuple, bool) {
	panic("not implemented")
}

// Parse the version number as with getOSVersion.  This should only be called
// with IOS or generic triples.
func (t *Triple) IOSVersion() support.VersionTuple {
	panic("not implemented")
}

// Parse the version number as with getOSVersion.  This should only be called
// with WatchOS or generic triples.
func (t *Triple) WatchOSVersion() support.VersionTuple {
	panic("not implemented")
}

// Parse the version number as with getOSVersion.
func (t *Triple) DriverKitVersion() support.VersionTuple {
	panic("not implemented")
}

// Parse the Vulkan version number from the OSVersion and SPIR-V version
// (SubArch).  This should only be called with Vulkan SPIR-V triples.
func (t *Triple) VulkanVersion() support.VersionTuple {
	panic("not implemented")
}

// Parse the DXIL version number from the OSVersion and DXIL version
// (SubArch).  This should only be called with DXIL triples.
func (t *Triple) DXILVersion() support.VersionTuple {
	panic("not implemented")
}

func (t *Triple) String() string {
	return t.data
}

func (t *Triple) Triple() string {
	return t.data
}

// Get the architecture (first) component of the triple.
func (t *Triple) ArchName() string {
	panic("not implemented")
}

// Get the vendor (second) component of the triple.
func (t *Triple) VendorName() string {
	panic("not implemented")
}

// Get the operating system (third) component of the triple.
func (t *Triple) OSName() string {
	panic("not implemented")
}

// Get the optional environment (fourth) component of the triple, or "" if
// empty.
func (t *Triple) EnvironmentName() string {
	panic("not implemented")
}

// Get the operating system and optional environment components as a single
// string (separated by a '-' if the environment component is present).
func (t *Triple) OSAndEnvironmentName() string {
	panic("not implemented")
}

// Get the version component of the environment component as a single
// string (the version after the environment).
//
// For example, "fooos1.2.3" would return "1.2.3".
func (t *Triple) EnvironmentVersionNString() string {
	panic("not implemented")
}

// Returns the pointer width of this architecture.
func TripleArchPointerBitWidth(arch TripleArchType) uint {
	panic("not implemented")
}

// Returns the pointer width of this architecture.
func (t *Triple) ArchPointerBitWidth() uint {
	return TripleArchPointerBitWidth(t.arch)
}

// Test whether the architecture is 64-bit
//
// Note that this tests for 64-bit pointer width, and nothing else. Note
// that we intentionally expose only three predicates, 64-bit, 32-bit, and
// 16-bit. The inner details of pointer width for particular architectures
// is not summed up in the triple, and so only a coarse grained predicate
// system is provided.
func (t *Triple) IsArch64Bit() bool {
	panic("not implemented")
}

// Test whether the architecture is 32-bit
//
// Note that this tests for 32-bit pointer width, and nothing else.
func (t *Triple) IsArch32Bit() bool {
	panic("not implemented")
}

// Test whether the architecture is 16-bit
//
// Note that this tests for 16-bit pointer width, and nothing else.
func (t *Triple) IsArch16Bit() bool {
	panic("not implemented")
}

// Helper function for doing comparisons against version numbers included in
// the target triple.
func (t *Triple) IsOSVersionLT(major uint, minor *uint, micro *uint) bool {
	var minor2 uint
	if minor != nil {
		minor2 = *minor
	}
	var micro2 uint
	if micro != nil {
		micro2 = *micro
	}
	if minor2 == 0 {
		return t.OSVersion().Cmp(support.NewVersionTuple2(major)) < 0
	}
	if micro2 == 0 {
		return t.OSVersion().Cmp(support.NewVersionTuple3(major, minor2)) < 0
	}
	return t.OSVersion().Cmp(support.NewVersionTuple4(major, minor2, micro2)) < 0
}

func (t *Triple) IsOSVersionLT2(other *Triple) bool {
	return t.OSVersion().Cmp(other.OSVersion()) < 0
}

// Comparison function for checking OS X version compatibility, which handles
// supporting skewed version numbering schemes used by the "darwin" triples.
func (t *Triple) IsMacOSXVersionLT(major uint, minor *uint, micro *uint) bool {
	panic("not implemented")
}

// Is this a Mac OS X triple. For legacy reasons, we support both "darwin"
// and "osx" as OS X triples.
func (t *Triple) IsMacOSX() bool {
	return t.os == TripleDarwin || t.os == TripleMacOSX
}

// Is this an iOS triple.
// Note: This identifies tvOS as a variant of iOS. If that ever
// changes, i.e., if the two operating systems diverge or their version
// numbers get out of sync, that will need to be changed.
// watchOS has completely different version numbers so it is not included.
func (t *Triple) IsiOS() bool {
	return t.os == TripleIOS || t.os == TripleTvOS
}

// Is this an Apple tvOS triple.
func (t *Triple) IsTvOS() bool {
	return t.os == TripleTvOS
}

// Is this an Apple watchOS triple.
func (t *Triple) IsWatchOS() bool {
	return t.os == TripleWatchOS
}

func (t *Triple) IsWatchABI() bool {
	return t.subArch == TripleARMSubArch_v7k
}

// Is this an Apple XROS triple.
func (t *Triple) IsXROS() bool {
	return t.os == TripleXROS
}

// Is this an Apple DriverKit triple.
func (t *Triple) IsDriverKit() bool {
	return t.os == TripleDriverKit
}

func (t *Triple) IsOSzOS() bool {
	return t.os == TripleZOS
}

// Is this a "Darwin" OS (macOS, iOS, tvOS, watchOS, XROS, or DriverKit).
func (t *Triple) IsOSDarwin() bool {
	return t.IsMacOSX() || t.IsiOS() || t.IsWatchOS() || t.IsDriverKit() || t.IsXROS()
}

func (t *Triple) IsSimulatorEnvironment() bool {
	return t.environment == TripleSimulator
}

func (t *Triple) IsMacCatalystEnvironment() bool {
	return t.environment == TripleMacABI
}

// Returns true for targets that run on a macOS machine.
func (t *Triple) IsTargetMachineMac() bool {
	return t.IsMacOSX() || (t.IsOSDarwin() && (t.IsSimulatorEnvironment() || t.IsMacCatalystEnvironment()))
}

func (t *Triple) IsOSNetBSD() bool {
	return t.os == TripleNetBSD
}

func (t *Triple) IsOSOpenBSD() bool {
	return t.os == TripleOpenBSD
}

func (t *Triple) IsOSFreeBSD() bool {
	return t.os == TripleFreeBSD
}

func (t *Triple) IsOSFuchsia() bool {
	return t.os == TripleFuchsia
}

func (t *Triple) IsOSDragonFly() bool {
	return t.os == TripleDragonFly
}

func (t *Triple) IsOSSolaris() bool {
	return t.os == TripleSolaris
}

func (t *Triple) IsOSIAMCU() bool {
	return t.os == TripleELFIAMCU
}

func (t *Triple) IsOSUnknown() bool {
	return t.os == TripleUnknownOS
}

func (t *Triple) IsGNUEnvironment() bool {
	env := t.environment
	return env == TripleGNU || env == TripleGNUT64 ||
		env == TripleGNUABIN32 || env == TripleGNUABI64 ||
		env == TripleGNUEABI || env == TripleGNUEABIT64 ||
		env == TripleGNUEABIHF || env == TripleGNUEABIHFT64 ||
		env == TripleGNUF32 || env == TripleGNUF64 ||
		env == TripleGNUSF || env == TripleGNUX32
}

// Tests whether the OS is Haiku.
func (t *Triple) IsOSHaiku() bool {
	return t.os == TripleHaiku
}

// Tests whether the OS is UEFI.
func (t *Triple) IsUEFI() bool {
	return t.os == TripleUEFI
}

// Tests whether the OS is Windows.
func (t *Triple) IsOSWindows() bool {
	return t.os == TripleWin32
}

// Checks if the environment is MSVC.
func (t *Triple) IsKnownWindowsMSVCEnvironment() bool {
	return t.IsOSWindows() && t.environment == TripleMSVC
}

// Checks if the environment could be MSVC.
func (t *Triple) IsWindowsMSVCEnvironment() bool {
	return t.IsKnownWindowsMSVCEnvironment() || (t.IsOSWindows() && t.environment == TripleUnknownEnvironment)
}

// Checks if we're using the Windows Arm64EC ABI.
func (t *Triple) IsWindowsArm64EC() bool {
	return t.arch == TripleAarch64 && t.subArch == TripleAArch64SubArch_arm64ec
}

func (t *Triple) IsWindowsCoreCLREnvironment() bool {
	return t.IsOSWindows() && t.environment == TripleCoreCLR
}

func (t *Triple) IsWindowsItaniumEnvironment() bool {
	return t.IsOSWindows() && t.environment == TripleItanium
}

func (t *Triple) IsWindowsCygwinEnvironment() bool {
	return t.IsOSWindows() && t.environment == TripleCygnus
}

func (t *Triple) IsWindowsGNUEnvironment() bool {
	return t.IsOSWindows() && t.environment == TripleGNU
}

// Tests for either Cygwin or MinGW OS
func (t *Triple) IsOSCygMing() bool {
	return t.IsWindowsCygwinEnvironment() || t.IsWindowsGNUEnvironment()
}

// Is this a "Windows" OS targeting a "MSVCRT.dll" environment.
func (t *Triple) IsOSmSVCRT() bool {
	return t.IsWindowsMSVCEnvironment() || t.IsWindowsGNUEnvironment() || t.IsWindowsItaniumEnvironment()
}

// Tests whether the OS is NaCl (Native Client)
func (t *Triple) IsOSNaCl() bool {
	return t.os == TripleNaCl
}

// Tests whether the OS is Linux.
func (t *Triple) IsOSLinux() bool {
	return t.os == TripleLinux
}

// Tests whether the OS is kFreeBSD.
func (t *Triple) IsOSKFreeBSD() bool {
	return t.os == TripleKFreeBSD
}

// Tests whether the OS is Hurd.
func (t *Triple) IsOSHurd() bool {
	return t.os == TripleHurd
}

// Tests whether the OS is WASI.
func (t *Triple) IsOSWASI() bool {
	return t.os == TripleWASI
}

// Tests whether the OS is Emscripten.
func (t *Triple) IsOSEmscripten() bool {
	return t.os == TripleEmscripten
}

// Tests whether the OS uses glibc.
func (t *Triple) IsOSGlibc() bool {
	return (t.os == TripleLinux || t.os == TripleKFreeBSD || t.os == TripleHurd) && !t.IsAndroid()
}

// Tests whether the OS is AIX.
func (t *Triple) IsOSAIX() bool {
	return t.os == TripleAIX
}

func (t *Triple) IsOSSerenity() bool {
	return t.os == TripleSerenity
}

// Tests whether the OS uses the ELF binary format.
func (t *Triple) IsOSBinFormatELF() bool {
	return t.objectFormat == TripleELF
}

// Tests whether the OS uses the COFF binary format.
func (t *Triple) IsOSBinFormatCOFF() bool {
	return t.objectFormat == TripleCOFF
}

// Tests whether the OS uses the GOFF binary format.
func (t *Triple) IsOSBinFormatGOFF() bool {
	return t.objectFormat == TripleGOFF
}

// Tests whether the environment is MachO.
func (t *Triple) IsOSBinFormatMacho() bool {
	return t.objectFormat == TripleMachO
}

// Tests whether the OS uses the Wasm binary format.
func (t *Triple) IsOSBinFormatWasm() bool {
	return t.objectFormat == TripleWasm
}

// Tests whether the OS uses the XCOFF binary format.
func (t *Triple) IsOSBinFormatXCOFF() bool {
	return t.objectFormat == TripleXCOFF
}

// Tests whether the OS uses the DXContainer binary format.
func (t *Triple) IsOSBinFormatDXContainer() bool {
	return t.objectFormat == TripleDXContainer
}

// Tests whether the target is the PS4 platform.
func (t *Triple) IsPS4() bool {
	return t.arch == TripleX86_64 && t.vendor == TripleSCEI && t.os == TriplePS4
}

// Tests whether the target is the PS5 platform.
func (t *Triple) IsPS5() bool {
	return t.arch == TripleX86_64 && t.vendor == TripleSCEI && t.os == TriplePS5
}

// Tests whether the target is the PS4 or PS5 platform.
func (t *Triple) IsPS() bool {
	return t.IsPS4() || t.IsPS5()
}

// Tests whether the target is Android
func (t *Triple) IsAndroid() bool {
	return t.environment == TripleAndroid
}

func (t *Triple) IsAndroidVersionLT(major uint) bool {
	if !t.IsAndroid() {
		panic("not an Android triple")
	}

	version := t.EnvironmentVersion()

	// 64-bit targets did not exist before API level 21 (Lollipop).
	if t.IsArch64Bit() && version.Major() < 21 {
		return support.NewVersionTuple2(21).Cmp(support.NewVersionTuple2(major)) < 0
	}

	return version.Cmp(support.NewVersionTuple2(major)) < 0
}

// Tests whether the environment is musl-libc
func (t *Triple) IsMusl() bool {
	return t.environment == TripleMusl ||
		t.environment == TripleMuslEABI ||
		t.environment == TripleMuslEABIHF ||
		t.environment == TripleMuslX32 ||
		t.environment == TripleOpenHOS || t.IsOSLiteOS()
}

// Tests whether the target is OHOS
// LiteOS default enviroment is also OHOS, but omited on triple.
func (t *Triple) IsOHOSFamily() bool {
	return t.IsOpenHOS() || t.IsOSLiteOS()
}

func (t *Triple) IsOpenHOS() bool {
	return t.environment == TripleOpenHOS
}

func (t *Triple) IsOSLiteOS() bool {
	return t.os == TripleLiteOS
}

// Tests whether the target is DXIL.
func (t *Triple) IsDXIL() bool {
	return t.arch == TripleDxil
}

func (t *Triple) IsShaderModelOS() bool {
	return t.os == TripleShaderModel
}

func (t *Triple) IsVulkanOS() bool {
	return t.os == TripleVulkan
}

func (t *Triple) IsShaderStageEnvironment() bool {
	env := t.environment
	return env == TriplePixel || env == TripleVertex ||
		env == TripleGeometry || env == TripleHull ||
		env == TripleDomain || env == TripleCompute ||
		env == TripleLibrary || env == TripleRayGeneration ||
		env == TripleIntersection || env == TripleAnyHit ||
		env == TripleClosestHit || env == TripleMiss ||
		env == TripleCallable || env == TripleMesh ||
		env == TripleAmplification
}

// Tests whether the target is SPIR (32- or 64-bit).
func (t *Triple) IsSPIR() bool {
	return t.arch == TripleSpir || t.arch == TripleSpir64
}

// Tests whether the target is SPIR-V (32/64-bit/Logical).
func (t *Triple) IsSPIRV() bool {
	return t.arch == TripleSpirv32 || t.arch == TripleSpirv64 || t.arch == TripleSpirv
}

// Tests whether the target is SPIR-V Logical
func (t *Triple) IsSPIRVLogical() bool {
	return t.arch == TripleSpirv
}

// Tests whether the target is NVPTX (32- or 64-bit).
func (t *Triple) IsNVPTX() bool {
	return t.arch == TripleNvptx || t.arch == TripleNvptx64
}

// Tests whether the target is AMDGCN
func (t *Triple) IsAMDGCN() bool {
	return t.arch == TripleAmdgcn
}

func (t *Triple) IsAMDGPU() bool {
	return t.arch == TripleR600 || t.arch == TripleAmdgcn
}

// Tests whether the target is Thumb (little and big endian).
func (t *Triple) IsThumb() bool {
	return t.arch == TripleThumb || t.arch == TripleThumbeb
}

// Tests whether the target is ARM (little and big endian).
func (t *Triple) IsARM() bool {
	return t.arch == TripleArm || t.arch == TripleArmeb
}

// Tests whether the target supports the EHABI exception
// handling standard.
func (t *Triple) IsARM_EABI() bool {
	return (t.IsARM() || t.IsThumb()) &&
		(t.environment == TripleEABI ||
			t.environment == TripleGNUEABI ||
			t.environment == TripleGNUEABIT64 ||
			t.environment == TripleMuslEABI ||
			t.environment == TripleEABIHF ||
			t.environment == TripleGNUEABIHF ||
			t.environment == TripleGNUEABIHFT64 ||
			t.environment == TripleOpenHOS ||
			t.environment == TripleMuslEABIHF || t.IsAndroid()) &&
		t.IsOSBinFormatELF()
}

// Tests whether the target is T32.
func (t *Triple) IsArmT32() bool {
	switch t.subArch {
	case TripleARMSubArch_v8m_baseline,
		TripleARMSubArch_v7s,
		TripleARMSubArch_v7k,
		TripleARMSubArch_v7ve,
		TripleARMSubArch_v6,
		TripleARMSubArch_v6m,
		TripleARMSubArch_v6k,
		TripleARMSubArch_v6t2,
		TripleARMSubArch_v5,
		TripleARMSubArch_v5te,
		TripleARMSubArch_v4t:
		return false
	default:
		return true
	}
}

// Tests whether the target is an M-class.
func (t *Triple) IsArmMClass() bool {
	switch t.subArch {
	case TripleARMSubArch_v6m,
		TripleARMSubArch_v7m,
		TripleARMSubArch_v7em,
		TripleARMSubArch_v8m_mainline,
		TripleARMSubArch_v8m_baseline,
		TripleARMSubArch_v8_1m_mainline:
		return true
	default:
		return false
	}
}

// Tests whether the target is AArch64 (little and big endian).
func (t *Triple) IsAArch64() bool {
	return t.arch == TripleAarch64 || t.arch == TripleAarch64_be
}

// Tests whether the target is AArch64 and pointers are the size specified by
// pointerWidth.
func (t *Triple) IsAArch642(pointerWidth int) bool {
	if !(pointerWidth == 32 || pointerWidth == 64) {
		panic("invalid pointer width")
	}
	if !t.IsAArch64() {
		return false
	}
	if t.arch == TripleAarch64_32 || t.environment == TripleGNUILP32 {
		return pointerWidth == 32
	} else {
		return pointerWidth == 64
	}
}

// Tests whether the target is 32-bit LoongArch.
func (t *Triple) IsLoongArch32() bool {
	return t.arch == TripleLoongarch32
}

// Tests whether the target is 64-bit LoongArch.
func (t *Triple) IsLoongArch64() bool {
	return t.arch == TripleLoongarch64
}

// Tests whether the target is LoongArch (32- and 64-bit).
func (t *Triple) IsLoongArch() bool {
	return t.IsLoongArch32() || t.IsLoongArch64()
}

// Tests whether the target is MIPS 32-bit (little and big endian).
func (t *Triple) IsMIPS32() bool {
	return t.arch == TripleMips || t.arch == TripleMipsel
}

// Tests whether the target is MIPS 64-bit (little and big endian).
func (t *Triple) IsMIPS64() bool {
	return t.arch == TripleMips64 || t.arch == TripleMips64el
}

// Tests whether the target is MIPS (little and big endian, 32- or 64-bit).
func (t *Triple) IsMIPS() bool {
	return t.IsMIPS32() || t.IsMIPS64()
}

// Tests whether the target is PowerPC (32- or 64-bit LE or BE).
func (t *Triple) IsPPC() bool {
	return t.arch == TriplePpc ||
		t.arch == TriplePpc64 ||
		t.arch == TriplePpcle ||
		t.arch == TriplePpc64le
}

// Tests whether the target is 32-bit PowerPC (little and big endian).
func (t *Triple) IsPPC32() bool {
	return t.arch == TriplePpc || t.arch == TriplePpcle
}

// Tests whether the target is 64-bit PowerPC (little and big endian).
func (t *Triple) IsPPC64() bool {
	return t.arch == TriplePpc64 || t.arch == TriplePpc64le
}

// Tests whether the target 64-bit PowerPC big endian ABI is ELFv2.
func (t *Triple) IsPPC64ELFv2ABI() bool {
	return (t.arch == TriplePpc64 &&
		((t.os == TripleFreeBSD &&
			(t.OSMajorVersion() >= 13 || t.OSVersion().Empty())) ||
			t.os == TripleOpenBSD || t.IsMusl()))
}

// Tests whether the target 32-bit PowerPC uses Secure PLT.
func (t *Triple) IsPPC32SecurePlt() bool {
	return ((t.arch == TriplePpc || t.arch == TriplePpcle) &&
		((t.os == TripleFreeBSD &&
			(t.OSMajorVersion() >= 13 || t.OSVersion().Empty())) ||
			t.os == TripleNetBSD || t.os == TripleOpenBSD || t.IsMusl()))
}

// Tests whether the target is 32-bit RISC-V.
func (t *Triple) IsRISCV32() bool {
	return t.arch == TripleRiscv32
}

// Tests whether the target is 64-bit RISC-V.
func (t *Triple) IsRISCV64() bool {
	return t.arch == TripleRiscv64
}

// Tests whether the target is RISC-V (32- and 64-bit).
func (t *Triple) IsRISCV() bool {
	return t.IsRISCV32() || t.IsRISCV64()
}

// Tests whether the target is 32-bit SPARC (little and big endian).
func (t *Triple) IsSparc32() bool {
	return t.arch == TripleSparc || t.arch == TripleSparcel
}

// Tests whether the target is 64-bit SPARC (big endian).
func (t *Triple) IsSparc64() bool {
	return t.arch == TripleSparcv9
}

// Tests whether the target is SPARC.
func (t *Triple) IsSparc() bool {
	return t.IsSparc32() || t.IsSparc64()
}

// Tests whether the target is SystemZ.
func (t *Triple) IsSystemZ() bool {
	return t.arch == TripleSystemz
}

// Tests whether the target is x86 (32- or 64-bit).
func (t *Triple) IsX86() bool {
	return t.arch == TripleX86 || t.arch == TripleX86_64
}

// Tests whether the target is VE
func (t *Triple) IsVE() bool {
	return t.arch == TripleVe
}

// Tests whether the target is wasm (32- and 64-bit).
func (t *Triple) IsWasm() bool {
	return t.arch == TripleWasm32 || t.arch == TripleWasm64
}

// Tests whether the target is CSKY
func (t *Triple) IsCSKY() bool {
	return t.arch == TripleCsky
}

// Tests whether the target is the Apple "arm64e" AArch64 subarch.
func (t *Triple) IsAArch64Arme() bool {
	return t.arch == TripleAarch64 && t.subArch == TripleAArch64SubArch_arm64e
}

// Tests whether the target is X32.
func (t *Triple) IsX32() bool {
	return t.environment == TripleGNUX32 || t.environment == TripleMuslX32
}

// Tests whether the target is eBPF.
func (t *Triple) IsBPF() bool {
	return t.arch == TripleBpfel || t.arch == TripleBpfeb
}

// Tests if the target forces 64-bit time_t on a 32-bit architecture.
func (t *Triple) IsTime64ABI() bool {
	env := t.environment
	return env == TripleGNUT64 || env == TripleGNUEABIT64 ||
		env == TripleGNUEABIHFT64
}

// Tests if the target forces hardfloat.
func (t *Triple) IsHardFloatABI() bool {
	env := t.environment
	return env == TripleGNUEABIHF ||
		env == TripleGNUEABIHFT64 ||
		env == TripleMuslEABIHF ||
		env == TripleEABIHF
}

// Tests whether the target supports comdat
func (t *Triple) SupportsCOMDAT() bool {
	return !(t.IsOSBinFormatMacho() || t.IsOSBinFormatXCOFF() || t.IsOSBinFormatDXContainer())
}

// Tests whether the target uses emulated TLS as default.
//
// Note: Android API level 29 (10) introduced ELF TLS.
func (t *Triple) HasDefaultEmulatedTLS() bool {
	return (t.IsAndroid() && t.IsAndroidVersionLT(29)) || t.IsOSOpenBSD() || t.IsWindowsCygwinEnvironment() || t.IsOHOSFamily()
}

// True if the target supports both general-dynamic and TLSDESC, and TLSDESC
// is enabled by default.
func (t *Triple) HasDefaultTLSDESC() bool {
	return t.IsAndroid() && t.IsRISCV64()
}

// Tests whether the target uses -data-sections as default.
func (t *Triple) HasDefaultDataSections() bool {
	return t.IsOSBinFormatXCOFF() || t.IsWasm()
}

// Tests if the environment supports dllimport/export annotations.
func (t *Triple) HasDLLImportExport() bool {
	return t.IsOSWindows() || t.IsPS()
}

// Set the architecture (first) component of the triple to a known type.
func (t *Triple) SetArch(kind TripleArchType, subArch *TripleSubArchType) {
	var subArch2 TripleSubArchType
	if subArch != nil {
		subArch2 = *subArch
	} else {
		subArch2 = TripleNoSubArch
	}
	_ = subArch2
	panic("not implemented")
}

// Set the vendor (second) component of the triple to a known type.
func (t *Triple) SetVendor(kind TripleVendorType) {
	panic("not implemented")
}

// Set the operating system (third) component of the triple to a known type.
func (t *Triple) SetOS(kind TripleOSType) {
	panic("not implemented")
}

// Set the environment (fourth) component of the triple to a known type.
func (t *Triple) SetEnvironment(kind TripleEnvironmentType) {
	panic("not implemented")
}

// Set the object file format.
func (t *Triple) SetObjectFormat(kind TripleObjectFormatType) {
	panic("not implemented")
}

// Set all components to the new triple str.
func (t *Triple) SetTriple(str string) {
	panic("not implemented")
}

// Set the architecture (first) component of the triple by name.
func (t *Triple) SetArchName(str string) {
	panic("not implemented")
}

// Set the vendor (second) component of the triple by name.
func (t *Triple) SetVendorName(str string) {
	panic("not implemented")
}

// Set the operating system (third) component of the triple by name.
func (t *Triple) SetOSName(str string) {
	panic("not implemented")
}

// Set the optional environment (fourth) component of the triple by name.
func (t *Triple) SetEnvironmentName(str string) {
	panic("not implemented")
}

// Set the operating system and optional environment components with a single
// string.
func (t *Triple) SetOSAndEnvironmentName(str string) {
	panic("not implemented")
}

// Form a triple with a 32-bit variant of the current architecture.
//
// This can be used to move across "families" of architectures where useful.
//
// Returns: A new triple with a 32-bit architecture or an unknown
//          architecture if no such variant can be found.
func (t *Triple) X32BitArchVariant() *Triple {
	panic("not implemented")
}

// Form a triple with a 64-bit variant of the current architecture.
//
// This can be used to move across "families" of architectures where useful.
//
// Returns: A new triple with a 64-bit architecture or an unknown
//          architecture if no such variant can be found.
func (t *Triple) X64BitArchVariant() *Triple {
	panic("not implemented")
}

// Form a triple with a big endian variant of the current architecture.
//
// This can be used to move across "families" of architectures where useful.
//
// Returns: A new triple with a big endian architecture or an unknown
//          architecture if no such variant can be found.
func (t *Triple) BigEndianArchVariant() *Triple {
	panic("not implemented")
}

// Form a triple with a little endian variant of the current architecture.
//
// This can be used to move across "families" of architectures where useful.
//
// Returns: A new triple with a little endian architecture or an unknown
//          architecture if no such variant can be found.
func (t *Triple) LittleEndianArchVariant() *Triple {
	panic("not implemented")
}

// Tests whether the target triple is little endian.
//
// Returns: true if the triple is little endian, false otherwise.
func (t *Triple) IsLittleEndian() bool {
	panic("not implemented")
}

// Test whether target triples are compatible.
func (t *Triple) IsCompatibleWith(other *Triple) bool {
	panic("not implemented")
}

// Merge target triples.
func (t *Triple) Merge(other *Triple) string {
	panic("not implemented")
}

// Some platforms have different minimum supported OS versions that
// varies by the architecture specified in the triple. This function
// returns the minimum supported OS version for this triple if one an exists,
// or an invalid version tuple if this triple doesn't have one.
func (t *Triple) MinimumSupportedOSVersion() support.VersionTuple {
	panic("not implemented")
}

// Get the canonical name for the kind architecture.
func TripleArchTypeName(kind TripleArchType) string {
	panic("not implemented")
}

// Get the architecture name based on kind and subArch.
func TripleArchName(kind TripleArchType, subArch *TripleSubArchType) string {
	var subArch2 TripleSubArchType
	if subArch != nil {
		subArch2 = *subArch
	} else {
		subArch2 = TripleNoSubArch
	}
	_ = subArch2
	panic("not implemented")
}

// Get the "prefix" canonical name for the kind architecture. This is the
// prefix used by the architecture specific builtins, and is suitable for
// passing to Intrinsic::getIntrinsicForClangBuiltin().
//
// return - The architecture prefix, or 0 if none is defined.
func TripleArchTypePrefix(kind TripleArchType) string {
	panic("not implemented")
}

// Get the canonical name for the kind vendor.
func TripleVendorTypeName(kind TripleVendorType) string {
	panic("not implemented")
}

// Get the canonical name for the kind operating system.
func TripleOSTypeName(kind TripleOSType) string {
	panic("not implemented")
}

// Get the canonical name for the kind environment.
func TripleEnvironmentTypeName(kind TripleEnvironmentType) string {
	panic("not implemented")
}

// Get the name for the object format.
func TripleObjectFormatTypeName(kind TripleObjectFormatType) string {
	panic("not implemented")
}

// The canonical type for the given LLVM architecture name (e.g., "x86").
func TripleArchTypeForLLVMName(str string) TripleArchType {
	panic("not implemented")
}

// Returns a canonicalized OS version number for the specified OS.
func TripleCanonicalVersionForOS(os TripleOSType, version support.VersionTuple) support.VersionTuple {
	panic("not implemented")
}