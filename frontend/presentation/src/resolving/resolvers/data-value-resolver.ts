import type { ConfigResolvedValue } from "frontend-types/site/data-resolving";
import { BackendResponseStatus } from "frontend-types/app/backend-response";
import type { ConfigValueResolver } from "../page-details-resolver";
import { sendHttpRequest } from "../../common/http-client";

export const dataValueResolver: ConfigValueResolver = {
  sourceMatcher: ["data"],
  cacheOption: {
    type: "resolve",
    parameterHash(parameters) {
      return JSON.stringify(parameters);
    },
  },
  async resolve(context, source, parameters) {
    const response = await sendHttpRequest(
      context.apiBaseUrl,
      `/data/${source[1]}`,
      parameters,
    );

    if (response.status !== BackendResponseStatus.OK) {
      return null;
    }

    return response.data as ConfigResolvedValue;
  },
};
