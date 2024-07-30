import type { ConfigResolvedValue } from "frontend-resources";
import { ResponseStatus } from "frontend-resources";
import { sendHttpRequest } from "frontend-resources";
import type { ConfigValueResolver } from "../page-details-resolver";

export const articleListValueResolver: ConfigValueResolver = {
  sourceMatcher: ["data", "articles"],
  cacheOption: {
    type: "resolve",
    parameterHash(parameters) {
      return JSON.stringify(parameters);
    },
  },
  async resolve(context, _, parameters) {
    try {
      const response = await sendHttpRequest(
        context.apiBaseUrl,
        "/data/article-list",
        parameters,
      );

      if (response.status !== ResponseStatus.OK) {
        console.log(JSON.stringify(response));
        return null;
      }

      return response.data as ConfigResolvedValue;
    } catch (e) {
      console.log(e);
      return null;
    }
  },
};
