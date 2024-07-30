import type {
  ConfigResolvedValue,
  ConfigResolvingInfo,
} from "frontend-resources";
import type { ConfigValueResolver } from "./page-details-resolver";

export type CacheResult =
  | {
    hit: true;
    value: ConfigResolvedValue;
  }
  | {
    hit: false;
  };

type CacheEntry =
  & {
    value: ConfigResolvedValue;
  }
  & (
    | {
      type: "time";
      expiredAt: number;
    }
    | {
      type: "resolve";
    }
  );

export class PageDetailsResolvingCache {
  private cachedValue: Record<string, CacheEntry>;

  public constructor() {
    this.cachedValue = {};
  }

  public clearAllCache() {
    for (const key of Object.keys(this.cachedValue)) {
      delete this.cachedValue[key];
    }
  }

  public clearCacheValuesAfterResolving() {
    for (const [key, entry] of Object.entries(this.cachedValue)) {
      if (entry.type === "resolve") {
        delete this.cachedValue[key];
      }

      if (entry.type === "time" && this.getCurrentTimeInSecs() > entry.expiredAt) {
        delete this.cachedValue[key];
      }
    }
  }

  public addCachedValue(
    resolvingInfo: ConfigResolvingInfo,
    resolver: ConfigValueResolver,
    resolvedParameters: Record<string, ConfigResolvedValue> | undefined,
    value: ConfigResolvedValue,
  ) {
    if (!resolver.cacheOption) {
      return;
    }

    const cachePath = [...resolvingInfo.source];
    if (resolvedParameters && resolver.cacheOption.parameterHash) {
      cachePath.push(resolver.cacheOption.parameterHash(resolvedParameters));
    }

    const cacheKey = cachePath.join("/");

    switch (resolver.cacheOption.type) {
      case "time":
        this.cachedValue[cacheKey] = {
          type: "time",
          expiredAt: this.getCurrentTimeInSecs() + resolver.cacheOption.timeToLive,
          value,
        };
        break;
      case "resolve":
        this.cachedValue[cacheKey] = {
          type: "resolve",
          value,
        };
        break;
    }
  }

  public getCachedValue(
    resolvingInfo: ConfigResolvingInfo,
    resolver: ConfigValueResolver,
    resolvedParameters?: Record<string, ConfigResolvedValue> | undefined,
  ): CacheResult {
    if (!resolver.cacheOption) {
      return { hit: false };
    }

    const cachePath = [...resolvingInfo.source];
    if (resolvedParameters && resolver.cacheOption.parameterHash) {
      cachePath.push(resolver.cacheOption.parameterHash(resolvedParameters));
    }

    const cacheKey = cachePath.join("/");
    const cacheEntry = this.cachedValue[cacheKey];

    if (!cacheEntry) {
      return { hit: false };
    }

    switch (cacheEntry.type) {
      case "time":
        if (this.getCurrentTimeInSecs() > cacheEntry.expiredAt) {
          return { hit: true, value: cacheEntry.value };
        } else {
          return { hit: false };
        }
      case "resolve":
        return { hit: true, value: cacheEntry.value };
      default:
        return { hit: false };
    }
  }

  private getCurrentTimeInSecs() {
    return Math.floor(Date.now() / 1000);
  }
}
