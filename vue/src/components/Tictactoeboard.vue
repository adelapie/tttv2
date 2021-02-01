<template>
  <div class="tictactoeboard">


<!--
<div v-for="(n, i) in 3">
      <div v-for="(n, j) in 3">
        <cell :value="board[i][j]"></cell> 

-->
<cell :value=pos0></cell>
<cell :value=pos1></cell>
<cell :value=pos2></cell>
<cell :value=pos3></cell>
<cell :value=pos4></cell>
<cell :value=pos5></cell>
<cell :value=pos6></cell>
<cell :value=pos7></cell>
<cell :value=pos8></cell>

      </div>
    </div>
  </div>
</template>
<script>
import Cell from "./Cell";                                                      
export default {
    name: 'tictactoeboard',
  props: { pos0: { type: String},
           pos1: { type: String},
           pos2: { type: String},
           pos3: { type: String},
           pos4: { type: String},
           pos5: { type: String},
           pos6: { type: String},
           pos7: { type: String},
           pos8: { type: String}
  },
    data() { return {
      board: [
          ['', '', ''],
          ['', 'x', ''],
          ['', '', '']
        ]
    } },
    
  }
</script>
<style>
  .tictactoeboard {
    display: flex;
    flex-wrap: wrap;
    width: 204px;
    height: 204px;
  }
</style>
