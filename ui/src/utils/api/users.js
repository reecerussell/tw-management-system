import Fetch from "./fetch";

const All = async (onSuccess, onError) =>
	await Fetch("users", null, onSuccess, onError);

const Get = async (id, onSuccess, onError) =>
	await Fetch("users/" + id, null, onSuccess, onError);

const Create = async (data, onSuccess, onError) =>
	await Fetch(
		"users",
		{
			method: "POST",
			body: JSON.stringify(data),
		},
		onSuccess,
		onError
	);

const Update = async (data, onSuccess, onError) =>
	await Fetch(
		"users",
		{
			method: "PUT",
			body: JSON.stringify(data),
		},
		onSuccess,
		onError
	);

const Delete = async (id, onSuccess, onError) =>
	await Fetch(
		"users/" + id,
		{
			method: "DELETE",
		},
		onSuccess,
		onError
	);

const ChangePassword = async (data, onSuccess, onError) =>
	await Fetch(
		"users/password",
		{
			method: "POST",
			body: JSON.stringify(data),
		},
		onSuccess,
		onError
	);

export { All, Get, Create, Update, Delete, ChangePassword };
