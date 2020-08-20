<template>
  <header>
    <nav role="navigation">
      <router-link to="/">
        <img width="100" height="100" id="logo" src="@/assets/logo.svg" />
        {{ applicationName }}
      </router-link>

      <router-link to="/">
        <font-awesome-icon icon="sync" />
        Requests
      </router-link>
      <router-link to="/logs">
        <font-awesome-icon icon="clipboard-list" />
        Logs
      </router-link>
      <router-link to="/info">
        <font-awesome-icon icon="server" />
        System
      </router-link>
    </nav>
  </header>
</template>

<script>
export default {
  name: "Navbar"
};
</script>

<style scoped></style>
