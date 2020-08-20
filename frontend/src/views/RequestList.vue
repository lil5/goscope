<template>
  <section>
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
            <font-awesome-icon icon="eye" />
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

@Component
export default class RequestList extends Vue {
  private requests: RequestsEndpointResponse = {
    data: [],
    applicationName: ""
  };

  private now: number = Math.round(new Date().getTime() / 1000);

  async mounted(): Promise<void> {
    this.requests = await RequestService.getRequests(1);
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
