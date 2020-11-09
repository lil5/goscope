<template>
  <section>
    <SearchBar
      v-on:searchEvent="handleSearch"
      v-on:cancelSearchEvent="cancelSearch"
      :search-enabled="this.searchModeEnabled"
      :autocomplete="[]"
    />
    <button :class="autoRefreshButtonClass" @click="autoRefresh = !autoRefresh">
      <font-awesome-icon icon="sync" />&nbsp;Auto-Refresh
    </button>
    <table>
      <thead>
        <tr>
          <th style="width: 70%">Message</th>
          <th style="width: 25%">Time</th>
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
      <button v-on:click="this.previousPage" class="navbar-link cursor-pointer">
        ← prev
      </button>
      <button v-on:click="this.nextPage" class="navbar-link cursor-pointer">
        next →
      </button>
    </nav>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { intervalToLevels } from "@/utils/time";
import { LogService } from "@/api/logs";
import { LogsEndpointResponse } from "@/interfaces/logs";
import SearchBar from "@/components/SearchBar.vue";
import { RequestService } from "@/api/requests";

export default Vue.extend({
  name: "LogsList",
  components: { SearchBar },
  data() {
    return {
      currentPage: 1,
      autoRefresh: false,
      timer: 0,
      searchModeEnabled: false,
      searchQuery: "",
      logs: {} as LogsEndpointResponse,
      now: Math.round(new Date().getTime() / 1000)
    };
  },
  computed: {
    autoRefreshButtonClass(): string {
      if (!this.autoRefresh) {
        return "";
      }
      return "active-auto-refresh";
    }
  },
  methods: {
    timeDiffToHuman(value: number): string {
      return intervalToLevels(value);
    },
    async nextPage(): Promise<void> {
      this.currentPage++;
      if (this.searchModeEnabled) {
        const received = await LogService.searchLogs(
          this.currentPage,
          this.searchQuery
        );
        if (received.data && received.data.length > 0) {
          this.logs = received;
        } else {
          this.currentPage--;
        }
      } else {
        const received = await LogService.getLogs(this.currentPage);
        if (received.data && received.data.length > 0) {
          this.logs = received;
        } else {
          this.currentPage--;
        }
      }
    },
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
    },
    async handleSearch(searchQuery: string): Promise<void> {
      this.currentPage = 1;
      this.searchModeEnabled = true;
      this.searchQuery = searchQuery;
      this.logs = await LogService.searchLogs(this.currentPage, searchQuery);
    },
    async cancelSearch(): Promise<void> {
      this.currentPage = 1;
      this.searchModeEnabled = false;
      this.searchQuery = "";
      this.logs = await LogService.getLogs(this.currentPage);
    }
  },
  async created(): Promise<void> {
    this.logs = await LogService.getLogs(this.currentPage);
    document.title = `${this.logs.applicationName} | Logs`;
    this.timer = setInterval(async () => {
      if (this.autoRefresh) {
        this.logs = await LogService.getLogs(this.currentPage);
      }
    }, 5000);
  }
});
</script>
