<template>
  <div id="app">
    <Navbar />
    <router-view />
    <Footer />
  </div>
</template>

<script lang="ts">
import Footer from "@/components/Footer";
import Navbar from "@/components/Navbar.vue";

export default {
  components: { Navbar, Footer }
};
</script>
