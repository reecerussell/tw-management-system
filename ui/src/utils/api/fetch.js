import { GetAccessToken } from "../user";

const Send = async (url, options) => {
	if (!options) {
		options = {
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + GetAccessToken(),
			},
		};
	} else if (!options.headers) {
		options.headers = {
			"Content-Type": "application/json",
			Authorization: "Bearer " + GetAccessToken(),
		};
	} else if (!options.headers["Authorization"]) {
		options.headers["Authorization"] = "Bearer " + GetAccessToken();
	}

	return await fetch(url, options);
};

const defaultFail = (err) => console.error(err);

const BaseUrl = document.getElementById("base-api-url").value;

const Fetch = async (
	url,
	options,
	onSuccess,
	onFail = defaultFail,
	baseUrl = BaseUrl
) => {
	try {
		const res = await Send(baseUrl + url, options);

		switch (res.status) {
			case 200:
			case 201:
				if (onSuccess) {
					await onSuccess(res);
				}
				break;
			case 401:
			case 403:
				window.location.hash = "/login";
				break;
			default:
				try {
					const { error } = await res.json();
					onFail(error);
				} catch {
					onFail("An error occured while reading the response.");
				}
				break;
		}
	} catch (e) {
		console.log(e);
		onFail(
			"It seems like we can't connect to the server. Try again later!"
		);
	}
};

export default Fetch;
export { Fetch, BaseUrl };
