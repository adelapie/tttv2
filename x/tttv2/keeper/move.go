package keeper

import (
	"reflect"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/adelapie/tttv2/x/tttv2/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetMoveCount get the total number of move
func (k Keeper) GetMoveCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.MoveCountPrefix)
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

// SetMoveCount set the total number of move
func (k Keeper) SetMoveCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.MoveCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateMove creates a move
func (k Keeper) CreateMove(ctx sdk.Context, msg types.MsgCreateMove) {
	// Create the move
	count := k.GetMoveCount(ctx)
    var move = types.Move{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        MatchID: msg.MatchID,
        Position: msg.Position,
    }



	// 1. recover match

	match_tmp, _ :=  k.GetMatch(ctx, msg.MatchID)
	// match found
	if match_tmp.ID == msg.MatchID {
 
		// this is a new game or a game in progress
		if match_tmp.Status == "New game" || match_tmp.Status == "In progress" {
			match_tmp.Status = "In progress"
			// i'm next ?
			if reflect.DeepEqual(msg.Creator, match_tmp.NextPlayer) {
				i, _ := strconv.Atoi(msg.Position)
				if match_tmp.Board[i] != "X" && match_tmp.Board[i] != "O" {
					// i'm the game's creator ?
					if reflect.DeepEqual(msg.Creator, match_tmp.Creator) {
						match_tmp.Board[i] = match_tmp.RoleCreator 
						match_tmp.NextPlayer = match_tmp.Opponent
					} else {
						match_tmp.Board[i] = match_tmp.RoleOpponent
						match_tmp.NextPlayer = match_tmp.Creator
					}
				}
			}
			
			// is the game finished ?

			// 1. check for finished game, board is full
			is_full := true

			for i := 0; i < 9; i++ {
				if match_tmp.Board[i] == "0" ||
				 match_tmp.Board[i] == "1" ||
				 match_tmp.Board[i] == "2" ||
				 match_tmp.Board[i] == "3" ||
				 match_tmp.Board[i] == "4" ||
				 match_tmp.Board[i] == "5" ||
				 match_tmp.Board[i] == "6" ||
				 match_tmp.Board[i] == "7" ||
				 match_tmp.Board[i] == "8" { 
					is_full = false 
				}
			}

			if is_full == true { match_tmp.Status = "Finished" }

			// 2. check for winning combination

			if (match_tmp.Board[0] == match_tmp.Board[1] && 
			   match_tmp.Board[1] == match_tmp.Board[2] && 
			   (match_tmp.Board[2] == "X" || match_tmp.Board[2] == "O")) ||  
			   (match_tmp.Board[3] == match_tmp.Board[4] && 
			   match_tmp.Board[4] == match_tmp.Board[5] &&
			   (match_tmp.Board[5] == "X" || match_tmp.Board[5] == "O")) ||  
			   (match_tmp.Board[6] == match_tmp.Board[7] && 
			   match_tmp.Board[7] == match_tmp.Board[8] &&
			   (match_tmp.Board[8] == "X" || match_tmp.Board[8] == "O")) ||  
			   (match_tmp.Board[0] == match_tmp.Board[3] && 
			   match_tmp.Board[3] == match_tmp.Board[6] &&
			   (match_tmp.Board[6] == "X" || match_tmp.Board[6] == "O")) ||  
			   (match_tmp.Board[1] == match_tmp.Board[4] && 
			   match_tmp.Board[4] == match_tmp.Board[7] &&
			   (match_tmp.Board[7] == "X" || match_tmp.Board[7] == "O")) ||  
			   (match_tmp.Board[2] == match_tmp.Board[5] && 
			   match_tmp.Board[5] == match_tmp.Board[8] &&
			   (match_tmp.Board[8] == "X" || match_tmp.Board[8] == "O")) ||  
			   (match_tmp.Board[0] == match_tmp.Board[4] && 
			   match_tmp.Board[4] == match_tmp.Board[8] &&
			   (match_tmp.Board[8] == "X" || match_tmp.Board[8] == "O")) ||  
			   (match_tmp.Board[6] == match_tmp.Board[4] && 
			   match_tmp.Board[4] == match_tmp.Board[2] &&
			   (match_tmp.Board[2] == "X" || match_tmp.Board[2] == "O")) {  
				   match_tmp.Status = "Finished"                                     
	               match_tmp.Winner = msg.Creator    
			}


		} else {
			// game is finished, do nothing
		}

		// update match
		k.SetMatch(ctx, match_tmp)
}


	// store move
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.MovePrefix + move.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(move)
	store.Set(key, value)

	// Update move count
    k.SetMoveCount(ctx, count+1)
}

// GetMove returns the move information
func (k Keeper) GetMove(ctx sdk.Context, key string) (types.Move, error) {
	store := ctx.KVStore(k.storeKey)
	var move types.Move
	byteKey := []byte(types.MovePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &move)
	if err != nil {
		return move, err
	}
	return move, nil
}

// SetMove sets a move
func (k Keeper) SetMove(ctx sdk.Context, move types.Move) {
	moveKey := move.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(move)
	key := []byte(types.MovePrefix + moveKey)
	store.Set(key, bz)
}

// DeleteMove deletes a move
func (k Keeper) DeleteMove(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.MovePrefix + key))
}

//
// Functions used by querier
//

func listMove(ctx sdk.Context, k Keeper) ([]byte, error) {
	var moveList []types.Move
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.MovePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var move types.Move
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &move)
		moveList = append(moveList, move)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, moveList)
	return res, nil
}

func getMove(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	move, err := k.GetMove(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, move)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetMoveOwner(ctx sdk.Context, key string) sdk.AccAddress {
	move, err := k.GetMove(ctx, key)
	if err != nil {
		return nil
	}
	return move.Creator
}


// Check if the key exists in the store
func (k Keeper) MoveExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.MovePrefix + key))
}
