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

    <button
      style="border: none; float: right;"
      v-if="hasFilter"
      v-on:click="toggleFilterOpen"
    >
      <font-awesome-icon icon="filter" />
      Filter
      <font-awesome-icon icon="angle-left" v-show="isFilterOpen" />
      <font-awesome-icon icon="angle-right" v-show="!isFilterOpen" />
    </button>
    <aside
      class="filter-bar"
      v-if="hasFilter"
      :class="{ 'filter-bar--open': isFilterOpen }"
    >
      <button v-on:click="toggleFilterOpen" class="filter-bar__close">
        <font-awesome-icon icon="angle-right" />
      </button>
      <h2>Filter</h2>
      <slot name="filter"></slot>
    </aside>
  </section>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";

export default Vue.extend({
  name: "SearchBar",
  props: {
    searchEnabled: Boolean as PropType<boolean>,
    hasFilter: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      searchQuery: "" as string,
      isFilterOpen: false as boolean
    };
  },
  methods: {
    toggleFilterOpen(): void {
      this.isFilterOpen = !this.isFilterOpen;
    },
    emitSearchEvent() {
      if (this.searchQuery !== "") {
        this.$emit("searchEvent", this.searchQuery);
      }
    },
    emitCancelSearchEvent() {
      this.$emit("cancelSearchEvent");
    }
  }
});
</script>
