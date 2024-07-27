import type {
  ConfigResolvedValue,
  ConfigResolvingInfo,
  ConfigValue,
  PageDetails,
} from "frontend-resources";
import { PageDetailsResolvingCache } from "./page-details-resolving-cache";

export type ConfigValueResolver = {
  sourceMatcher: string[];
  cacheOption?: ConfigValueCacheOption;
  resolve(
    context: PageContext,
    source: string[],
    parameters?: Record<string, ConfigResolvedValue>,
  ): Promise<ConfigResolvedValue>;
};

export type ConfigValueCacheOption =
  & {
    parameterHash?(parameters: Record<string, ConfigResolvedValue>): string;
  }
  & (
    | {
      type: "time";
      timeToLive: number;
    }
    | {
      type: "resolve";
    }
  );

export type PageContext = {
  apiBaseUrl: string;
  pageBaseUrl: string;
  pageUrlPath: string;
  pageUrlParams: Record<string, string>;
  pageUrlQueries: Record<string, string | string[]>;
};

export type ResolvedPageDetails = {
  template: string;
  language: string;
  title: string;
  config: Record<string, ConfigResolvedValue>;
  slots: Record<string, ResolvedComponentDetails[]>;
};

export type ResolvedComponentDetails = {
  component: string;
  config: Record<string, ConfigResolvedValue>;
};

// eslint-disable-next-line @typescript-eslint/no-empty-object-type
interface ConfigValueResolverLookup
  extends Record<string, ConfigValueResolver | ConfigValueResolverLookup> {}

function isConfigValueResolver(
  l: ConfigValueResolver | ConfigValueResolverLookup,
): l is ConfigValueResolver {
  return (l as ConfigValueResolver).sourceMatcher
    && (l as ConfigValueResolver).resolve
    && Object.keys(l).every((key) => ["sourceMatcher", "resolve", "cacheOption"].includes(key));
}

function isConfigResolvingInfo(c: ConfigValue): c is ConfigResolvingInfo {
  return c !== null
    && typeof c === "object"
    && (c as ConfigResolvingInfo).source
    && Object.keys(c).every((key) => ["source", "parameters", "defaultValue", "field"].includes(key));
}

export class PageDetailsResolver {
  private resolverLookup: ConfigValueResolverLookup;
  private cache: PageDetailsResolvingCache;

  public constructor() {
    this.resolverLookup = {};
    this.cache = new PageDetailsResolvingCache();
  }

  public with(resolver: ConfigValueResolver): this {
    let currentLookup: ConfigValueResolverLookup = this.resolverLookup;

    for (let i = 0; i < resolver.sourceMatcher.length - 1; i++) {
      const lookupKey = resolver.sourceMatcher[i];
      if (!currentLookup[lookupKey]) {
        currentLookup[lookupKey] = {};
      }
      if (isConfigValueResolver(currentLookup[lookupKey])) {
        currentLookup[lookupKey] = {};
      }
      currentLookup = currentLookup[lookupKey];
    }

    currentLookup[resolver.sourceMatcher[resolver.sourceMatcher.length - 1]] = resolver;

    return this;
  }

  public async resolve(pageContext: PageContext, pageDetails: PageDetails): Promise<ResolvedPageDetails> {
    const resolvedLanguageAndTitle = await this.resolveConfig(
      pageContext,
      {
        language: pageDetails.language,
        title: pageDetails.title,
      },
    );
    const resolvedConfig = await this.resolveConfig(pageContext, pageDetails.config);

    const resolvedSlots: Record<string, ResolvedComponentDetails[]> = {};
    for (const [slotName, componentDetailsArray] of Object.entries(pageDetails.slots)) {
      const resolvedComponentDetailsArray: ResolvedComponentDetails[] = [];

      for (const componentDetails of componentDetailsArray) {
        const resolvedComponentDetails = {
          component: componentDetails.component,
          config: await this.resolveConfig(pageContext, componentDetails.config),
        };

        resolvedComponentDetailsArray.push(resolvedComponentDetails);
      }

      resolvedSlots[slotName] = resolvedComponentDetailsArray;
    }

    this.cache.clearCacheValuesAfterResolving();

    return {
      template: pageDetails.template,
      language: resolvedLanguageAndTitle.language as string,
      title: resolvedLanguageAndTitle.title as string,
      config: resolvedConfig,
      slots: resolvedSlots,
    };
  }

  private async resolveConfig(
    pageContext: PageContext,
    config: Record<string, ConfigValue>,
  ): Promise<Record<string, ConfigResolvedValue>> {
    const result: Record<string, ConfigResolvedValue> = {};
    for (const [configKey, configValue] of Object.entries(config)) {
      if (isConfigResolvingInfo(configValue)) {
        result[configKey] = await this.resolveConfigUsingResolver(pageContext, configValue);
      } else {
        result[configKey] = configValue;
      }
    }

    return result;
  }

  private async resolveConfigUsingResolver(
    pageContext: PageContext,
    resolvingInfo: ConfigResolvingInfo,
  ): Promise<ConfigResolvedValue> {
    let currentLookup: ConfigValueResolver | ConfigValueResolverLookup = this.resolverLookup;

    for (const lookupKey of resolvingInfo.source) {
      if (!this.resolverLookup[lookupKey]) {
        return null;
      }

      currentLookup = this.resolverLookup[lookupKey];
      if (isConfigValueResolver(currentLookup)) {
        break;
      }
    }

    const resolver = currentLookup as ConfigValueResolver;

    let resolvedParameters: Record<string, ConfigResolvedValue> | undefined;
    if (resolvingInfo.parameters) {
      resolvedParameters = await this.resolveConfig(pageContext, resolvingInfo.parameters);
    }

    let resolvedValue: ConfigResolvedValue | undefined;

    const cacheResult = this.cache.getCachedValue(resolvingInfo, resolver, resolvedParameters);
    if (cacheResult.hit) {
      resolvedValue = cacheResult.value;
    } else {
      try {
        resolvedValue = await resolver.resolve(pageContext, resolvingInfo.source, resolvedParameters);
        this.cache.addCachedValue(resolvingInfo, resolver, resolvedParameters, resolvedValue);
      } catch {
        resolvedValue = undefined;
      }
    }

    if (resolvingInfo.field) {
      for (const fieldKey of resolvingInfo.field) {
        /* eslint-disable @typescript-eslint/no-explicit-any */
        if (resolvedValue ?? ((resolvedValue as any)[fieldKey])) {
          resolvedValue = (resolvedValue as any)[fieldKey];
        } else {
          return resolvingInfo.defaultValue ?? null;
        }
        /* eslint-enable @typescript-eslint/no-explicit-any */
      }
    }

    return resolvedValue ?? resolvingInfo.defaultValue ?? null;
  }
}
