import React, { useState, useEffect } from "react";
import * as Api from "../../utils/api";
import PropTypes from "prop-types";
import { Details } from "../../components/queueBusters";

const propTypes = {
	department: PropTypes.string.isRequired,
};
const defaultProps = {};

const DetailsContainer = ({ department }) => {
	const [queueBuster, setQueueBuster] = useState(null);
	const [error, setError] = useState(null);
	const [loading, setLoading] = useState(false);

	const fetchQueueBuster = async () =>
		await Api.QueueBusters.Get(
			department,
			async (res) => setQueueBuster(await res.json()),
			setError
		);

	const handleToggle = async (e) => {
		e.preventDefault();

		if (loading) {
			return;
		}

		setLoading(true);

		const toggle = queueBuster.enabled
			? Api.QueueBusters.Disable
			: Api.QueueBusters.Enable;
		await toggle(department, fetchQueueBuster, setError);

		setLoading(false);
	};

	useEffect(() => {
		fetchQueueBuster();
	}, [department]);

	useEffect(() => {
		if (error !== null) {
			setTimeout(() => setError(null), 4000);
		}
	}, [error]);

	if (queueBuster === null) {
		return null;
	}

	return (
		<Details
			error={error}
			queueBuster={queueBuster}
			handleToggle={handleToggle}
		/>
	);
};

DetailsContainer.propTypes = propTypes;
DetailsContainer.defaultProps = defaultProps;

export default DetailsContainer;
