<template>
  <div>
    <router-view />
  </div>
</template>

<style>
.sp-container {
  margin: 0 auto;
  max-width: 800px;
  padding: 1rem;
}
</style>

<script>
export default {
  created() {
    this.$store.dispatch("cosmos/init");
    //this.$store.dispatch("cosmos/entitySubmit", {type: "move", module: "tttv2"});
    this.$store.dispatch("cosmos/entitySubmit", {type: "match", module: "tttv2"});
    this.$store.dispatch("cosmos/entityFetch", {type: "match", module: "tttv2"});
  },
};
</script>
