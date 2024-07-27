import axios from "axios";

import { ResponseStatus, type Response } from "../types/app/response";

export async function sendHttpRequest<Q, B, R>(
  baseUrl: string,
  path: string,
  query?: Q,
  body?: B,
): Promise<Response<R>> {
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
      status: ResponseStatus.UNKNOWN_ERROR,
    };
  }
}
