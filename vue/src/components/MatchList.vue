<template>
  <div>
    <SpH3>
      Invite someone to play tic-tac-toe!
    </SpH3>
    <div>
      <sp-input placeholder="Opponent" v-model="opponent" />                          
      <sp-button @click.native="submit">Create match</sp-button>   
    </div>
    <div v-for="match in matches" v-bind:key="match.id">
      <SpH3>
      Match {{ match.id }} {{ match.status }}
      </SpH3>
      <app-text type="subtitle_red">Created by: {{ match.creator }}</app-text>
      <app-text type="subtitle_red">Opponent: {{ match.opponent }}</app-text>
      <app-text type="subtitle_red">Next player: {{ match.nextplayer }}</app-text>
      <br>
      <tic-tac-toe :pos0=match.board[0] :pos1=match.board[1] :pos2=match.board[2]
                   :pos3=match.board[3] :pos4=match.board[4] :pos5=match.board[5]
                   :pos6=match.board[6] :pos7=match.board[7] :pos8=match.board[8]
        />
      <br>
      <app-text type="subtitle_red">Winner: {{ match.winner }}</app-text>


      <app-text type="h2">Make a move</app-text>
      
      <div>                                                                         
       <sp-input placeholder="Position" v-model="position"/>                      
       <sp-button @click.native="submitmove(match.id)">Move</sp-button>                          
     </div>  

    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
//import {countBy } from "lodash"
import AppText from "./AppText";
import Cell from "./Cell";
import tictactoeboard from "./Tictactoeboard"
export default {
  components: {AppText,  ...sp },
  data() {
    return {
      opponent: ""
    };
  },
  computed: {
    matches() {
      return this.$store.state.cosmos.data["tttv2/match"] || [];
    }
  },
  methods: {
    async submit() {
      const type = { type: "match" };
      const payload = {
        type: "match",
        module: "tttv2",
        body: {
          opponent: this.opponent
        }
      };
      await this.$store.dispatch("cosmos/entitySubmit", payload);
      await this.$store.dispatch("cosmos/entityFetch", {...type, module: "tttv2"});
    },
 
    async submitmove(match_id) {
      const type = { type: "move" };
      const payload = {
        type: "move",
        module: "tttv2",
        body: {
          matchID: match_id,
          position: this.position
        }
      };
      await this.$store.dispatch("cosmos/entitySubmit", payload);
    }
  }
};
</script>

