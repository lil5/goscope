import { DetailedLogsReponse, LogsEndpointResponse } from "@/interfaces/logs";
import axios from "axios";

export abstract class LogService {
  private static logsAxios = axios.create();

  static async getLogs(page: number): Promise<LogsEndpointResponse> {
    const url = process.env.VUE_APP_API_LOGS_URL;
    const offset: number = (page - 1) * 50;
    const response = await this.logsAxios.get<LogsEndpointResponse>(url, {
      params: {
        offset: offset.toString()
      }
    });
    return response.data;
  }

  static async searchLogs(
    page: number,
    query: string
  ): Promise<LogsEndpointResponse> {
    const url = process.env.VUE_APP_API_SEARCH_LOGS_URL;
    const offset: number = (page - 1) * 50;
    const response = await this.logsAxios.post<LogsEndpointResponse>(
      url,
      { query },
      {
        params: {
          offset: offset.toString()
        }
      }
    );
    return response.data;
  }

  static async getLog(uuid: string) {
    const url = `${process.env.VUE_APP_API_LOGS_URL}/${uuid}`;
    const response = await this.logsAxios.get<DetailedLogsReponse>(url);
    return response.data;
  }
}
