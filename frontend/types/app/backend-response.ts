export enum BackendResponseStatus {
  OK = "OK",
  PARSE_ERROR = "PARSE_ERROR",
  VALIDATION_FAILED = "VALIDATION_FAILED",
  AUTHENTICATION_FAILED = "AUTHENTICATION_FAILED",
  ENTITY_NOT_FOUND = "ENTITY_NOT_FOUND",
  UNKNOWN_ERROR = "UNKNOWN_ERROR",
} 

export type BackendResponse<D> =
| {
  status: BackendResponseStatus.OK,
  data: D,
}
| {
  status: Exclude<BackendResponseStatus, BackendResponseStatus.OK>,
  error?: string,
};
