<template>
  <footer>
    <p>
      Made with ❤️ and hosted on
      <a
        class="navbar-link"
        target="_blank"
        rel="noopener noreferrer"
        href="https://github.com/averageflow/goscope"
        >GitHub</a
      >
    </p>
    <dark-mode-toggle
      id="dark-mode-toggle-1"
      legend=""
      appearance="switch"
      dark="Dark"
      light="Light"
      remember="Remember setting?"
    ></dark-mode-toggle>
  </footer>
</template>

<script>
import Vue from "vue";
export default Vue.extend({
  name: "Footer"
});
</script>

<style scoped>
footer {
  border-top: 1px dashed white;
}
</style>
