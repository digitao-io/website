import type { ConfigResolvedValue } from "frontend-resources";
import { ResponseStatus } from "frontend-resources";
import { sendHttpRequest } from "frontend-resources";
import type { ConfigValueResolver } from "../page-details-resolver";

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

    if (response.status !== ResponseStatus.OK) {
      return null;
    }

    return response.data as ConfigResolvedValue;
  },
};
