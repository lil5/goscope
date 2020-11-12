export interface SystemInfoDetailsResponse {
  applicationName: string;
  cpu: {
    coreCount: string;
    modelName: string;
  };
  disk: {
    freeSpace: string;
    partitionType: string;
    mountPath: string;
    totalSpace: string;
  };
  host: {
    kernelArch: string;
    kernelVersion: string;
    hostname: string;
    hostOS: string;
    hostPlatform: string;
    uptime: string;
  };
  memory: {
    availableMemory: string;
    totalMemory: string;
    usedSwap: string;
  };
  environment: Record<string, string>;
}
