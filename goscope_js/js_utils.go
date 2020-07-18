package goscope_js

func JsUtils() string {
	const script = `
const now = Math.round((new Date()).getTime() / 1000);
const viewMoreImage = '<svg class="eye-image" style="width:1.2em;height:1.2em;" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 22 16"> <path d="M16.56 13.66a8 8 0 0 1-11.32 0L.3 8.7a1 1 0 0 1 0-1.42l4.95-4.95a8 8 0 0 1 11.32 0l4.95 4.95a1 1 0 0 1 0 1.42l-4.95 4.95-.01.01zm-9.9-1.42a6 6 0 0 0 8.48 0L19.38 8l-4.24-4.24a6 6 0 0 0-8.48 0L2.4 8l4.25 4.24h.01zM10.9 12a4 4 0 1 1 0-8 4 4 0 0 1 0 8zm0-2a2 2 0 1 0 0-4 2 2 0 0 0 0 4z"> </path> </svg>';

const intervalToLevels = (interval, levels) => {
  const cbFun = (d, c) => {
    let bb = d[1] % c[0],
      aa = (d[1] - bb) / c[0];
    aa = aa > 0 ? aa + c[1] : '';

    return [d[0] + aa, bb];
  };

  let rslt = levels.scale.map((d, i, a) => a.slice(i).reduce((d, c) => d * c))
    .map((d, i) => ([d, levels.units[i]]))
    .reduce(cbFun, ['', interval]);
  return rslt[0];
};

const TimeLevels = {
  scale: [24, 60, 60, 1],
  units: ['d ', 'h ', 'm ', 's ']
};
const secondsToString = interval => intervalToLevels(interval, TimeLevels);

function applyMethodColor(method) {
    if (method === "GET") {
        return '<span class="badge-secondary">' + method + '</span>'
    } else if (method === "POST") {
        return '<span class="badge-info">' + method + '</span>'
    } else if (method === "PUT") {
        return '<span class="badge-info">' + method + '</span>'
    } else if (method === "PATCH") {
        return '<span class="badge-turq">' + method + '</span>'
    } else if (method === "DELETE") {
        return '<span class="badge-danger">' + method + '</span>'
    }
    return '<span class="badge-secondary">' + method + '</span>'
}

function applyStatusColor(status) {
    status = parseInt(status)
    if (status >= 200 && status < 300) {
        return '<span class="badge-success">' + status + '</span>'
    } else if (status >= 300 && status < 400) {
        return '<span class="badge-info">' + status + '</span>'
    } else if (status >= 400 && status < 500) {
        return '<span class="badge-warning">' + status + '</span>'
    }
    return '<span class="badge-danger">' + status + '</span>'
}
`
	return MinifyJs(script)
}
