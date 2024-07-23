export type PageDetails = {
  template: string;
  staticConfig: Record<string, StaticConfigValue>;
  slots: Record<string, ComponentDetails[]>;
}

export type ComponentDetails = {
  component: string;
  staticConfig: Record<string, StaticConfigValue>;
}

export type StaticConfigValue =
| StaticConfigResolvingInfo
| StaticConfigResolvedValue;

export type StaticConfigResolvingInfo = {
  source: string[];
  parameters?: Record<string, StaticConfigValue>;
  defaultValue?: StaticConfigResolvedValue;
  field?: (string|number)[];
};

export type StaticConfigResolvedValue =
| string
| number
| boolean
| null
| object;

