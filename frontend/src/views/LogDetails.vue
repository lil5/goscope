<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <section>
      <h2>Log details</h2>
      <dl>
        <dt>Message</dt>
        <dd>
          <pre><code>{{ computedLogError }}</code></pre>
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
import { epochToHumanDate } from "@/utils/time";

export default Vue.extend({
  name: "LogDetails",
  props: {
    logUUID: String as PropType<string>
  },
  computed: {
    computedLogError(): string {
      if (!this.logDetails.data) {
        return "";
      }
      return this.logDetails.data.logDetails.error;
    },
    computedLogTime(): string {
      if (!this.logDetails.data) {
        return "";
      }
      return epochToHumanDate(this.logDetails.data.logDetails.time);
    }
  },
  async created(): Promise<void> {
    this.logDetails = await LogService.getLog(this.$route.params.id);
    document.title = `${this.logDetails.applicationName} | Log ${this.$route.params.id}`;
  },
  data() {
    return {
      logDetails: {} as DetailedLogsReponse
    };
  }
});
</script>

<style scoped>
pre {
  white-space: pre-wrap;       /* css-3 */
  white-space: -moz-pre-wrap;  /* Mozilla, since 1999 */
  white-space: -pre-wrap;      /* Opera 4-6 */
  white-space: -o-pre-wrap;    /* Opera 7 */
  word-wrap: break-word;       /* Internet Explorer 5.5+ */
}
</style>