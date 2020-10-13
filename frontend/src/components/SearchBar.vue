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

    <TagsInput element-id="tags"
    v-model="selectedTags"
    placeholder="Search ..."
    id-field="id"
    text-field="name"
    :existing-tags="existingTags"
    :typeahead='true'
  />
  </section>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";
import { Method } from "@/interfaces/filter";
import {EnumReflection} from '@/utils/enum-reflection';
import VoerroTagsInput from '@voerro/vue-tagsinput';

interface SelectedTag {
  key: string
  value: string
}
interface ExistingTag {
  id: number
  name: string
}

function generateExistingTags(): ExistingTag[] {
  const existingTags: ExistingTag[] = [];
    let existingTagIndex = 0;

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
]
    methods.forEach(m=> {
      existingTags.push({
        id: existingTagIndex,
        name: m
      });

      existingTagIndex++;
    })
  
    const statuses = [
  '1xx', '2xx', '3xx', '4xx', '5xx'
]
    statuses.forEach((s) => {
      existingTags.push({
        id: existingTagIndex,
        name: s
      })

      existingTagIndex++
    })

    console.log(existingTags)

    return existingTags
}

export default Vue.extend({
  name: "SearchBar",
  components: { TagsInput: VoerroTagsInput },
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
      selectedTags: [] as SelectedTag[],
      existingTags: generateExistingTags(),
    };
  },
  methods: {
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
