import Fetch from "./fetch";

const All = async (onError) => {
	let data = null;

	await Fetch(
		"users",
		null,
		async (res) => (data = await res.json()),
		onError
	);

	return data;
};

const Get = async (id, onError) => {
	let data = null;

	await Fetch(
		"users/" + id,
		null,
		async (res) => (data = await res.json()),
		onError
	);

	return data;
};

export { All, Get };
