export function intervalToLevels(interval: number): string {
  const levels = {
    scale: [24, 60, 60, 1],
    units: ['d ', 'h ', 'm ', 's '],
  };
  const cbFun = (d, c) => {
    const bb = d[1] % c[0];
    let aa = (d[1] - bb) / c[0];
    aa = aa > 0 ? aa + c[1] : '';
    return [d[0] + aa, bb];
  };
  const rslt = levels.scale
    .map((d, i, a) => {
      return a.slice(i).reduce((e, f) => e * f);
    })
    .map((d, i) => [d, levels.units[i]])
    .reduce(cbFun, ['', interval]);
  return rslt[0];
}
