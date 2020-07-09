package goscope_js
const LogsJs = `
let logOffset = 0;

const logTableHeaders = '\
<thead>Logs</thead>\
<tr>\
	<th class="custom-td">Message</th>\
	<th class="custom-td">Time</th>\
	<th class="custom-td"></th>\
</tr>\
';

async function getLogs(offset) {
    try {
        const response = await axios.get('/goscope/log-records', {
            params: {
                "offset": offset,
            }
        });
        return response.data
    } catch (error) {
        console.error(error);
    }
}

function fillLogTable(logData) {
    let logTable = document.getElementById("log-table");
    logTable.innerHTML = logTableHeaders;
    logData.forEach(function (item) {
        let requestMoment = item.time;
        let elapsed = (now - requestMoment).toString().toHumanDate();
        logTable.innerHTML += '\
            <tr class="text-center">\
			<td class="p-3 custom-td">' + item.error + '...</td>\
			\<td class="p-3 custom-td">' + elapsed + '</td>\
            <td class="p-3 custom-td">\
                 <a class="cursor-pointer" href="/goscope/requests/' + item.uid + '" target="_blank" rel="noopener noreferrer">' + viewMoreImage + '</a></td></tr>';
    })
}

function increaseLogOffset() {
    logOffset += 10;
}

function decreaseLogOffset() {
    if (logOffset !== 0) {
        logOffset -= 10;
    }
}

let prevLogPage = document.getElementById("logs-prev-page");
prevLogPage.onclick = async function () {
    decreaseLogOffset();
    const data = await getRequests(logOffset)
    if (data !== null && data.length > 0) {
        fillRequestTable(data);
    } else {
        increaseLogOffset();
    }
}
let nextLogPage = document.getElementById("logs-next-page");
nextLogPage.onclick = async function () {
    increaseLogOffset();
    const data = await getRequests(logOffset);
    if (data !== null && data.length > 0) {
        fillRequestTable(data);
    } else {
        decreaseLogOffset();
    }
}

document.addEventListener("DOMContentLoaded", async function () {
    let logData = await getLogs(logOffset);
    fillLogTable(logData);
});
`