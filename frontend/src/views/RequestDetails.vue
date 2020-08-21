<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <section>
      <h2>Request details</h2>
      <dl>
        <dt>Client IP</dt>
        <dd>
          <code>{{ this.requestDetails.clientIP }}</code>
        </dd>
        <dt>Headers</dt>
        <dd class="code-block">
          <pre><code class="language-json">{{ this.requestDetails.headers }}</code></pre>
        </dd>
        <dt>Host</dt>
        <dd>
          <code>{{ this.requestDetails.host }}</code>
        </dd>
        <dt>Method</dt>
        <dd>
          <code>{{ this.requestDetails.method }}</code>
        </dd>
        <dt>Path</dt>
        <dd>
          <code>{{ this.requestDetails.path }}</code>
        </dd>
        <dt>Referrer</dt>
        <dd>
          <code>{{ this.requestDetails.referrer }}</code>
        </dd>
        <dt>Time</dt>
        <dd>
          <code>{{ this.requestDetails.time }}</code>
        </dd>
        <dt>UID</dt>
        <dd>
          <code>{{ this.requestDetails.uid }}</code>
        </dd>
        <dt>URL</dt>
        <dd>
          <code>{{ this.requestDetails.url }}</code>
        </dd>
        <dt>User Agent</dt>
        <dd>
          <code>{{ this.requestDetails.userAgent }}</code>
        </dd>
        <dt>Body</dt>
        <dd>
          <pre><code>{{ this.requestDetails.body }}</code></pre>
        </dd>
      </dl>

      <h1>Response details</h1>
      <dl>
        <dt>Client IP</dt>
        <dd>
          <code>{{ this.responseDetails.clientIP }}</code>
        </dd>
        <dt>Headers</dt>
        <dd class="code-block">
          <pre><code class="language-json">{{ this.responseDetails.headers }}</code></pre>
        </dd>
        <dt>Path</dt>
        <dd>
          <code>{{ this.responseDetails.path }}</code>
        </dd>
        <dt>Size</dt>
        <dd>
          <code>{{ this.responseDetails.size }}</code>
        </dd>
        <dt>Status</dt>
        <dd>
          <code>{{ this.responseDetails.status }}</code>
        </dd>
        <dt>Time</dt>
        <dd>
          <code>{{ this.responseDetails.time }}</code>
        </dd>
        <dt>RequestUID</dt>
        <dd>
          <code>{{ this.responseDetails.requestUID }}</code>
        </dd>
        <dt>UID</dt>
        <dd>
          <code>{{ this.responseDetails.uid }}</code>
        </dd>
        <dt>Body</dt>
        <dd class="code-block">
          <pre><code class="language-json">{{ this.responseDetails.body }}</code></pre>
        </dd>
      </dl>
    </section>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { PropType } from "vue";
import { DetailedRequest, DetailedResponse } from "@/interfaces/requests";
import { RequestService } from "@/api/requests";

@Component({
  props: {
    requestUUID: String as PropType<string>
  }
})
export default class RequestDetails extends Vue {
  private requestDetails: DetailedRequest = {
    body: "",
    clientIP: "",
    headers: "",
    host: "",
    method: "",
    path: "",
    referrer: "",
    time: 0,
    uid: "",
    url: "",
    userAgent: ""
  };
  private responseDetails: DetailedResponse = {
    body: "",
    clientIP: "",
    headers: "",
    path: "",
    size: 0,
    status: "",
    time: 0,
    requestUID: "",
    uid: ""
  };

  async mounted(): Promise<void> {
    const requestedData = await RequestService.getRequest(
      this.$route.params.id
    );
    this.requestDetails = requestedData.data.request;
    this.responseDetails = requestedData.data.response;
  }
}
</script>