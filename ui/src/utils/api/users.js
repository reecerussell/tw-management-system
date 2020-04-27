import Fetch from "./fetch";

const All = async (onSuccess, onError) =>
	await Fetch("users", null, onSuccess, onError);

const Get = async (id, onSuccess, onError) =>
	await Fetch("users/" + id, null, onSuccess, onError);

export { All, Get };
