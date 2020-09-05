<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <section>
      <h2>Request details</h2>
      <dl>
        <div v-if="hasContent(this.requestDetails.clientIP)">
          <dt>Client IP</dt>
          <dd>
            <code>{{ this.requestDetails.clientIP }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.headers)">
          <dt>Headers</dt>
          <dd class="code-block">
            <pre><code class="language-json">{{ this.requestDetails.headers }}</code></pre>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.host)">
          <dt>Host</dt>
          <dd>
            <code>{{ this.requestDetails.host }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.method)">
          <dt>Method</dt>
          <dd>
            <code>{{ this.requestDetails.method }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.path)">
          <dt>Path</dt>
          <dd>
            <code>{{ this.requestDetails.path }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.referrer)">
          <dt>Referrer</dt>
          <dd>
            <code>{{ this.requestDetails.referrer }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.computedRequestTime)">
          <dt>Time</dt>
          <dd>
            <code>{{ this.computedRequestTime }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.uid)">
          <dt>UID</dt>
          <dd>
            <code>{{ this.requestDetails.uid }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.url)">
          <dt>URL</dt>
          <dd>
            <code>{{ this.requestDetails.url }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.userAgent)">
          <dt>User Agent</dt>
          <dd>
            <code>{{ this.requestDetails.userAgent }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.body)">
          <dt>Body</dt>
          <dd>
            <pre><code>{{ this.requestDetails.body }}</code></pre>
          </dd>
        </div>
      </dl>

      <h1>Response details</h1>
      <dl>
        <div v-if="hasContent(this.responseDetails.clientIP)">
          <dt>Client IP</dt>
          <dd>
            <code>{{ this.responseDetails.clientIP }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.responseDetails.headers)">
          <dt>Headers</dt>
          <dd class="code-block">
            <pre><code class="language-json">{{ this.responseDetails.headers }}</code></pre>
          </dd>
        </div>
        <div v-if="hasContent(this.responseDetails.path)">
          <dt>Path</dt>
          <dd>
            <code>{{ this.responseDetails.path }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.responseDetails.size)">
          <dt>Size</dt>
          <dd>
            <code>{{ this.responseDetails.size }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.responseDetails.status)">
          <dt>Status</dt>
          <dd>
            <code>{{ this.responseDetails.status }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.computedResponseTime)">
          <dt>Time</dt>
          <dd>
            <code>{{ this.computedResponseTime }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.responseDetails.requestUID)">
          <dt>RequestUID</dt>
          <dd>
            <code>{{ this.responseDetails.requestUID }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.responseDetails.uid)">
          <dt>UID</dt>
          <dd>
            <code>{{ this.responseDetails.uid }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.responseDetails.body)">
          <dt>Body</dt>
          <dd class="code-block">
            <pre><code class="language-json">{{ this.responseDetails.body }}</code></pre>
          </dd>
        </div>
      </dl>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { DetailedRequest, DetailedResponse } from "@/interfaces/requests";
import { RequestService } from "@/api/requests";
import { epochToHumanDate } from "@/utils/time";
import { hasContent } from "@/utils/values";

export default Vue.extend({
  name: "RequestDetails",
  data() {
    return {
      requestDetails: {} as DetailedRequest,
      responseDetails: {} as DetailedResponse
    };
  },
  methods: {
    hasContent
  },
  computed: {
    computedRequestTime(): string {
      if (!this.requestDetails) {
        return "";
      }
      return epochToHumanDate(this.requestDetails.time);
    },
    computedResponseTime(): string {
      if (!this.responseDetails) {
        return "";
      }
      return epochToHumanDate(this.responseDetails.time);
    }
  },
  async created(): Promise<void> {
    const requestedData = await RequestService.getRequest(
      this.$route.params.id
    );
    this.$data.requestDetails = requestedData.data.request;
    this.$data.responseDetails = requestedData.data.response;
  }
});
</script>
