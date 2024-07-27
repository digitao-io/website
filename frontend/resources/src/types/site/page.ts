import type { PageDetails } from "./page-details";

export type Page = {
  title: string;
  description: string;
  urlPattern: string;
  details: PageDetails;
} & PageIdentifier;

export type PageIdentifier = {
  key: string;
};
