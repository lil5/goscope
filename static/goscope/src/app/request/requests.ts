export interface Requests {
  method: string
  path: string
  time: number
  uid: string
  responseStatus: number
}

export interface RequestsEndpointResponse {
  applicationName: string
  data: Requests[]
}
