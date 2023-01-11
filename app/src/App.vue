<template>
  <navbar />
  <router-view />
</template>

<script>
import Navbar from "@/components/Navbar.vue";

export default {
  components: {
    Navbar,
  },
};
</script>

<style lang="scss">
.content-area {
  margin-left: 5%;
  background: white;
  width: 87%;
  height: 70%;
}
</style>
