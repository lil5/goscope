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
import { Component, Vue } from "vue-property-decorator";
import { PropType } from "vue";

@Component({
  props: {
    searchEnabled: Boolean as PropType<boolean>
  }
})
export default class SearchBar extends Vue {
  private searchQuery = "";

  emitSearchEvent() {
    if (this.searchQuery !== "") {
      this.$emit("searchEvent", this.searchQuery);
    }
  }

  emitCancelSearchEvent() {
    this.$emit("cancelSearchEvent");
  }
}
</script>
