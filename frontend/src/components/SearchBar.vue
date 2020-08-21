<template>
  <section>
    <input
      v-model="searchQuery"
      placeholder="Search ..."
      id="search-input"
      style="float: left"
      type="search"
    />
    <button
      v-on:click="emitEvent"
      id="search-button"
      style="display: flex; border: none; float:left;"
    >
      <font-awesome-icon icon="search" />
    </button>
    <button
      id="search-cancel-button"
      style="border: none; display: none; float:left;"
    >
      <font-awesome-icon icon="times" />
    </button>
  </section>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class SearchBar extends Vue {
  private searchQuery = "";

  emitEvent() {
    if (this.searchQuery !== "") {
      this.$emit("searchEvent", this.searchQuery);
    }
  }
}
</script>
