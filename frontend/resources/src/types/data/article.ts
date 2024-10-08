export type Article =
  & ArticleIdentifier
  & ArticleData
  & ArticleExtra;

export type ArticleIdentifier = {
  id: string;
};

export type ArticleData = {
  title: string;
  type: string;
  tagKeys: string[];
  summary: string;
  content: string;
  thumbnailUrl: string;
};

export type ArticleExtra = {
  createdAt: string;
  updatedAt: string;
};

export type ArticleSearchQuery = {
  q?: string;
  type?: string;
  tag?: string[];
  sort: string;
  order: string;
  take: number;
  skip: number;
};
