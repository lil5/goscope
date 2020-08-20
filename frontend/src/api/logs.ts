import { DetailedLogsReponse, LogsEndpointResponse } from "@/interfaces/logs";
import axios from "axios";

export abstract class LogService {
  private static logsAxios = axios.create();

  static async getLogs(page: number): Promise<LogsEndpointResponse> {
    const url = "http://localhost:7005/goscope/api/logs";
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
    const url = "http://localhost:7005/goscope/api/search/logs";
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
    const url = `http://localhost:7005/goscope/api/logs/${uuid}`;
    const response = await this.logsAxios.get<DetailedLogsReponse>(url);
    return response.data;
  }
}
