package keeper

import (
	"crypto/sha256"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/adelapie/tttv2/x/tttv2/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetMatchCount get the total number of match
func (k Keeper) GetMatchCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.MatchCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetMatchCount set the total number of match
func (k Keeper) SetMatchCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.MatchCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateMatch creates a match
func (k Keeper) CreateMatch(ctx sdk.Context, msg types.MsgCreateMatch) {
	// Create the match
	count := k.GetMatchCount(ctx)

	op_tmp, _ := sdk.AccAddressFromBech32(msg.Opponent)

	var match = types.Match{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        Opponent: op_tmp, //msg.Opponent,
		Status: "New game",
    }

	// ease position look in the board
	
	for i := 0; i < 9; i++ {
		match.Board[i] = strconv.Itoa(i)
	}

	// decide nextPlayer before game starts

	creator_byte := []byte(msg.Creator)
	opponent_byte := []byte(msg.Opponent)

	hash_input := append(creator_byte, opponent_byte...)
	h := sha256.New()
	h.Write(hash_input)
	digest := h.Sum(nil)
	bit := digest[0] & 1

	if bit == 1 {
		match.NextPlayer = msg.Creator 
		match.RoleOpponent = "O"
		match.RoleCreator = "X"
	} else {
		match.NextPlayer = op_tmp
		match.RoleOpponent = "X"
		match.RoleCreator = "O"
	}

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.MatchPrefix + match.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(match)
	store.Set(key, value)

	// Update match count
    k.SetMatchCount(ctx, count+1)
}

// GetMatch returns the match information
func (k Keeper) GetMatch(ctx sdk.Context, key string) (types.Match, error) {
	store := ctx.KVStore(k.storeKey)
	var match types.Match
	byteKey := []byte(types.MatchPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &match)
	if err != nil {
		return match, err
	}
	return match, nil
}

// SetMatch sets a match
func (k Keeper) SetMatch(ctx sdk.Context, match types.Match) {
	matchKey := match.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(match)
	key := []byte(types.MatchPrefix + matchKey)
	store.Set(key, bz)
}

// DeleteMatch deletes a match
func (k Keeper) DeleteMatch(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.MatchPrefix + key))
}

//
// Functions used by querier
//

func listMatch(ctx sdk.Context, k Keeper) ([]byte, error) {
	var matchList []types.Match
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.MatchPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var match types.Match
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &match)
		matchList = append(matchList, match)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, matchList)
	return res, nil
}

func getMatch(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	match, err := k.GetMatch(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, match)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetMatchOwner(ctx sdk.Context, key string) sdk.AccAddress {
	match, err := k.GetMatch(ctx, key)
	if err != nil {
		return nil
	}
	return match.Creator
}


// Check if the key exists in the store
func (k Keeper) MatchExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.MatchPrefix + key))
}
