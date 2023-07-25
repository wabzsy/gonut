package gonut

// ArchType target architecture
const (
	DONUT_ARCH_ANY = -1 // just for vbs,js and xsl files
	DONUT_ARCH_X86 = 1  // x86
	DONUT_ARCH_X64 = 2  // AMD64
	DONUT_ARCH_X96 = 3  // AMD64 + x86
)

// FormatType format type
const (
	DONUT_FORMAT_BINARY     = 1
	DONUT_FORMAT_BASE64     = 2
	DONUT_FORMAT_C          = 3
	DONUT_FORMAT_RUBY       = 4
	DONUT_FORMAT_PYTHON     = 5
	DONUT_FORMAT_POWERSHELL = 6
	DONUT_FORMAT_CSHARP     = 7
	DONUT_FORMAT_HEX        = 8
	DONUT_FORMAT_UUID       = 9
)

// CompressionType compression engine
const (
	DONUT_COMPRESS_NONE   = 1
	DONUT_COMPRESS_APLIB  = 2
	DONUT_COMPRESS_LZNT1  = 3 // COMPRESSION_FORMAT_LZNT1
	DONUT_COMPRESS_XPRESS = 4 // COMPRESSION_FORMAT_XPRESS
)

// EntropyType entropy level
const (
	DONUT_ENTROPY_NONE    = 1 // don't use any entropy
	DONUT_ENTROPY_RANDOM  = 2 // use random names
	DONUT_ENTROPY_DEFAULT = 3 // use random names + symmetric encryption
)

// ExitType misc options
const (
	DONUT_OPT_EXIT_THREAD  = 1 // return to the caller which calls RtlExitUserThread
	DONUT_OPT_EXIT_PROCESS = 2 // call RtlExitUserProcess to terminate host process
	DONUT_OPT_EXIT_BLOCK   = 3 // after the main shellcode ends, do not exit or cleanup and block indefinitely
)

// BypassType AMSI/WLDP/ETW options
const (
	DONUT_BYPASS_NONE     = 1 // Disables bypassing AMSI/WDLP/ETW
	DONUT_BYPASS_ABORT    = 2 // If bypassing AMSI/WLDP/ETW fails, the loader stops running
	DONUT_BYPASS_CONTINUE = 3 // If bypassing AMSI/WLDP/ETW fails, the loader continues running
)

// HeadersType Preserve PE headers options
const (
	DONUT_HEADERS_OVERWRITE = 1 // Overwrite PE headers
	DONUT_HEADERS_KEEP      = 2 // Preserve PE headers
)
