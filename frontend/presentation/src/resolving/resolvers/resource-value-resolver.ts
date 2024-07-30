import type { Resource, ResourceIdentifier } from "frontend-resources";
import { ResponseStatus } from "frontend-resources";
import { sendHttpRequest } from "frontend-resources";
import type { ConfigValueResolver } from "../page-details-resolver";

export const resourceValueResolver: ConfigValueResolver = {
  sourceMatcher: ["resource"],
  cacheOption: {
    type: "resolve",
  },
  async resolve(context, source) {
    try {
      const response = await sendHttpRequest<ResourceIdentifier, undefined, Resource>(
        context.apiBaseUrl,
        "/site/resource-get",
        { key: source[1] },
      );

      if (response.status !== ResponseStatus.OK) {
        console.log(JSON.stringify(response));
        return null;
      }

      return response.data.details;
    } catch (e) {
      console.log(e);
      return null;
    }
  },
};
