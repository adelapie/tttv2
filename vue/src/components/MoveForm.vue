<template>
  <div>
    <sp-input placeholder="Match" v-model="matchID" />
    <sp-input placeholder="Position" v-model="position" />
    <sp-button @click.native="submit">Move</sp-button>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
export default {
  components: { ...sp },
  data() {
    return {
      matchID: "",
      position: ""
    };
  },
  methods: {
    async submit() {
      const payload = {
        type: "move",
        module: "tttv2",
        body: {
          matchID: this.matchID,
          position: this.position
        }
      };
      await this.$store.dispatch("cosmos/entitySubmit", payload);
        //await this.$store.dispatch("cosmos/entityFetch", payload);
    }
  }
};


</script>
