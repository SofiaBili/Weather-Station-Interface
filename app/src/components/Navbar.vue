<template>
  <div class="sidenav">
    <router-link to="/"><i class="bi bi-map-fill" /></router-link>
    <router-link to="/data"><i class="bi bi-cloud-fill" /></router-link>
  </div>
</template>

<style lang="scss" scoped>
.sidenav {
  height: 100%;
  width: 80px;
  position: fixed;
  top: 0;
  background-color: rgb(255, 255, 255);
  padding-top: 1rem;

  a {
    padding: 1rem 1rem 1rem 2rem;
    margin-bottom: 4px;
    font-size: 1.8rem;
    color: #0094ce;
    display: block;
    background-position: right;

    &:hover {
      background: linear-gradient(90deg, #eeeeee 50%, #0094ce 50%);
      background-size: 170%;
      transition: .3s ease-out;
    }
  }
}
</style>
