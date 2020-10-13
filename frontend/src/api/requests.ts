import axios from "axios";
import {
  DetailedRequestResponse,
  RequestsEndpointResponse
} from "@/interfaces/requests";
import { Tag } from "@/interfaces/filter";

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
    filter: Tag[]
  ): Promise<RequestsEndpointResponse> {
    const url = process.env.VUE_APP_API_SEARCH_REQUESTS_URL;
    const offset: number = (page - 1) * 50;
    const stringFilter: string[] = filter.map(f => f.value);
    const response = await this.requestAxios.post<RequestsEndpointResponse>(
      url,
      {
        query,
        filter: {
          method: stringFilter,
          status: []
        }
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
