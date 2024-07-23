import type {
  PageDetails,
  StaticConfigResolvedValue,
  StaticConfigResolvingInfo,
  StaticConfigValue,
} from "frontend-types/site/data-resolving";

export type ResolvedTemplate = {
  template: string;
  staticConfig: Record<string, StaticConfigResolvedValue>;
};

export type ResolvedComponent = {
  component: string;
  staticConfig: Record<string, StaticConfigResolvedValue>;
};

export class StaticConfigResolver {
  private pageDetails: PageDetails;

  private resolvedTemplate: ResolvedTemplate;
  private resolvedComponentsBySlot: Record<string, ResolvedComponent[]>;

  public constructor(pageDetails: PageDetails) {
    this.pageDetails = pageDetails;

    this.resolvedTemplate = {
      template: "",
      staticConfig: {},
    };

    this.resolvedComponentsBySlot = {};
  }

  public async resolve() {
    await this.resolveTemplate();
    await this.resolveComponents();
  }

  public getTempalte(): ResolvedTemplate {
    return this.resolvedTemplate;
  }

  public getSlots(): string[] {
    return Object.keys(this.resolvedComponentsBySlot);
  }

  public getComponents(slot: string): ResolvedComponent[] {
    return this.resolvedComponentsBySlot[slot];
  }

  private async resolveTemplate() {
    this.resolvedTemplate = {
      template: this.pageDetails.template,
      staticConfig: await this.resolveConfig(this.pageDetails.staticConfig),
    };
  }

  private async resolveComponents() {
    this.resolvedComponentsBySlot = {};

    for (const [slotName, componentDetailsArray] of Object.entries(this.pageDetails.slots)) {
      const resolvedComponentDetailsArray: ResolvedComponent[] = [];

      for (const componentDetails of componentDetailsArray) {
        const resolvedComponentDetails: ResolvedComponent = {
          component: componentDetails.component,
          staticConfig: await this.resolveConfig(componentDetails.staticConfig),
        };

        resolvedComponentDetailsArray.push(resolvedComponentDetails);
      }

      this.resolvedComponentsBySlot[slotName] = resolvedComponentDetailsArray;
    }
  }

  private async resolveConfig(
    config: Record<string, StaticConfigValue>,
  ): Promise<Record<string, StaticConfigResolvedValue>> {
    const result: Record<string, StaticConfigResolvedValue> = {};
    for (const [configKey, configValue] of Object.entries(config)) {
      if (configValueIsResolvingInfo(configValue)) {
        result[configKey] = null;
      } else {
        result[configKey] = configValue;
      }
    }

    return result;
  }
}

function configValueIsResolvingInfo(c: StaticConfigValue): c is StaticConfigResolvingInfo {
  if (c === null) {
    return false;
  }

  if (typeof c !== "object") {
    return false;
  }

  return Object.keys(c).every((key) => ["source", "parameters", "defaultValue", "field"].includes(key));
}
