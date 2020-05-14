import Fetch from "./fetch";

const Token = async (creds, onSuccess, onError) =>
	await Fetch(
		"token",
		{
			method: "POST",
			body: JSON.stringify(creds),
		},
		onSuccess,
		onError
	);

export { Token };
