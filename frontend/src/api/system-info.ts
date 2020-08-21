import { SystemInfoDetailsResponse } from "@/interfaces/system-info";
import axios from "axios";

export abstract class SystemInfoService {
  private static systemInfoAxios = axios.create();

  static async getSystemInfo(): Promise<SystemInfoDetailsResponse> {
    const url = process.env.VUE_APP_API_SYS_INFO_URL;
    const response = await this.systemInfoAxios.get<SystemInfoDetailsResponse>(
      url
    );
    return response.data;
  }
}
