import { SystemInfoDetailsResponse } from "@/interfaces/system-info";
import axios from "axios";

export abstract class SystemInfoService {
  private static systemInfoAxios = axios.create();

  static async getSystemInfo(): Promise<SystemInfoDetailsResponse> {
    const url = "http://localhost:7005/goscope/api/info";
    const response = await this.systemInfoAxios.get<SystemInfoDetailsResponse>(
      url
    );
    return response.data;
  }
}
