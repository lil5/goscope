<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <section>
      <h2>Log details</h2>
      <dl>
        <dt>Message</dt>
        <dd>
          <code>{{ this.logDetails.data.logDetails.error }}</code>
        </dd>
        <dt>Time</dt>
        <dd>
          <code>{{ this.logDetails.data.logDetails.time }}</code>
        </dd>
      </dl>
    </section>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { DetailedLogsReponse } from "@/interfaces/logs";
import { LogService } from "@/api/logs";
import { PropType } from "vue";

@Component({
  props: {
    logUUID: String as PropType<string>
  }
})
export default class LogDetails extends Vue {
  private logDetails: DetailedLogsReponse = {
    data: {
      logDetails: {
        error: "",
        time: 0,
        uid: ""
      }
    },
    applicationName: ""
  };

  async mounted(): Promise<void> {
    this.logDetails = await LogService.getLog(this.$route.params.id);
  }
}
</script>
