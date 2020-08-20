<template>
  <section>
    <table>
      <thead>
        <tr>
          <th>Message</th>
          <th>Time</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr :key="index" v-for="(log, index) in logs.data">
          <td>{{ log.error }}</td>
          <td>{{ timeDiffToHuman(now - log.time) }} ago</td>
          <td><font-awesome-icon icon="eye" /></td>
        </tr>
      </tbody>
    </table>
    <nav class="text-center">
      <button v-on:click="previousPage()" class="navbar-link cursor-pointer">
        ← prev
      </button>
      <button v-on:click="nextPage()" class="navbar-link cursor-pointer">
        next →
      </button>
    </nav>
  </section>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { LogsEndpointResponse } from "@/interfaces/logs";
import { LogService } from "@/api/logs";
import { intervalToLevels } from "@/utils/time";

@Component
export default class LogsList extends Vue {
  private logs: LogsEndpointResponse = {
    data: [],
    applicationName: "",
    entriesPerPage: 50
  };

  private now: number = Math.round(new Date().getTime() / 1000);

  async mounted(): Promise<void> {
    this.logs = await LogService.getLogs(1);
  }

  timeDiffToHuman(value: number): string {
    return intervalToLevels(value);
  }
}
</script>

<style scoped></style>
