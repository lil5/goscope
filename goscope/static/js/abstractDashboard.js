/*
    License: MIT
    Authors:
        - Josep Bigorra (averageflow)
 */

/**
 * Abstract class that encapsulates common functionalities to the GoScope dashboards.
 * You can extend this class by satisfying the constructor constraints and
 * implementing a fillTable method.
 */
class AbstractDashboard {
    /**
     * @param {string} entriesUrl
     * @param {string} searchUrl
     */
    constructor(entriesUrl, searchUrl) {
        this.entriesUrl = entriesUrl;
        this.searchUrl = searchUrl;
    }

    _entryOffset = 0;

    get entryOffset() {
        return this._entryOffset;
    }

    /**
     * @param {number} value
     */
    set entryOffset(value) {
        this._entryOffset = value;
    }

    _searchOffset = 0;

    get searchOffset() {
        return this._searchOffset;
    }

    /**
     * @param {number} value
     */
    set searchOffset(value) {
        this._searchOffset = value;
    }

    _searchModeEnabled = false;

    get searchModeEnabled() {
        return this._searchModeEnabled;
    }

    /**
     * @param {boolean} value
     */
    set searchModeEnabled(value) {
        this._searchModeEnabled = value;
    }

    get searchUrl() {
        return this._searchUrl;
    }

    /**
     * @param {string} value
     */
    set searchUrl(value) {
        this._searchUrl = value;
    }

    get entriesUrl() {
        return this._entriesUrl;
    }

    /**
     * @param {string} value
     */
    set entriesUrl(value) {
        this._entriesUrl = value;
    }

    /**
     * @param {number} offset
     */
    async getEntries(offset) {
        try {
            let config = {
                params: {
                    "offset": offset,
                }
            }
            const response = await axios.get(this.entriesUrl, config);
            return response.data;
        } catch (error) {
            console.error(error);
        }
    }

    /**
     * @param {number} offset
     * @param {string} searchString
     */
    async searchEntries(offset, searchString) {
        try {
            let config = {
                params: {
                    "offset": offset,
                }
            };
            let postBody = {
                query: searchString,
            };
            const response = await axios.post(this.searchUrl, postBody, config);
            return response.data;
        } catch (error) {
            console.error(error);
        }
    }

    increaseEntryOffset() {
        this.entryOffset += goscopeEntriesPerPage;
    }

    decreaseEntryOffset() {
        if (this.entryOffset !== 0) {
            this.entryOffset -= goscopeEntriesPerPage;
        }
    }

    increaseSearchOffset() {
        this.searchOffset += goscopeEntriesPerPage;
    }

    decreaseSearchOffset() {
        if (this.searchOffset !== 0) {
            this.searchOffset -= goscopeEntriesPerPage;
        }
    }

    async nextPage() {
        if (this.searchModeEnabled) {
            let searchString = document.getElementById("search-input").value;
            if (searchString === "" || searchString === null || searchString === undefined) {
                return
            }
            this.increaseSearchOffset();
            const data = await this.searchEntries(this.searchOffset, searchString);
            if (data !== null && data.length > 0) {
                this.fillTable(data);
            } else {
                this.decreaseSearchOffset();
            }
        } else {
            this.increaseEntryOffset();
            const data = await this.getEntries(this.entryOffset);
            if (data !== null && data.length > 0) {
                this.fillTable(data);
            } else {
                this.decreaseEntryOffset();
            }
        }
    }

    async previousPage() {
        if (this.searchModeEnabled) {
            let searchString = document.getElementById("search-input").value;
            if (searchString === "" || searchString === null || searchString === undefined) {
                return
            }
            this.decreaseSearchOffset();
            const data = await this.searchEntries(this.searchOffset, searchString);
            if (data !== null && data.length > 0) {
                this.fillTable(data);
            } else {
                this.increaseSearchOffset();
            }
        } else {
            this.decreaseEntryOffset();
            const data = await this.getEntries(this.entryOffset)
            if (data !== null && data.length > 0) {
                this.fillTable(data);
            } else {
                this.increaseEntryOffset();
            }
        }
    }

    async searchInputEnter(event) {
        if (event.key !== "Enter") {
            return;
        }
        document.getElementById("search-button").click();
        event.preventDefault();
    }

    async cancelSearch() {
        document.getElementById("search-input").value = "";
        this.searchModeEnabled = false;
        document.getElementById("search-cancel-button").style.display = "none";
        let requestData = await this.getEntries(this.entryOffset);
        this.fillTable(requestData);
    }

    async performSearch() {
        let searchString = document.getElementById("search-input").value;
        if (searchString === "" || searchString === null || searchString === undefined) {
            return
        }
        this.searchModeEnabled = true;
        let requestData = await this.searchEntries(this.searchOffset, searchString);
        this.fillTable(requestData);
        document.getElementById("search-cancel-button").style.display = "initial";
    }

    /**
     * @param {AbstractDashboard} instance
     */
    commonEventListeners(instance) {
        document.getElementById("prev-page").addEventListener("click", async () => {
            await instance.previousPage()
        });

        document.getElementById("next-page").addEventListener("click", async () => {
            await instance.nextPage()
        });

        document.getElementById("search-input").addEventListener("keyup", event => {
            instance.searchInputEnter(event)
        });

        document.getElementById("search-button").addEventListener("click", async () => {
            await instance.performSearch()
        });

        document.getElementById("search-cancel-button").addEventListener("click", async () => {
            await instance.cancelSearch()
        });
    }
}
