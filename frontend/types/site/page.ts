import type { PageDetails } from "./data-resolving";

export type Page = {
  title: string;
  description: string;
  urlPattern: string;
  details: PageDetails;
} & PageIdentifier;

export type PageIdentifier = {
  key: string;
};
