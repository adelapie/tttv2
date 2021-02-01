<template>
  <div>
    <div :class="[`${type}`]">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.h1 {
  font-size: 2rem;
  font-weight: 800;
  letter-spacing: 0.03em;
  margin-bottom: 2rem;
}
.h2 {
  margin-top: 1rem;
  font-weight: 800;
  font-size: 0.85rem;
  color: rgb(0,0,0);
}
.subtitle {
  color: rgba(0, 0, 0, 0.5);
  font-size: 0.85rem;
}
.subtitle_red {
  color: rgb(255, 0, 0 );
  font-size: 0.85rem;
}
.board_no {
  color: rgba(255, 0, 0, 0.5 );
  font-size: 0.85rem;
}
</style>

<script>
export default {
  props: {
    type: {
      default: "p1",
    },
  },
};
</script>

