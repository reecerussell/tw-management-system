import React, { useState, useEffect } from "react";
import * as Api from "../../utils/api";
import { Edit } from "../../components/users";
import PropTypes from "prop-types";

const propTypes = {
	id: PropTypes.string.isRequired,
};
const defaultProps = {};

const EditContainer = ({ id }) => {
	const [user, setUser] = useState([]);
	const [error, setError] = useState(null);
	const [loading, setLoading] = useState(false);
	const [deleteModal, setDeleteModalOpen] = useState(false);
	const [redirect, setRedirect] = useState(null);

	const toggleModal = () => setDeleteModalOpen(!deleteModal);

	const handleUpdateText = (e) => {
		const data = { ...user };
		const { name, value } = e.target;
		data[name] = value;
		setUser(data);
	};

	const handleSubmit = async (e) => {
		e.preventDefault();

		if (loading) {
			return;
		}

		setLoading(true);

		await Api.Users.Update(
			user,
			async (res) => setUser(await res.json()),
			setError
		);

		setLoading(false);
	};

	const deleteSuccess = async () => {
		toggleModal();
		setRedirect("/users");
	};

	useEffect(() => {
		Api.Users.Get(id, async (res) => setUser(await res.json()), setError);
	}, []);

	if (!user) {
		return null;
	}

	return (
		<Edit
			user={user}
			error={error}
			loading={loading}
			handleUpdateText={handleUpdateText}
			handleSubmit={handleSubmit}
			toggleModal={toggleModal}
			deleteModal={deleteModal}
			deleteSuccess={deleteSuccess}
			redirect={redirect}
		/>
	);
};

EditContainer.propTypes = propTypes;
EditContainer.defaultProps = defaultProps;

export default EditContainer;
