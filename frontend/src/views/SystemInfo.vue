<template>
  <div>
    <button v-on:click="$router.go(-1)">← Back</button>
    <br />
    <section v-if="systemInformation">
      <h2>Environment</h2>
      <details>
        <summary>Click to view variables...</summary>
        <table>
          <thead>
            <tr>
              <th>Variable</th>
              <th>Value</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(value, key) in systemInformation.environment"
              :key="key"
            >
              <td>
                <code>{{ key }}</code>
              </td>
              <td>
                <code>{{ value }}</code>
              </td>
            </tr>
          </tbody>
        </table>
      </details>
      <h2>CPU</h2>
      <dl>
        <dt>Core Count</dt>
        <dd>
          <code>{{ systemInformation.cpu.coreCount }}</code>
        </dd>
        <dt>Model</dt>
        <dd>
          <code>{{ systemInformation.cpu.modelName }}</code>
        </dd>
      </dl>

      <h2>Disk</h2>
      <dl>
        <dt>Free space</dt>
        <dd>
          <code>{{ systemInformation.disk.freeSpace }}</code>
        </dd>
        <dt>Partition Type</dt>
        <dd>
          <code>{{ systemInformation.disk.partitionType }}</code>
        </dd>
        <dt>Mount Path</dt>
        <dd>
          <code>{{ systemInformation.disk.mountPath }}</code>
        </dd>
        <dt>Total Space</dt>
        <dd>
          <code>{{ systemInformation.disk.totalSpace }}</code>
        </dd>
      </dl>

      <h2>Host</h2>
      <dl>
        <dt>Kernel Arch</dt>
        <dd>
          <code>{{ systemInformation.host.kernelArch }}</code>
        </dd>
        <dt>Kernel Version</dt>
        <dd>
          <code>{{ systemInformation.host.kernelVersion }}</code>
        </dd>
        <dt>Hostname</dt>
        <dd>
          <code>{{ systemInformation.host.hostname }}</code>
        </dd>
        <dt>Host OS</dt>
        <dd>
          <code>{{ systemInformation.host.hostOS }}</code>
        </dd>
        <dt>Platform</dt>
        <dd>
          <code>{{ systemInformation.host.hostPlatform }}</code>
        </dd>
        <dt>Uptime</dt>
        <dd>
          <code>{{ systemInformation.host.uptime }}</code>
        </dd>
      </dl>

      <h2>Memory</h2>
      <dl>
        <dt>Status</dt>
        <dd>
          <code
            >{{ systemInformation.memory.availableMemory }} available of
            {{ systemInformation.memory.totalMemory }}</code
          >
        </dd>
        <dt>Swap usage</dt>
        <dd>
          <code>{{ systemInformation.memory.usedSwap }}</code>
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
  data() {
    return {
      systemInformation: null as SystemInfoDetailsResponse | null
    };
  },
  async created(): Promise<void> {
    this.systemInformation = await SystemInfoService.getSystemInfo();
    document.title = `${this.systemInformation.applicationName} | System Information`;
  }
});
</script>
