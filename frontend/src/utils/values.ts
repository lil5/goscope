// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function hasContent(value: any): boolean {
  return !(!value || value === "");
}
