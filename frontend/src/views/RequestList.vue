<template>
  <section>
    <SearchBar
      v-on:searchEvent="this.handleSearch"
      v-on:cancelSearchEvent="this.cancelSearch"
      :search-enabled="this.searchModeEnabled"
      :autocomplete="autocomplete"
    >
    </SearchBar>
    <table>
      <thead>
        <tr>
          <th style="width: 10%">Status</th>
          <th style="width: 15%">Method</th>
          <th style="width: 55%">Path</th>
          <th style="width: 15%">Happened</th>
          <th style="width: 5%;"></th>
        </tr>
      </thead>
      <tbody>
        <tr :key="index" v-for="(request, index) in requests.data">
          <td v-html="applyStatusColor(request.responseStatus)"></td>
          <td v-html="applyMethodColor(request.method)"></td>
          <td>{{ request.path }}</td>
          <td>
            <small>{{ timeDiffToHuman(now - request.time) }} ago</small>
          </td>
          <td>
            <router-link class="eye-link" :to="`/requests/${request.uid}`">
              <font-awesome-icon icon="eye" />
            </router-link>
          </td>
        </tr>
      </tbody>
    </table>
    <nav class="text-center">
      <button v-on:click="this.previousPage">← prev</button>
      <button v-on:click="this.nextPage">next →</button>
    </nav>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import SearchBar from "../components/SearchBar.vue";
import { RequestService } from "../api/requests";
import {
  RequestsEndpointResponse,
  Method,
  Status,
  FilterRequest
} from "../interfaces/requests";
import { Tag } from "../interfaces/filter";
import { EnumReflection } from "../utils/enum-reflection";
import { intervalToLevels } from "../utils/time";

function generateAutocomplete(): Tag[] {
  const autocomplete: Tag[] = [];

  const methods = EnumReflection.getNames(Method);
  methods.forEach(m => {
    autocomplete.push({
      text: "method:" + m.toLowerCase(),
      group: "method",
      value: m
    });
  });

  // TODO: add status filter
  // const statuses = EnumReflection.getNames(Status);
  // statuses.forEach(s => {
  //   autocomplete.push({
  //     text: `status:${Status[s]}xx`,
  //     group: "status",
  //     value: s
  //   });
  // });

  return autocomplete;
}

export default Vue.extend({
  name: "RequestList",
  components: { SearchBar },
  data() {
    return {
      requests: {} as RequestsEndpointResponse,
      currentPage: 1,
      searchModeEnabled: false,
      searchQuery: "",
      searchTags: [] as Tag[],
      autocomplete: generateAutocomplete(),
      now: Math.round(new Date().getTime() / 1000)
    };
  },
  async created(): Promise<void> {
    this.requests = await RequestService.getRequests(this.currentPage);
    document.title = `${this.requests.applicationName} | Requests`;
  },
  methods: {
    async nextPage(): Promise<void> {
      this.currentPage++;
      if (this.searchModeEnabled) {
        const filter = this.getFilter();
        const received = await RequestService.searchRequests(
          this.currentPage,
          this.searchQuery,
          filter
        );
        if (received.data && received.data.length > 0) {
          this.requests = received;
        } else {
          this.currentPage--;
        }
      } else {
        const received = await RequestService.getRequests(this.currentPage);
        if (received.data && received.data.length > 0) {
          this.requests = received;
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
        const filter = this.getFilter();
        this.requests = await RequestService.searchRequests(
          this.currentPage,
          this.searchQuery,
          filter
        );
      } else {
        this.requests = await RequestService.getRequests(this.currentPage);
      }
    },
    async handleSearch(searchQuery: string, searchTags: Tag[]): Promise<void> {
      this.currentPage = 1;
      this.searchModeEnabled = true;
      this.searchQuery = searchQuery;
      this.searchTags = searchTags;

      const filter = this.getFilter();
      this.requests = await RequestService.searchRequests(
        this.currentPage,
        searchQuery,
        filter
      );
    },
    async cancelSearch(): Promise<void> {
      this.currentPage = 1;
      this.searchModeEnabled = false;
      this.searchQuery = "";
      this.searchTags = [];
      this.requests = await RequestService.getRequests(this.currentPage);
    },
    timeDiffToHuman(value: number): string {
      return intervalToLevels(value);
    },
    getFilter(): FilterRequest {
      const status: Status[] = [];
      const method: Method[] = [];

      this.searchTags.forEach((tag: Tag) => {
        switch (tag.group) {
          case "method":
            //@ts-ignore
            method.push(tag.value);
            break;
          case "status":
            //@ts-ignore
            status.push(tag.value);
            break;
        }
      });

      return {
        status,
        method
      };
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
