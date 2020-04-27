import React, { useState } from "react";
import * as Api from "../../utils/api";
import { Delete } from "../../components/users";
import PropTypes from "prop-types";

const propTypes = {
	id: PropTypes.string.isRequired,
	onSuccess: PropTypes.func.isRequired,
	toggle: PropTypes.func.isRequired,
};
const defaultProps = {};

const DeleteContainer = ({ id, onSuccess, toggle }) => {
	const [error, setError] = useState(null);
	const [loading, setLoading] = useState(false);

	const handleSubmit = async (e) => {
		e.preventDefault();

		if (loading) {
			return;
		}

		setLoading(true);

		await Api.Users.Delete(id, onSuccess, setError);

		setLoading(false);
	};

	return (
		<Delete
			error={error}
			loading={loading}
			handleSubmit={handleSubmit}
			toggle={toggle}
		/>
	);
};

DeleteContainer.propTypes = propTypes;
DeleteContainer.defaultProps = defaultProps;

export default DeleteContainer;
