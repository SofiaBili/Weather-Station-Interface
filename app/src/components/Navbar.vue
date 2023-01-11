<template>
  <div class="sidenav">
    <img
      src="@/assets/AEOLUS_LOGO.png"
      style="position: relative; top: 0px"
      class="img-responsive"
      alt="logo"
    />
    <router-link to="/"><i class="bi bi-map-fill" /> Map</router-link>
    <router-link to="/data"><i class="bi bi-cloud-fill" /> Data</router-link>
    <img
      src="@/assets/logo_pada.png"
      style="position: relative; bottom: 0px"
      class="img-responsive"
      alt="logo"
    />
  </div>
</template>

<style lang="scss" scoped>
.sidenav {
  height: 100%;
  width: 8%;
  position: fixed;
  top: 0;
  background-color: #011627;
  padding-top: 3px;

  a {
    padding: 6px 6px 6px 28px;
    margin-bottom: 4px;
    font-size: 20px;
    color: #437f97;
    display: block;
    background-position: right;
    text-decoration: none;

    &:hover {
      background: linear-gradient(90deg, #a7b9c8 50%, #437f97 50%);
      background-size: 170%;
      transition: 0.3s ease-out;
    }

    .router-link-active,
    .router-link-exact-active {
      background: linear-gradient(90deg, #ffffff 85%, #437f97 15%);
    }
  }

  .img-responsive {
    width: 100%;
  }
}
</style>
