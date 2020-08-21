<template>
  <section>
    <SearchBar v-on:searchEvent="handleSearch" />
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
      <button v-on:click="previousPage()">← prev</button>
      <button v-on:click="nextPage()">next →</button>
    </nav>
  </section>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { intervalToLevels } from "@/utils/time";
import { RequestsEndpointResponse } from "@/interfaces/requests";
import { RequestService } from "@/api/requests";
import SearchBar from "@/components/SearchBar.vue";

@Component({
  components: { SearchBar }
})
export default class RequestList extends Vue {
  private currentPage = 1;
  private searchModeEnabled = false;
  private searchQuery = "";

  private requests: RequestsEndpointResponse = {
    data: [],
    applicationName: ""
  };

  private now: number = Math.round(new Date().getTime() / 1000);

  async mounted(): Promise<void> {
    this.requests = await RequestService.getRequests(this.currentPage);
  }

  async nextPage(): Promise<void> {
    this.currentPage++;
    if (this.searchModeEnabled) {
      const received = await RequestService.searchRequests(
        this.currentPage,
        this.searchQuery
      );
      if (received.data !== null) {
        this.requests = received;
      } else {
        this.currentPage--;
      }
    } else {
      const received = await RequestService.getRequests(this.currentPage);
      if (received.data !== null) {
        this.requests = received;
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
      this.requests = await RequestService.searchRequests(
        this.currentPage,
        this.searchQuery
      );
    } else {
      this.requests = await RequestService.getRequests(this.currentPage);
    }
  }

  async handleSearch(searchQuery: string): void {
    this.currentPage = 1;
    this.searchModeEnabled = true;
    this.searchQuery = searchQuery;
    this.requests = await RequestService.searchRequests(
      this.currentPage,
      searchQuery
    );
  }

  timeDiffToHuman(value: number): string {
    return intervalToLevels(value);
  }

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
  }

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
</script>
