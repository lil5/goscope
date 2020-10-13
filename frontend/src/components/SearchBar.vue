<template>
  <section class="search-bar">
    <TagsInput
      v-model="searchTag"
      placeholder="Search ..."
      :autocomplete-items="filteredAutocomplete"
      add-only-from-autocomplete
      :tags="selectedTags"
      @tags-changed="editSelectedTags"
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
import { Tag } from "@/interfaces/filter";
import VueTagsInput from "@johmun/vue-tags-input";

function generateAutocomplete(): Tag[] {
  const autocomplete: Tag[] = [];

  const methods: string[] = [
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH"
  ];
  methods.forEach(m => {
    autocomplete.push({
      text: "method:" + m.toLowerCase(),
      group: "method",
      value: m
    });
  });

  const statuses = ["1xx", "2xx", "3xx", "4xx", "5xx"];
  statuses.forEach(s => {
    autocomplete.push({
      text: "status:" + s,
      group: "status",
      value: s
    });
  });

  return autocomplete;
}

export default Vue.extend({
  name: "SearchBar",
  components: { TagsInput: VueTagsInput },
  props: {
    searchEnabled: Boolean as PropType<boolean>,
    hasFilter: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      searchTag: "" as string,
      selectedTags: [] as Tag[],
      autocomplete: generateAutocomplete()
    };
  },
  computed: {
    filteredAutocomplete(): Tag[] {
      return this.autocomplete.filter(i => {
        return (
          i.text.toLowerCase().indexOf(this.searchTag.toLowerCase()) !== -1
        );
      });
    }
  },
  methods: {
    editSelectedTags(newSelectedTags: Tag[]) {
      this.selectedTags = newSelectedTags;
    },
    emitSearchEvent() {
      if (this.searchTag !== "") {
        this.$emit("searchEvent", this.searchTag, this.selectedTags);
      }
    },
    emitCancelSearchEvent() {
      this.$emit("cancelSearchEvent");
    }
  }
});
</script>
