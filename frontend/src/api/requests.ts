import axios from "axios";
import {
  DetailedRequestResponse,
  RequestsEndpointResponse
} from "@/interfaces/requests";

export abstract class RequestService {
  private static requestAxios = axios.create();

  static async getRequests(page: number): Promise<RequestsEndpointResponse> {
    const url = "http://localhost:7005/goscope/api/requests";
    const offset: number = (page - 1) * 50;
    const response = await this.requestAxios.get<RequestsEndpointResponse>(
      url,
      {
        params: {
          offset: offset.toString()
        }
      }
    );
    return response.data;
  }

  static async searchRequests(
    page: number,
    query: string
  ): Promise<RequestsEndpointResponse> {
    const url = "http://localhost:7005/goscope/api/search/requests";
    const offset: number = (page - 1) * 50;
    const response = await this.requestAxios.post<RequestsEndpointResponse>(
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

  static async getRequest(uuid: string) {
    const url = `http://localhost:7005/goscope/api/requests/${uuid}`;
    const response = await this.requestAxios.get<DetailedRequestResponse>(url);
    return response.data;
  }
}
