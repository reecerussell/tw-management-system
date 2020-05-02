import React, { useState } from "react";
import * as Api from "../../utils/api";
import { Create } from "../../components/users";

const CreateContainer = () => {
	const [error, setError] = useState(null);
	const [loading, setLoading] = useState(false);
	const [username, setUsername] = useState("");
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [confirm, setConfirm] = useState("");
	const [redirect, setRedirect] = useState(null);

	const isValid = () => {
		if (username === "") {
			setError("Username is required.");
			return;
		}

		if (email === "") {
			setError("Email is required.");
			return;
		}

		if (password === "") {
			setError("Password is required.");
			return;
		}

		if (password < 6) {
			setError("Password must be greater than 5 characters.");
			return;
		}

		if (confirm === "") {
			setError("Please confirm password.");
			return;
		}

		if (confirm !== password) {
			setError("Passwords do not match.");
			return;
		}

		return true;
	};

	const handleUpdateText = (e) => {
		const { name, value } = e.target;
		switch (name) {
			case "username":
				setUsername(value);
				break;
			case "email":
				setEmail(value);
				break;
			case "password":
				setPassword(value);
				break;
			case "confirm":
				setConfirm(value);
				break;
			default:
				console.error("unhandled input change");
				break;
		}
	};

	const handleSubmit = async (e) => {
		e.preventDefault();

		if (!isValid() || loading) {
			return;
		}

		setLoading(true);

		await Api.Users.Create(
			{
				username,
				email,
				password,
			},
			async (res) => {
				const { id } = await res.json();
				setRedirect(`/users/${id}/details`);
			},
			setError
		);

		setLoading(false);
	};

	return (
		<Create
			error={error}
			loading={loading}
			redirect={redirect}
			username={username}
			email={email}
			password={password}
			confirm={confirm}
			handleUpdateText={handleUpdateText}
			handleSubmit={handleSubmit}
		/>
	);
};

export default CreateContainer;
