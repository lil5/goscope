/*
    License: MIT
    Authors:
        - Josep Bigorra (averageflow)
 */
class LogsDashboard extends AbstractDashboard {
	/**
	 * @param {string} activeLink
	 * @param {string} activeSymbol
	 */
	constructor(activeLink, activeSymbol) {
		super('/goscope/log-records', '/goscope/search/logs')
		document.getElementById(activeLink).classList.add('active-navbar-link')
		document.getElementById(activeSymbol).style.fill =
			'var(--main-highlight-color)'
	}

	_requestTableHeaders = ''

	get requestTableHeaders() {
		return this._requestTableHeaders
	}

	/**
	 * @param {string} value
	 */
	set requestTableHeaders(value) {
		this._requestTableHeaders = value
	}

	fillTable(logData) {
		let logTable = document.getElementById('log-table')
		if (logData === null || logData === undefined || logData.length === 0) {
			logTable.innerHTML = `<h3>No results could be found</h3>`
			return
		}
		logTable.innerHTML = this.requestTableHeaders
		logData.forEach(function (item) {
			let requestMoment = item.time
			let elapsed = secondsToString(now - requestMoment)
			logTable.innerHTML += `
            <tr class="text-center">
			    <td class="monospaced p-3 custom-td">${item.error}</td>
			    <td class="p-3 custom-td">${elapsed}</td>
                <td class="p-3 custom-td">
                    <a class="cursor-pointer" href="/goscope/log-records/${item.uid}" target="_blank" rel="noopener noreferrer">
                        ${viewMoreImage}
                    </a>
                </td>
            </tr>`
		})
	}
}

document.addEventListener('DOMContentLoaded', async function () {
	let instance = new LogsDashboard('logs-link', 'logs-symbol')
	instance.requestTableHeaders = `
        <tr>
            <th class="custom-td">Message</th>
            <th class="custom-td">Time</th>
            <th class="custom-td"></th>
        </tr>
    `
	let requestData = await instance.getEntries(instance.entryOffset)
	instance.fillTable(requestData)
	instance.commonEventListeners(instance)
})
