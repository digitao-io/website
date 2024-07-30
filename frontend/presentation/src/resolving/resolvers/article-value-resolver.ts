import type { ConfigResolvedValue } from "frontend-resources";
import { ResponseStatus } from "frontend-resources";
import { sendHttpRequest } from "frontend-resources";
import type { ConfigValueResolver } from "../page-details-resolver";

export const articleValueResolver: ConfigValueResolver = {
  sourceMatcher: ["data", "article"],
  cacheOption: {
    type: "resolve",
  },
  async resolve(context, source) {
    try {
      const response = await sendHttpRequest(
        context.apiBaseUrl,
        "/data/article-get",
        { id: source[2] },
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
