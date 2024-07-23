import axios from "axios";

import type { BackendResponse } from "frontend-types/app/backend-response";

export async function sendHttpRequest<Q, B, R>(
  path: string,
  query?: Q,
  body?: B,
): Promise<BackendResponse<R>> {
  const response = await axios({
    method: "POST",
    baseURL: "http://localhost:3000",
    url: path,
    params: query,
    data: body,
  });

  return response.data;
}
