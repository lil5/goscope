import axios from "axios";
import {
  DetailedRequestResponse,
  FilterRequest,
  RequestsEndpointResponse
} from "../interfaces/requests";

export abstract class RequestService {
  private static requestAxios = axios.create();

  static async getRequests(page: number): Promise<RequestsEndpointResponse> {
    const url = process.env.VUE_APP_API_REQUESTS_URL;
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
    query: string,
    filter: FilterRequest
  ): Promise<RequestsEndpointResponse> {
    const url = process.env.VUE_APP_API_SEARCH_REQUESTS_URL;
    const offset: number = (page - 1) * 50;
    const response = await this.requestAxios.post<RequestsEndpointResponse>(
      url,
      {
        query,
        filter
      },
      {
        params: {
          offset: offset.toString()
        }
      }
    );
    return response.data;
  }

  static async getRequest(uuid: string) {
    const url = `${process.env.VUE_APP_API_REQUESTS_URL}/${uuid}`;
    const response = await this.requestAxios.get<DetailedRequestResponse>(url);
    return response.data;
  }
}
