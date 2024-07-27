export type File =
  & FileIdentifier
  & FileData
  & FileExtra;

export type FileIdentifier = {
  id: string;
};

export type FileData = {
  title: string;
  mimeType: string;
  sizeInBytes: number;
};

export type FileExtra = {
  createdAt: string;
};

export type FileSearchParams = {
  q: string;
  mimeType: string;
  sort: string;
  order: string;
  take: number;
  skip: number;
};
