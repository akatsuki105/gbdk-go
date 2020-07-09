package types

var PrimitiveTypes []string = []string{
	"int", "int8", "int16", "uint", "byte", "uint16", "int32", "uint32", "float32",
}

var GBDKTypes []string = []string{
	"BOOLEAN", "INT8", "UINT8", "UINT16", "UINT32", "BYTE", "WORD", "DWORD",
}

func IsTypeFunc(name string) bool {
	for _, t := range PrimitiveTypes {
		if name == t {
			return true
		}
	}

	for _, t := range GBDKTypes {
		if name == t {
			return true
		}
	}

	return false
}
