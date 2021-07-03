package errs

const (
	UnknowVersion = iota
	Marshal
	Unmarshal
	Encrypt
	Decrypt
	WriteFile
	ReadFile
)

const (
	UnknowVersionDesc = "未知版本信息"
	MarshalDesc       = "json marshal error"
	UnmarshalDesc     = "json unmarshal error"
	EncryptDesc       = "encrypt error"
	DecryptDesc       = "decrypt error"
	WriteFileDesc     = "write file error"
	ReadFileDesc      = "read file error"
)

// space相关异常
const (
	UniqueSpaceName              = 11001
	UniqueSpaceNameDesc          = "重复的空间名称"
	GenerateSpaceStorageData     = 11002
	GenerateSpaceStorageDataDesc = "空间数据转存储内容失败"
)

// record相关异常
const (
	AddRecordBodyError     = 21001
	AddRecordValidateError = 21002
)
