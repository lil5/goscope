import axios from "axios";
import { ApplicationDetailsResponse } from "@/interfaces/app-details";

export abstract class ApplicationDetailsService {
  private static systemInfoAxios = axios.create();

  static async getApplicationDetails(): Promise<ApplicationDetailsResponse> {
    const url = process.env.VUE_APP_API_APP_DETAILS_URL;
    const response = await this.systemInfoAxios.get<ApplicationDetailsResponse>(
      url
    );
    return response.data;
  }
}
