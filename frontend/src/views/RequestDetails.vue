<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <section>
      <h1>Request details</h1>
      <dl>
        <div v-if="hasContent(this.requestDetails.clientIP)">
          <dt>Client IP</dt>
          <dd>
            <code>{{ this.requestDetails.clientIP }}</code>
          </dd>
        </div>
        <div v-if="hasContent(this.requestDetails.headers)">
          <dt>Headers</dt>
          <details>
            <summary>Click to view headers...</summary>
            <pre>
                <code class="language-json">{{ this.requestDetails.headers }}</code>
              </pre>
          </details>
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
          <details>
            <summary>Click to view body...</summary>
            <pre>
              <code class="language-json">{{ this.requestDetails.body }}</code>
            </pre>
          </details>
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
          <details>
            <summary>Click to view headers... </summary>
            <pre>
              <code class="language-json" >{{ this.responseDetails.headers }}</code>
            </pre>
          </details>
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
          <details>
            <summary>Click to view body...</summary>
            <pre>
                <code class="language-json">{{ this.responseDetails.body }}</code>
              </pre>
          </details>
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
import Prism from "prismjs";

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
  async mounted(): Promise<void> {
    const requestedData = await RequestService.getRequest(
      this.$route.params.id
    );
    this.requestDetails = requestedData.data.request;
    this.responseDetails = requestedData.data.response;
    document.title = `${requestedData.applicationName} | Request ${this.$route.params.id}`;
    Prism.highlightAll();
  }
});
</script>
<style scoped>
pre {
  white-space: pre-wrap; /* css-3 */
  white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
  white-space: -pre-wrap; /* Opera 4-6 */
  white-space: -o-pre-wrap; /* Opera 7 */
  word-wrap: break-word; /* Internet Explorer 5.5+ */
}
</style>
