package goscope_js

func DashboardJs() string {
	const script = `
let requestOffset = 0;

const requestTableHeaders = '\
<thead>Incoming Requests</thead>\
<tr>\
	<th class="custom-td">Status</th>\
	<th class="custom-td">Verb</th>\
	<th class="custom-td text-center">Path</th>\
	<th class="custom-td">Happened</th>\
	<th class="custom-td"></th>\
</tr>\
';

async function getRequests(offset) {
    try {
        const response = await axios.get('/goscope/requests', {
            params: {
                "offset": offset,
            }
        });
        return response.data;
    } catch (error) {
        console.error(error);
    }
}

function fillRequestTable(requestData) {
    let requestTable = document.getElementById("request-table");
    requestTable.innerHTML = requestTableHeaders;
    requestData.forEach(function (item) {
        let requestMoment = item.time;
        let elapsed = (now - requestMoment).toString().toHumanDate();
        requestTable.innerHTML += '\
            <tr class="text-center">\
			<td class="p-3 custom-td">' + applyStatusColor(item.response_status) + '</td>\
            <td class="p-3 custom-td">' + applyMethodColor(item.method) + '</td>\
            \<td class="p-3 custom-td">' + item.path + '</td>\
            \<td class="p-3 custom-td">' + elapsed + '</td>\
            <td class="p-3 custom-td">\
                 <a class="cursor-pointer" href="/goscope/requests/' + item.uid + '" target="_blank" rel="noopener noreferrer">' + viewMoreImage + '</a></td></tr>';
    })
}

function increaseRequestOffset() {
    requestOffset += 100;
}

function decreaseRequestOffset() {
    if (requestOffset !== 0) {
        requestOffset -= 100;
    }
}

let prevRequestPage = document.getElementById("requests-prev-page");
prevRequestPage.onclick = async function () {
    decreaseRequestOffset();
    const data = await getRequests(requestOffset)
    if (data !== null && data.length > 0) {
        fillRequestTable(data);
    } else {
        increaseRequestOffset();
    }
}
let nextRequestPage = document.getElementById("requests-next-page");
nextRequestPage.onclick = async function () {
    increaseRequestOffset();
    const data = await getRequests(requestOffset);
    if (data !== null && data.length > 0) {
        fillRequestTable(data);
    } else {
        decreaseRequestOffset();
    }
}

document.addEventListener("DOMContentLoaded", async function () {
    let requestData = await getRequests(requestOffset);
    fillRequestTable(requestData);
});
`
	return MinifyJs(script)
}
