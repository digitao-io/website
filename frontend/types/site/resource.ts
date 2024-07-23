export type Resource = {
  title: string;
  description: string;
  type: string;
  details: ResourceDetails;
} & ResourceIdentifier;

export type ResourceIdentifier = {
  key: string;
};

export type ResourceDetails =
| string
| number
| boolean
| null
| object;
