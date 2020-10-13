export class EnumReflection {
  private static REGEXP = /^[0-9]+$/g;

  private static isString<T>(name: string): boolean {
    if (name.match(this.REGEXP)) return false;

    return true;
  }

  public static getNames<T>(object: T): Array<string> {
    const result = new Array<string>();

    for (const name in object) {
      if (this.isString(name)) result.push(name);
    }

    return result;
  }

  public static getValues<T>(object: T): Array<string | number> {
    const result = new Array<string | number>();

    for (const name in object) {
      if (this.isString(name)) result.push(object[name] as any);
    }

    return result;
  }
}
