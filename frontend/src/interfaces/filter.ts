export interface Tag {
  text: string;
  style?: string;
  classes?: string;
  group: Group;
  value: string;
}

export type Group = "method" | "status";
