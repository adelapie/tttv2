package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers tttv2-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/tttv2/move", createMoveHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/tttv2/move", listMoveHandler(cliCtx, "tttv2")).Methods("GET")
		r.HandleFunc("/tttv2/move/{key}", getMoveHandler(cliCtx, "tttv2")).Methods("GET")
		r.HandleFunc("/tttv2/move", setMoveHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/tttv2/move", deleteMoveHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/tttv2/match", createMatchHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/tttv2/match", listMatchHandler(cliCtx, "tttv2")).Methods("GET")
		r.HandleFunc("/tttv2/match/{key}", getMatchHandler(cliCtx, "tttv2")).Methods("GET")
		r.HandleFunc("/tttv2/match", setMatchHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/tttv2/match", deleteMatchHandler(cliCtx)).Methods("DELETE")

		
}
