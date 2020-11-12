export interface Requests {
  method: string;
  path: string;
  time: number;
  uid: string;
  responseStatus: number;
}

export interface RequestsEndpointResponse {
  applicationName: string;
  data: Requests[];
}

export interface DetailedResponse {
  body: string;
  clientIP: string;
  headers: string;
  path: string;
  size: number;
  status: string;
  time: number;
  requestUID: string;
  uid: string;
}

export interface DetailedRequest {
  body: string;
  clientIP: string;
  headers: string;
  host: string;
  method: string;
  path: string;
  referrer: string;
  time: number;
  uid: string;
  url: string;
  userAgent: string;
}

export interface DetailedRequestResponse {
  applicationName: string;
  data: {
    response: DetailedResponse;
    request: DetailedRequest;
  };
}

export interface FilterRequest {
  method: Method[];
  status: Status[];
}

export enum Method {
  GET = "GET",
  HEAD = "HEAD",
  POST = "POST",
  PUT = "PUT",
  DELETE = "DELETE",
  CONNECT = "CONNECT",
  OPTIONS = "OPTIONS",
  TRACE = "TRACE",
  PATCH = "PATCH"
}

export enum Status {
  INFORMATIONAL = 1,
  SUCCESS = 2,
  REDIRECTION = 3,
  CLIENT_ERROR = 4,
  SERVER_ERROR = 5
}
