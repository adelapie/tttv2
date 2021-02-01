<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
     <!-- <sp-bank-balances />
     <sp-token-send /> -->
      <!-- this line is used by starport scaffolding # 4 -->
      <!-- <move-form /> -->
       <!-- <sp-type-form type="move" :fields="['matchID', 'position', ]" module="tttv2" />
       	    <sp-type-form type="match" :fields="['opponent', ]" module="tttv2" /> -->
      <match-list />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
//import MoveForm from "../components/MoveForm";
import MatchList from "../components/MatchList";
export default {
  components: {MatchList, ...sp },
};
</script>
