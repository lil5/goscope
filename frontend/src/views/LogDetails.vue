<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <section>
      <h2>Log details</h2>
      <dl>
        <dt>Message</dt>
        <dd>
          <code>{{ computedLogError }}</code>
        </dd>
        <dt>Time</dt>
        <dd>
          <code>{{ computedLogTime }}</code>
        </dd>
      </dl>
    </section>
  </div>
</template>

<script lang="ts">
import { DetailedLogsReponse } from "@/interfaces/logs";
import { LogService } from "@/api/logs";
import Vue, { PropType } from "vue";

export default Vue.extend({
  name: "LogDetails",
  props: {
    logUUID: String as PropType<string>
  },
  computed: {
    computedLogError(): string {
      if (!this.$data.logDetails.data) {
        return "";
      }
      return this.$data.logDetails.data.logDetails.error;
    },
    computedLogTime(): string {
      if (!this.$data.logDetails.data) {
        return "";
      }
      return new Date(
        this.$data.logDetails.data.logDetails.time * 1000
      ).toLocaleString("nl-NL");
    }
  },
  async created(): Promise<void> {
    this.logDetails = await LogService.getLog(this.$route.params.id);
  },
  data() {
    return {
      logDetails: {} as DetailedLogsReponse
    };
  }
});
</script>
