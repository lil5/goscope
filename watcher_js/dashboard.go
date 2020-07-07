package watcher_js
const DashboardJs = `
let requestOffset = 0;
let responseOffset = 0;

const now = Math.round((new Date()).getTime() / 1000);
const viewMoreImage = '<svg style="width:1.2em;height:1.2em;" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 22 16"> <path fill="white" d="M16.56 13.66a8 8 0 0 1-11.32 0L.3 8.7a1 1 0 0 1 0-1.42l4.95-4.95a8 8 0 0 1 11.32 0l4.95 4.95a1 1 0 0 1 0 1.42l-4.95 4.95-.01.01zm-9.9-1.42a6 6 0 0 0 8.48 0L19.38 8l-4.24-4.24a6 6 0 0 0-8.48 0L2.4 8l4.25 4.24h.01zM10.9 12a4 4 0 1 1 0-8 4 4 0 0 1 0 8zm0-2a2 2 0 1 0 0-4 2 2 0 0 0 0 4z"> </path> </svg>';

async function getRequests(offset) {
    try {
        const response = await axios.get('/watcher/requests', {
            params: {
                "offset": offset,
            }
        });
        return response.data
    } catch (error) {
        console.error(error);
    }
}

async function getResponses(offset) {
    try {
        const response = await axios.get('/watcher/responses', {
            params: {
                "offset": offset,
            }
        });
        return response.data
    } catch (error) {
        console.error(error);
    }
}

String.prototype.toHumanDate = function () {
    let sec_num = parseInt(this, 10); // don't forget the second param
    let hours = Math.floor(sec_num / 3600);
    let hourString = hours + 'h';
    let minutes = Math.floor((sec_num - (hours * 3600)) / 60);
    let minuteString = minutes + 'm';
    let seconds = sec_num - (hours * 3600) - (minutes * 60);
    let secondString = seconds + 's';
    let resultingString = "";
    if (hourString !== "0h") {
        resultingString += hourString;
    }
    if (minuteString !== "0m") {
        resultingString += ' ' + minuteString;
    }
    if (secondString !== "0s") {
        resultingString += ' ' + secondString;
    }
    return resultingString + ' ago';
}

function fillRequestTable(requestData) {
    let requestTable = document.getElementById("request-table");
    requestTable.innerHTML = "";
    requestData.forEach(function (item) {
        let requestMoment = item.time;
        let elapsed = (now - requestMoment).toString().toHumanDate();
        requestTable.innerHTML += '\
            <tr class="text-center">\
            <td class="p-3 custom-td">' + item.method + '</td>\
            \<td class="p-3 custom-td">' + item.path + '</td>\
            \<td class="p-3 custom-td">' + elapsed + '</td>\
            <td class="p-3 custom-td">\
                 <a class="cursor-pointer" href="/watcher/responses/${item.uid}" target="_blank" rel="noopener noreferrer">' + viewMoreImage + '</a></td></tr>';
    })
}

function fillResponseTable(responseData) {
    let responseTable = document.getElementById("response-table");
    responseTable.innerHTML = "";
    responseData.forEach(function (item) {
        let requestMoment = item.time;
        let elapsed = (now - requestMoment).toString().toHumanDate();
        responseTable.innerHTML += '\
            <tr class="text-center"> \
            <td class="p-3 custom-td">' + item.status + '</td>\
            <td class="p-3 custom-td">' + item.client_ip + '</td>\
            <td class="p-3 custom-td">' + item.path + '</td>\
            <td class="p-3 custom-td">' + elapsed + '</td>\
            <td class="p-3 custom-td">\
                <a class="cursor-pointer" href="/watcher/responses/${item.uid}" target="_blank" rel="noopener noreferrer">' + viewMoreImage + '</a></td></tr>';
    })
}

document.addEventListener("DOMContentLoaded", async function () {
    let requestData = await getRequests(requestOffset);
    fillRequestTable(requestData)
    let responseData = await getResponses(responseOffset);
    fillResponseTable(responseData)
})

function increaseRequestOffset() {
    requestOffset += 10;
}

function decreaseRequestOffset() {
    if (requestOffset !== 0) {
        requestOffset -= 10;
    }
}

function increaseResponseOffset() {
    responseOffset += 10;
}

function decreaseResponseOffset() {
    if (responseOffset !== 0) {
        responseOffset -= 10;
    }
}

let prevRequestPage = document.getElementById("requests-prev-page");
prevRequestPage.onclick = async function () {
    decreaseRequestOffset();
    const data = await getRequests(requestOffset)
    if (data !== null && data.length > 0) {
        fillRequestTable(data)
    } else {
        increaseRequestOffset();
    }
}
let prevResponsePage = document.getElementById("responses-prev-page");
prevResponsePage.onclick = async function () {
    decreaseResponseOffset();
    const data = await getResponses(responseOffset)
    if (data !== null && data.length > 0) {
        fillResponseTable(data)
    } else {
        increaseResponseOffset();
    }
}
let nextRequestPage = document.getElementById("requests-next-page");
nextRequestPage.onclick = async function () {
    increaseRequestOffset();
    const data = await getRequests(requestOffset)
    if (data !== null && data.length > 0) {
        fillRequestTable(data)
    } else {
        decreaseRequestOffset()
    }
}
let nextResponsePage = document.getElementById("responses-next-page");
nextResponsePage.onclick = async function () {
    increaseResponseOffset();
    const data = await getResponses(responseOffset)
    if (data !== null && data.length > 0) {
        fillResponseTable(data)
    } else {
        decreaseResponseOffset()
    }
}
`