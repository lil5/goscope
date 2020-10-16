type EnumName<T> = Extract<keyof T, string>;
type EnumValue<T> = T[EnumName<T>];

export class EnumReflection {
  public static getNames<T>(object: T): EnumName<T>[] {
    const result: EnumName<T>[] = [];

    for (const name in object) {
      if (typeof name === "string") result.push(name);
    }

    return result;
  }

  public static getValues<T>(object: T): EnumValue<T>[] {
    const result: EnumValue<T>[] = [];

    for (const name in object) {
      if (typeof name === "string") result.push(object[name]);
    }

    return result;
  }
}
