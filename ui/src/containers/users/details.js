import React, { useState, useEffect } from "react";
import * as Api from "../../utils/api";
import { Details } from "../../components/users";
import PropTypes from "prop-types";

const propTypes = {
	id: PropTypes.string.isRequired,
};
const defaultProps = {};

const DetailsContainer = ({ id }) => {
	const [user, setUser] = useState([]);
	const [error, setError] = useState(null);

	useEffect(() => {
		Api.Users.Get(id, async (res) => setUser(await res.json()), setError);
	}, [id]);

	return <Details user={user} error={error} />;
};

DetailsContainer.propTypes = propTypes;
DetailsContainer.defaultProps = defaultProps;

export default DetailsContainer;
