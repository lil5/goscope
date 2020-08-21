import axios from "axios";
import { ApplicationDetailsResponse } from "@/interfaces/app-details";

export abstract class ApplicationDetailsService {
  private static systemInfoAxios = axios.create();

  static async getApplicationDetails(): Promise<ApplicationDetailsResponse> {
    const url = "http://localhost:7005/goscope/api/application-name";
    const response = await this.systemInfoAxios.get<ApplicationDetailsResponse>(
      url
    );
    return response.data;
  }
}
