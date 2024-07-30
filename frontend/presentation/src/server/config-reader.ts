import { readFileSync } from "node:fs";

export type Config = {
  apiBaseUrl: string;
  pageBaseUrl: string;
  port: number;
};

export function readConfig(): Config {
  const configFilePath = process.env.SERVICE_CONFIG_PATH!;

  const configFileContent = readFileSync(configFilePath, { encoding: "utf8" });

  return JSON.parse(configFileContent);
}
