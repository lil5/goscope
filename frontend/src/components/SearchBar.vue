<template>
  <section>
    <label for="search-input" content="Search"></label
    ><input
      v-model="searchQuery"
      placeholder="Search ..."
      id="search-input"
      style="float: left"
      type="search"
    />
    <button
      v-on:click="emitSearchEvent"
      id="search-button"
      style="display: flex; border: none; float:left;"
    >
      <font-awesome-icon icon="search" />
    </button>
    <button
      v-if="this.searchEnabled"
      v-on:click="emitCancelSearchEvent"
      id="search-cancel-button"
      style="border: none; float:left;"
    >
      <font-awesome-icon icon="times" />
    </button>
  </section>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";

export default Vue.extend({
  name: "SearchBar",
  props: {
    searchEnabled: Boolean as PropType<boolean>
  },
  data() {
    return {
      searchQuery: ""
    };
  },
  methods: {
    emitSearchEvent() {
      if (this.$data.searchQuery !== "") {
        this.$emit("searchEvent", this.$data.searchQuery);
      }
    },
    emitCancelSearchEvent() {
      this.$emit("cancelSearchEvent");
    }
  }
});
</script>
