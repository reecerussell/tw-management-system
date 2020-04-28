import React, { useState, useEffect } from "react";
import * as Api from "../../utils/api";
import PropTypes from "prop-types";
import { ChangePassword } from "../../components/users";

const propTypes = {
	id: PropTypes.string.isRequired,
};
const defaultProps = {};

const ChangePasswordContainer = ({ id }) => {
	const [error, setError] = useState(null);
	const [success, setSuccess] = useState(null);
	const [loading, setLoading] = useState(false);
	const [current, setCurrent] = useState("");
	const [password, setPassword] = useState("");
	const [confirm, setConfirm] = useState("");

	const handleUpdateText = (e) => {
		const { name, value } = e.target;
		switch (name) {
			case "current":
				setCurrent(value);
				break;
			case "password":
				setPassword(value);
				break;
			case "confirm":
				setConfirm(value);
				break;
			default:
				console.error("unhandled update value");
				break;
		}
	};

	const isValid = () => {
		if (current === "") {
			setError("Current password is required.");
			return false;
		}

		if (password === "") {
			setError("Enter a new password.");
			return false;
		}

		if (confirm !== password) {
			setError("Passwords do not match.");
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

		await Api.Users.ChangePassword(
			{
				id,
				current,
				new: password,
			},
			async () => {
				setSuccess("Password successfully changed!");
			},
			setError
		);

		setLoading(false);
	};

	useEffect(() => {
		if (error !== null) {
			setSuccess(null);
		}

		if (success !== null) {
			setError(null);
			setTimeout(() => setSuccess(null), 3000);
		}
	}, [error, success]);

	return (
		<ChangePassword
			error={error}
			success={success}
			loading={loading}
			current={current}
			password={password}
			confirm={confirm}
			handleSubmit={handleSubmit}
			handleUpdateText={handleUpdateText}
		/>
	);
};

ChangePasswordContainer.propTypes = propTypes;
ChangePasswordContainer.defaultProps = defaultProps;

export default ChangePasswordContainer;
