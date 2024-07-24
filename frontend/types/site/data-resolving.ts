export type PageDetails = {
  template: string;
  language: string;
  title: string;
  config: Record<string, ConfigValue>;
  slots: Record<string, ComponentDetails[]>;
}

export type ComponentDetails = {
  component: string;
  config: Record<string, ConfigValue>;
}

export type ConfigValue =
| ConfigResolvingInfo
| ConfigResolvedValue;

export type ConfigResolvingInfo = {
  source: string[];
  parameters?: Record<string, ConfigValue>;
  defaultValue?: ConfigResolvedValue;
  field?: (string|number)[];
};

export type ConfigResolvedValue =
| string
| number
| boolean
| null
| object;

