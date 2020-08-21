<template>
  <section>
    <SearchBar
      v-on:searchEvent="handleSearch"
      v-on:cancelSearchEvent="cancelSearch"
      :search-enabled="this.searchModeEnabled"
    />
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
          <td>
            <router-link class="eye-link" :to="`/logs/${log.uid}`">
              <font-awesome-icon icon="eye" />
            </router-link>
          </td>
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
import SearchBar from "@/components/SearchBar.vue";
@Component({
  components: { SearchBar }
})
export default class LogsList extends Vue {
  private currentPage = 1;
  private searchModeEnabled = false;
  private searchQuery = "";

  private logs: LogsEndpointResponse = {
    data: [],
    applicationName: "",
    entriesPerPage: 50
  };

  private now: number = Math.round(new Date().getTime() / 1000);

  async mounted(): Promise<void> {
    this.logs = await LogService.getLogs(this.currentPage);
  }

  timeDiffToHuman(value: number): string {
    return intervalToLevels(value);
  }

  async nextPage(): Promise<void> {
    this.currentPage++;
    if (this.searchModeEnabled) {
      const received = await LogService.searchLogs(
        this.currentPage,
        this.searchQuery
      );
      if (received.data !== null) {
        this.logs = received;
      } else {
        this.currentPage--;
      }
    } else {
      const received = await LogService.getLogs(this.currentPage);
      if (received.data !== null) {
        this.logs = received;
      } else {
        this.currentPage--;
      }
    }
  }

  async previousPage(): Promise<void> {
    if (this.currentPage > 1) {
      this.currentPage--;
    }
    if (this.searchModeEnabled) {
      this.logs = await LogService.searchLogs(
        this.currentPage,
        this.searchQuery
      );
    } else {
      this.logs = await LogService.getLogs(this.currentPage);
    }
  }

  async handleSearch(searchQuery: string): Promise<void> {
    this.currentPage = 1;
    this.searchModeEnabled = true;
    this.searchQuery = searchQuery;
    this.logs = await LogService.searchLogs(this.currentPage, searchQuery);
  }

  async cancelSearch(): Promise<void> {
    this.currentPage = 1;
    this.searchModeEnabled = false;
    this.searchQuery = "";
    this.logs = await LogService.getLogs(this.currentPage);
  }
}
</script>
