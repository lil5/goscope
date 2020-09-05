<template>
  <section>
    <SearchBar
      v-on:searchEvent="this.handleSearch"
      v-on:cancelSearchEvent="this.cancelSearch"
      :search-enabled="this.searchModeEnabled"
    />
    <table>
      <thead>
        <tr>
          <th>Status</th>
          <th>Verb</th>
          <th>Path</th>
          <th>Happened</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr :key="index" v-for="(request, index) in requests.data">
          <td v-html="applyStatusColor(request.responseStatus)"></td>
          <td v-html="applyMethodColor(request.method)"></td>
          <td>{{ request.path }}</td>
          <td>{{ timeDiffToHuman(now - request.time) }} ago</td>
          <td>
            <router-link class="eye-link" :to="`/requests/${request.uid}`">
              <font-awesome-icon icon="eye" />
            </router-link>
          </td>
        </tr>
      </tbody>
    </table>
    <nav class="text-center">
      <button v-on:click="this.previousPage()">← prev</button>
      <button v-on:click="this.nextPage()">next →</button>
    </nav>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import SearchBar from "@/components/SearchBar.vue";
import { RequestService } from "@/api/requests";
import { RequestsEndpointResponse } from "@/interfaces/requests";
import { intervalToLevels } from "@/utils/time";

export default Vue.extend({
  name: "RequestList",
  components: { SearchBar },
  data() {
    return {
      requests: {} as RequestsEndpointResponse,
      currentPage: 1,
      searchModeEnabled: false,
      searchQuery: "",
      now: Math.round(new Date().getTime() / 1000)
    };
  },
  async created(): Promise<void> {
    this.$data.requests = await RequestService.getRequests(
      this.$data.currentPage
    );
  },
  methods: {
    async nextPage(): Promise<void> {
      this.$data.currentPage++;
      if (this.$data.searchModeEnabled) {
        const received = await RequestService.searchRequests(
          this.$data.currentPage,
          this.$data.searchQuery
        );
        if (received.data && received.data.length > 0) {
          this.$data.requests = received;
        } else {
          this.$data.currentPage--;
        }
      } else {
        const received = await RequestService.getRequests(
          this.$data.currentPage
        );
        if (received.data && received.data.length > 0) {
          this.$data.requests = received;
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
        this.$data.requests = await RequestService.searchRequests(
          this.$data.currentPage,
          this.$data.searchQuery
        );
      } else {
        this.$data.requests = await RequestService.getRequests(
          this.$data.currentPage
        );
      }
    },
    async handleSearch(searchQuery: string): Promise<void> {
      this.$data.currentPage = 1;
      this.$data.searchModeEnabled = true;
      this.$data.searchQuery = searchQuery;
      this.$data.requests = await RequestService.searchRequests(
        this.$data.currentPage,
        searchQuery
      );
    },
    async cancelSearch(): Promise<void> {
      this.$data.currentPage = 1;
      this.$data.searchModeEnabled = false;
      this.$data.searchQuery = "";
      this.$data.requests = await RequestService.getRequests(
        this.$data.currentPage
      );
    },
    timeDiffToHuman(value: number): string {
      return intervalToLevels(value);
    },
    applyMethodColor(method: string): string {
      if (method === "GET") {
        return `<span class="badge-secondary">${method}</span>`;
      } else if (method === "POST") {
        return `<span class="badge-info">${method}</span>`;
      } else if (method === "PUT") {
        return `<span class="badge-info">${method}</span>`;
      } else if (method === "PATCH") {
        return `<span class="badge-turq">${method}</span>`;
      } else if (method === "DELETE") {
        return `<span class="badge-danger">${method}</span>`;
      }
      return `<span class="badge-secondary">${method}</span>`;
    },
    applyStatusColor(status: number): string {
      if (status >= 200 && status < 300) {
        return `<span class="badge-success">${status}</span>`;
      } else if (status >= 300 && status < 400) {
        return `<span class="badge-info">${status}</span>`;
      } else if (status >= 400 && status < 500) {
        return `<span class="badge-warning">${status}</span>`;
      }
      return `<span class="badge-danger">${status}</span>`;
    }
  }
});
</script>
