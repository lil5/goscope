export interface Logs {
  error: string
  time: number
  uid: string
}

export interface LogsEndpointResponse {
  data: Logs[]
  applicationName: string
  entriesPerPage: number
}
