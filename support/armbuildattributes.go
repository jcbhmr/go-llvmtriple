package support

// Legal Values for CPU_arch, (=6), uleb128
type ARMBuildAttributesCPUArch int
const (
	ARMBuildAttributesCPUArchPre_v4 ARMBuildAttributesCPUArch = 0
	ARMBuildAttributesCPUArchV4 ARMBuildAttributesCPUArch = 1           // e.g. SA110
	ARMBuildAttributesCPUArchV4T ARMBuildAttributesCPUArch = 2          // e.g. ARM7TDMI
	ARMBuildAttributesCPUArchV5T ARMBuildAttributesCPUArch = 3          // e.g. ARM9TDMI
	ARMBuildAttributesCPUArchV5TE ARMBuildAttributesCPUArch = 4         // e.g. ARM946E_S
	ARMBuildAttributesCPUArchV5TEJ ARMBuildAttributesCPUArch = 5        // e.g. ARM926EJ_S
	ARMBuildAttributesCPUArchV6 ARMBuildAttributesCPUArch = 6           // e.g. ARM1136J_S
	ARMBuildAttributesCPUArchV6KZ ARMBuildAttributesCPUArch = 7         // e.g. ARM1176JZ_S
	ARMBuildAttributesCPUArchV6T2 ARMBuildAttributesCPUArch = 8         // e.g. ARM1156T2_S
	ARMBuildAttributesCPUArchV6K ARMBuildAttributesCPUArch = 9          // e.g. ARM1176JZ_S
	ARMBuildAttributesCPUArchV7 ARMBuildAttributesCPUArch = 10          // e.g. Cortex A8, Cortex M3
	ARMBuildAttributesCPUArchV6_M ARMBuildAttributesCPUArch = 11        // e.g. Cortex M1
	ARMBuildAttributesCPUArchV6S_M ARMBuildAttributesCPUArch = 12       // v6_M with the System extensions
	ARMBuildAttributesCPUArchV7E_M ARMBuildAttributesCPUArch = 13       // v7_M with DSP extensions
	ARMBuildAttributesCPUArchV8_A ARMBuildAttributesCPUArch = 14        // v8_A AArch32
	ARMBuildAttributesCPUArchV8_R ARMBuildAttributesCPUArch = 15        // e.g. Cortex R52
	ARMBuildAttributesCPUArchV8_M_Base ARMBuildAttributesCPUArch = 16   // v8_M_Base AArch32
	ARMBuildAttributesCPUArchV8_M_Main ARMBuildAttributesCPUArch = 17   // v8_M_Main AArch32
	ARMBuildAttributesCPUArchV8_1_M_Main ARMBuildAttributesCPUArch = 21 // v8_1_M_Main AArch32
	ARMBuildAttributesCPUArchV9_A ARMBuildAttributesCPUArch = 22        // v9_A AArch32
)

