package types

const (
	// ModuleName defines the module name
	ModuleName = "dezen"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dezen"
)

var (
	ParamsKey = []byte("p_dezen")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
