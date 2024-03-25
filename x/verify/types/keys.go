package types

const (
	// ModuleName defines the module name
	ModuleName = "verify"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_verify"
)

var (
	ParamsKey = []byte("p_verify")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
