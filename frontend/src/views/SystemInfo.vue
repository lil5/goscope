<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <br />
    <section>
      <h1>CPU</h1>
      <dl>
        <dt>Core Count</dt>
        <dd>
          <code>{{ this.computedDetails.cpu.coreCount }}</code>
        </dd>
        <dt>Model</dt>
        <dd>
          <code>{{ this.computedDetails.cpu.modelName }}</code>
        </dd>
      </dl>

      <h1>Disk</h1>
      <dl>
        <dt>Free space</dt>
        <dd>
          <code>{{ this.computedDetails.disk.freeSpace }}</code>
        </dd>
        <dt>Partition Type</dt>
        <dd>
          <code>{{ this.computedDetails.disk.partitionType }}</code>
        </dd>
        <dt>Mount Path</dt>
        <dd>
          <code>{{ this.computedDetails.disk.mountPath }}</code>
        </dd>
        <dt>Total Space</dt>
        <dd>
          <code>{{ this.computedDetails.disk.totalSpace }}</code>
        </dd>
      </dl>

      <h1>Host</h1>
      <dl>
        <dt>Kernel Arch</dt>
        <dd>
          <code>{{ this.computedDetails.host.kernelArch }}</code>
        </dd>
        <dt>Kernel Version</dt>
        <dd>
          <code>{{ this.computedDetails.host.kernelVersion }}</code>
        </dd>
        <dt>Hostname</dt>
        <dd>
          <code>{{ this.computedDetails.host.hostname }}</code>
        </dd>
        <dt>Host OS</dt>
        <dd>
          <code>{{ this.computedDetails.host.hostOS }}</code>
        </dd>
        <dt>Platform</dt>
        <dd>
          <code>{{ this.computedDetails.host.hostPlatform }}</code>
        </dd>
        <dt>Uptime</dt>
        <dd>
          <code>{{ this.computedDetails.host.uptime }}</code>
        </dd>
      </dl>

      <h1>Memory</h1>
      <dl>
        <dt>Status</dt>
        <dd>
          <code
            >{{ this.computedDetails.memory.availableMemory }} available of
            {{ this.computedDetails.memory.totalMemory }}</code
          >
        </dd>
        <dt>Swap usage</dt>
        <dd>
          <code>{{ this.computedDetails.memory.usedSwap }}</code>
        </dd>
      </dl>
    </section>
  </div>
</template>

<script lang="ts">
import { SystemInfoDetailsResponse } from "@/interfaces/system-info";
import { SystemInfoService } from "@/api/system-info";
import Vue from "vue";

export default Vue.extend({
  name: "SystemInfo",
  computed: {
    computedDetails(): SystemInfoDetailsResponse {
      if (
        !this.$data.systemInformation ||
        Object.keys(this.$data.systemInformation).length === 0
      ) {
        return {
          cpu: {
            coreCount: "",
            modelName: ""
          },
          disk: {
            freeSpace: "",
            partitionType: "",
            mountPath: "",
            totalSpace: ""
          },
          host: {
            kernelArch: "",
            kernelVersion: "",
            hostname: "",
            hostOS: "",
            hostPlatform: "",
            uptime: ""
          },
          memory: {
            availableMemory: "",
            totalMemory: "",
            usedSwap: ""
          }
        } as SystemInfoDetailsResponse;
      } else {
        return this.$data.systemInformation;
      }
    }
  },
  data() {
    return {
      systemInformation: {} as SystemInfoDetailsResponse
    };
  },
  async created(): Promise<void> {
    this.$data.systemInformation = await SystemInfoService.getSystemInfo();
  }
});
</script>
