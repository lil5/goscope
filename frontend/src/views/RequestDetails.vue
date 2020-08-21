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
          <code class="language-json">{{ this.requestDetails.headers }}</code>
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
          <code>{{ this.requestDetails.body }}</code>
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
          <code class="language-json">{{ this.responseDetails.headers }}</code>
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
          <code class="language-json">{{ this.responseDetails.body }}</code>
        </dd>
      </dl>
    </section>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class RequestDetails extends Vue {}
</script>
