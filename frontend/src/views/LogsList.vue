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
import Vue from "vue";
import { intervalToLevels } from "@/utils/time";
import { LogService } from "@/api/logs";
import { LogsEndpointResponse } from "@/interfaces/logs";
import SearchBar from "@/components/SearchBar.vue";

export default Vue.extend({
  name: "LogsList",
  components: { SearchBar },
  data() {
    return {
      currentPage: 1,
      searchModeEnabled: false,
      searchQuery: "",
      logs: {} as LogsEndpointResponse,
      now: Math.round(new Date().getTime() / 1000)
    };
  },
  methods: {
    timeDiffToHuman(value: number): string {
      return intervalToLevels(value);
    },
    async nextPage(): Promise<void> {
      this.$data.currentPage++;
      if (this.$data.searchModeEnabled) {
        const received = await LogService.searchLogs(
          this.$data.currentPage,
          this.$data.searchQuery
        );
        if (!received.data || received.data.length === 0) {
          this.$data.logs = received;
        } else {
          this.$data.currentPage--;
        }
      } else {
        const received = await LogService.getLogs(this.$data.currentPage);
        if (received.data !== null) {
          this.$data.logs = received;
        } else {
          this.$data.currentPage--;
        }
      }
    },
    async previousPage(): Promise<void> {
      if (this.$data.currentPage > 1) {
        this.$data.currentPage--;
      }
      if (this.$data.searchModeEnabled) {
        this.$data.logs = await LogService.searchLogs(
          this.$data.currentPage,
          this.$data.searchQuery
        );
      } else {
        this.$data.logs = await LogService.getLogs(this.$data.currentPage);
      }
    },
    async handleSearch(searchQuery: string): Promise<void> {
      this.$data.currentPage = 1;
      this.$data.searchModeEnabled = true;
      this.$data.searchQuery = searchQuery;
      this.$data.logs = await LogService.searchLogs(
        this.$data.currentPage,
        searchQuery
      );
    },
    async cancelSearch(): Promise<void> {
      this.$data.currentPage = 1;
      this.$data.searchModeEnabled = false;
      this.$data.searchQuery = "";
      this.$data.logs = await LogService.getLogs(this.$data.currentPage);
    }
  },
  async created(): Promise<void> {
    this.logs = await LogService.getLogs(this.currentPage);
  }
});
</script>
