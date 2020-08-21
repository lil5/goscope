<template>
  <header>
    <nav role="navigation">
      <router-link to="/">
        <img id="logo" src="@/assets/logo.svg" />
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

<style scoped>
#logo {
  height: 1.2em;
  margin-right: 5px;
  transform: scale(2);
}

header {
  margin-top: -20px;
}
header nav {
  margin-bottom: 10px;
  padding-left: 5px;
  padding-right: 5px;
  background-color: var(--background);
  border-radius: 0 0 6px 6px;
}
header nav > a {
  text-decoration: none !important;
  border-bottom-style: solid;
  border-bottom-color: transparent;
  border-bottom-width: 2px;
  padding-bottom: 8px;
  padding-top: 10px;
  padding-left: 5px;
  padding-right: 5px;
  display: inline-block;
  color: var(--form-text);
}
header nav > a:not([href="/"]):hover {
  background-color: var(--button-hover);
  border-bottom-color: var(--button-links);
}
header nav > a:not([href="/"]):active,
header nav > a:not([href="/"])[active],
header nav > a:not([href="/"]).active {
  color: var(--links);
  border-bottom-color: var(--links);
}
</style>
