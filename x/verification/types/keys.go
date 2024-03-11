package types

const (
	// ModuleName defines the module name
	ModuleName = "verification"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_verification"
)

var (
	ParamsKey = []byte("p_verification")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
