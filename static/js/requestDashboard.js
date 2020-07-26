/*
    License: MIT
    Authors:
        - Josep Bigorra (averageflow)
 */
class RequestDashboard extends AbstractDashboard {
	/**
	 * @param {string} activeLink
	 * @param {string} activeSymbol
	 */
	constructor(activeLink, activeSymbol) {
		super('/goscope/requests', '/goscope/search/requests')
		document.getElementById(activeLink).className = 'active-navbar-link'
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

	fillTable(requestData) {
		let requestTable = document.getElementById('request-table')
		if (
			requestData === null ||
			requestData === undefined ||
			requestData.length === 0
		) {
			requestTable.innerHTML = `<h3>No results could be found</h3>`
			return
		}
		requestTable.innerHTML = this.requestTableHeaders
		requestData.forEach(function (item) {
			let requestMoment = item.time
			let elapsed = secondsToString(now - requestMoment)
			requestTable.innerHTML += `
            <tr class="text-center">\
			    <td class="p-3 custom-td">${applyStatusColor(item.response_status)}</td>
                <td class="p-3 custom-td">${applyMethodColor(item.method)}</td>
                <td class="monospaced p-3 custom-td">${item.path}</td>
                <td class="p-3 custom-td">${elapsed}</td>
                <td class="p-3 custom-td">
                    <a class="cursor-pointer" href="/goscope/requests/${
																					item.uid
																				}" target="_blank" rel="noopener noreferrer">
                        ${viewMoreImage}
                    </a>
                </td>
            </tr>
        `
		})
	}
}

document.addEventListener('DOMContentLoaded', async function () {
	let instance = new RequestDashboard('requests-link', 'requests-symbol')
	instance.requestTableHeaders = `
        <tr>
            <th class="custom-td">Status</th>
            <th class="custom-td">Verb</th>
            <th class="custom-td text-center">Path</th>
            <th class="custom-td">Happened</th>
            <th class="custom-td"></th>
        </tr>
    `
	let requestData = await instance.getEntries(instance.entryOffset)
	instance.fillTable(requestData)
	instance.commonEventListeners(instance)
})
