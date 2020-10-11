<template>
  <fieldset :class="{ filter__inactive: false }">
    <legend>Method</legend>

    <template v-for="method in methods">
      <label :key="method">
        {{ method }}
        <input type="checkbox" />
      </label>
      <br :key="method" />
    </template>
  </fieldset>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";
import { Method } from "@/interfaces/filter";
export default Vue.extend({
  props: {
    input: {
      required: true,
      type: Array as PropType<Method[]>
    }
  },
  data() {
    return {
      methods: [
        Method.GET,
        Method.HEAD,
        Method.POST,
        Method.PUT,
        Method.DELETE,
        Method.CONNECT,
        Method.OPTIONS,
        Method.TRACE,
        Method.PATCH
      ]
    };
  },
  computed: {}
});
</script>
