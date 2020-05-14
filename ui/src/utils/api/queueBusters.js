import Fetch from "./fetch";

const Get = async (department, onSuccess, onError) =>
	await Fetch("queuebuster/" + department, null, onSuccess, onError);

const GetAll = async (onSuccess, onError) =>
	await Fetch("queuebuster", null, onSuccess, onError);

const Create = async (data, onSuccess, onError) =>
	await Fetch(
		"queuebuster",
		{
			method: "POST",
			body: JSON.stringify(data),
		},
		onSuccess,
		onError
	);

const Enable = async (department, onSuccess, onError) =>
	await Fetch(
		"queuebuster/enable/" + department,
		{
			method: "POST",
		},
		onSuccess,
		onError
	);

const Disable = async (department, onSuccess, onError) =>
	await Fetch(
		"queuebuster/disable/" + department,
		{
			method: "POST",
		},
		onSuccess,
		onError
	);

const Delete = async (department, onSuccess, onError) =>
	await Fetch(
		"queuebuster/" + department,
		{
			method: "DELETE",
		},
		onSuccess,
		onError
	);

export { Get, GetAll, Create, Enable, Disable, Delete };
