package types

const (
	// ModuleName is the name of the module
	ModuleName = "tttv2"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	MatchPrefix = "match-value-"
	MatchCountPrefix = "match-count-"
)
		
const (
	MovePrefix = "move-value-"
	MoveCountPrefix = "move-count-"
)
		