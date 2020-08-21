<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <br />
    <section>
      <h1>CPU</h1>
      <dl>
        <dt>Core Count</dt>
        <dd>
          <code>{{ this.systemInformation.cpu.coreCount }}</code>
        </dd>
        <dt>Model</dt>
        <dd>
          <code>{{ this.systemInformation.cpu.modelName }}</code>
        </dd>
      </dl>

      <h1>Disk</h1>
      <dl>
        <dt>Free space</dt>
        <dd>
          <code>{{ this.systemInformation.disk.freeSpace }}</code>
        </dd>
        <dt>Partition Type</dt>
        <dd>
          <code>{{ this.systemInformation.disk.partitionType }}</code>
        </dd>
        <dt>Mount Path</dt>
        <dd>
          <code>{{ this.systemInformation.disk.mountPath }}</code>
        </dd>
        <dt>Total Space</dt>
        <dd>
          <code>{{ this.systemInformation.disk.totalSpace }}</code>
        </dd>
      </dl>

      <h1>Host</h1>
      <dl>
        <dt>Kernel Arch</dt>
        <dd>
          <code>{{ this.systemInformation.host.kernelArch }}</code>
        </dd>
        <dt>Kernel Version</dt>
        <dd>
          <code>{{ this.systemInformation.host.kernelVersion }}</code>
        </dd>
        <dt>Hostname</dt>
        <dd>
          <code>{{ this.systemInformation.host.hostname }}</code>
        </dd>
        <dt>Host OS</dt>
        <dd>
          <code>{{ this.systemInformation.host.hostOS }}</code>
        </dd>
        <dt>Platform</dt>
        <dd>
          <code>{{ this.systemInformation.host.hostPlatform }}</code>
        </dd>
        <dt>Uptime</dt>
        <dd>
          <code>{{ this.systemInformation.host.uptime }}</code>
        </dd>
      </dl>

      <h1>Memory</h1>
      <dl>
        <dt>Status</dt>
        <dd>
          <code
            >{{ this.systemInformation.memory.availableMemory }} available of
            {{ this.systemInformation.memory.totalMemory }}</code
          >
        </dd>
        <dt>Swap usage</dt>
        <dd>
          <code>{{ this.systemInformation.memory.usedSwap }}</code>
        </dd>
      </dl>
    </section>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { SystemInfoDetailsResponse } from "@/interfaces/system-info";
import { SystemInfoService } from "@/api/system-info";

@Component
export default class SystemInfo extends Vue {
  private routerInstance;
  private systemInformation: SystemInfoDetailsResponse = {
    applicationName: "",
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
  };

  async mounted(): Promise<void> {
    this.systemInformation = await SystemInfoService.getSystemInfo();
  }
}
</script>

<style scoped></style>
