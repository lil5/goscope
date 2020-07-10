package goscope_css

func WatcherStyles() string {
	const styles = `
	body {
		background: #333;
		color: #fff;
		font-family: "Fira Mono", monospace;
	}

	.navbar-link {
		text-decoration: none;
		color: #fff;
	}
	.navbar-link:hover {
		text-decoration: underline;
		color: rgb(181, 224, 222);
	}
	.custom-td {
		background: #2B2B2D;
	}

	th {
		border-radius: 4px;
	}

	table {
		border-spacing: 2px;
		border-collapse: separate;
	}

	.paginate-button {
		color: #fff;
		border: none;
	}
	.tab-button {
		color: #fff;
		border: none;
		background: #2B2B2D;
		border-radius: 4px;
	}
	.badge-success {
		padding: 4px;
		border-radius: 4px;
		color: #fff;
		background-color: #1f9d55;
	}
	.badge-info {
		padding: 4px;
		border-radius: 4px;
		color: #fff;
		background-color: #1c3d5a;
	}
	.badge-secondary {
		border-radius: 4px;
		padding: 4px;
		color: #fff;
		background-color: #494444;
	}
	.badge-danger {
		border-radius: 4px;
		padding: 4px;
		color: #fff;
		background-color: #621b18;
	}
	.badge-warning {
		border-radius: 4px;
		padding: 4px;
		color: #fff;
		background-color: #684f1d;
	}
	`
	return MinifyCss(styles)
}
