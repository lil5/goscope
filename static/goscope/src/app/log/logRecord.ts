export interface LogRecord {
  error: string;
  time: number;
  uid: string;
}

export interface LogsEndpointResponse {
  data: LogRecord[];
  applicationName: string;
  entriesPerPage: number;
}

export interface DetailedLogsReponse {
  data: {
    logDetails: LogRecord
  };
  applicationName: string;
}
