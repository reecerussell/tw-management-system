import React, { useState, useEffect } from "react";
import { Token } from "../../utils/api";
import { Login } from "../../components/auth";
import { Login as SetLogin } from "../../utils/user";

const LoginContainer = () => {
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [error, setError] = useState(null);
	const [loading, setLoading] = useState(false);
	const [data, setData] = useState(null);

	const isValid = () => {
		if (username === "") {
			setError("Please enter your username!");
			return false;
		}

		if (password === "") {
			setError("Please enter your password!");
			return false;
		}

		return true;
	};

	const handleSubmit = async (e) => {
		e.preventDefault();

		if (!isValid() || loading) {
			return;
		}

		setLoading(true);

		await Token(
			{
				username,
				password,
			},
			async (res) => {
				setData(await res.json());
			},
			setError
		);

		setLoading(false);
	};

	const handleUpdateText = (e) => {
		switch (e.target.name) {
			case "username":
				setUsername(e.target.value);
				break;
			case "password":
				setPassword(e.target.value);
				break;
			default:
				console.error("unhandled input update");
				break;
		}
	};

	useEffect(() => {
		if (loading || !data) {
			return;
		}

		const { accessToken, expires } = data;
		SetLogin(accessToken, expires);
	}, [data, loading]);

	return (
		<Login
			username={username}
			password={password}
			error={error}
			loading={loading}
			handleSubmit={handleSubmit}
			handleUpdateText={handleUpdateText}
		/>
	);
};

export default LoginContainer;
