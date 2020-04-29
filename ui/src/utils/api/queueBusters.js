import Fetch from "./fetch";

const Get = async (department, onSuccess, onError) =>
	await Fetch("queuebusters/" + department, null, onSuccess, onError);

const GetAll = async (onSuccess, onError) =>
	await Fetch("queuebusters", null, onSuccess, onError);

const Create = async (data, onSuccess, onError) =>
	await Fetch(
		"queuebusters",
		{
			method: "POST",
			body: JSON.stringify(data),
		},
		onSuccess,
		onError
	);

const Enable = async (department, onSuccess, onError) =>
	await Fetch(
		"queuebusters/enable/" + department,
		{
			method: "POST",
		},
		onSuccess,
		onError
	);

const Disable = async (department, onSuccess, onError) =>
	await Fetch(
		"queuebusters/disable/" + department,
		{
			method: "POST",
		},
		onSuccess,
		onError
	);

const Delete = async (department, onSuccess, onError) =>
	await Fetch(
		"queuebusters/" + department,
		{
			method: "DELETE",
		},
		onSuccess,
		onError
	);

export { Get, GetAll, Create, Enable, Disable, Delete };
