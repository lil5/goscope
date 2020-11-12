<template>
  <header>
    <nav role="navigation">
      <router-link to="/">
        <strong>{{ this.applicationDetails.applicationName }}</strong>
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

<script lang="ts">
import { ApplicationDetailsResponse } from "@/interfaces/app-details";
import { ApplicationDetailsService } from "@/api/app-details";
import Vue from "vue";

export default Vue.extend({
  name: "Navbar",

  async mounted(): Promise<void> {
    this.$data.applicationDetails = await ApplicationDetailsService.getApplicationDetails();
  },
  data() {
    return {
      applicationDetails: {} as ApplicationDetailsResponse
    };
  }
});
</script>

<style scoped>
.navbar-logo {
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
  padding: 10px 5px 8px;
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
