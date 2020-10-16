<template>
  <section class="search-bar">
    <TagsInput
      v-model="searchTag"
      placeholder="Search ..."
      :autocomplete-items="filteredAutocomplete"
      add-only-from-autocomplete
      :tags="selectedTags"
      @tags-changed="editSelectedTags"
      :autocomplete-min-length="0"
      @before-adding-tag="beforeAddingTag"
    />

    <button
      v-on:click="emitSearchEvent"
      id="search-button"
      style="display: flex; border: none; float:left;"
    >
      <font-awesome-icon icon="search" />
    </button>
    <button
      v-show="showDisabled"
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
import { Tag } from "../interfaces/filter";
//@ts-ignore
import VueTagsInput from "@johmun/vue-tags-input";

export default Vue.extend({
  name: "SearchBar",
  components: { TagsInput: VueTagsInput },
  props: {
    searchEnabled: Boolean as PropType<boolean>,
    autocomplete: {
      type: Array as PropType<Tag[]>,
      default: []
    }
  },
  data() {
    return {
      showDisabled: false as boolean,
      searchTag: "" as string,
      selectedTags: [] as Tag[]
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
    checkShowDisabled(): void {
      const emptySearch = this.searchTag.length === 0;
      const emptyTags = this.selectedTags.length === 0;

      if (emptySearch && emptyTags) {
        this.showDisabled = false;
      }

      this.showDisabled = true;
    },
    beforeAddingTag(o: { tag: Tag; addTag: Function }) {
      o.addTag();

      if (o.tag.text === this.searchTag) {
        this.emitSearchEvent();
      }
    },
    editSelectedTags(newSelectedTags: Tag[]) {
      this.selectedTags = newSelectedTags;
    },
    emitSearchEvent() {
      if (this.searchTag === "" && this.selectedTags.length === 0) return;

      this.checkShowDisabled();
      this.$emit("searchEvent", this.searchTag, this.selectedTags);
    },
    emitCancelSearchEvent() {
      this.searchTag = "";
      this.selectedTags = [];
      this.checkShowDisabled();
      this.$emit("cancelSearchEvent");
    }
  }
});
</script>
