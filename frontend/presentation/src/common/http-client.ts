import axios from "axios";

import type { BackendResponse } from "frontend-types/app/backend-response";
import { BackendResponseStatus } from "frontend-types/app/backend-response";

export async function sendHttpRequest<Q, B, R>(
  baseUrl: string,
  path: string,
  query?: Q,
  body?: B,
): Promise<BackendResponse<R>> {
  try {
    const response = await axios({
      method: "POST",
      baseURL: baseUrl,
      url: path,
      params: query,
      paramsSerializer: { indexes: null },
      data: body,
    });

    return response.data;
  } catch (e) {
    console.log(e);
    return {
      status: BackendResponseStatus.UNKNOWN_ERROR,
    };
  }
}
