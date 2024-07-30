import type { ConfigResolvedValue } from "frontend-resources";
import { ResponseStatus } from "frontend-resources";
import { sendHttpRequest } from "frontend-resources";
import type { ConfigValueResolver } from "../page-details-resolver";

export const tagValueResolver: ConfigValueResolver = {
  sourceMatcher: ["data", "tags"],
  cacheOption: {
    type: "resolve",
  },
  async resolve(context) {
    try {
      const response = await sendHttpRequest(
        context.apiBaseUrl,
        "/data/tag-list",
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
