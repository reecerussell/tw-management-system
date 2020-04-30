import React, { useState, useEffect } from "react";
import * as Api from "../../utils/api";
import { Create } from "../../components/queueBusters";

const CreateContainer = () => {
	const [department, setDepartment] = useState("");
	const [enabled, setEnabled] = useState(false);
	const [error, setError] = useState(null);
	const [loading, setLoading] = useState(false);
	const [redirect, setRedirect] = useState(null);

	const isValid = () => {
		if (department === "") {
			setError("Department is required.");
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

		await Api.QueueBusters.Create(
			{
				department,
				enabled,
			},
			async () => setRedirect(`/queueBusters/${department}/details`),
			setError
		);

		setLoading(false);
	};

	const handleUpdateDepartment = (e) => setDepartment(e.target.value);

	const handleUpdateEnabled = (e) => setEnabled(e.target.checked);

	useEffect(() => {
		if (error !== null) {
			setTimeout(() => setError(null), 4000);
		}
	}, [error]);

	return (
		<Create
			error={error}
			department={department}
			enabled={enabled}
			loading={loading}
			redirect={redirect}
			handleSubmit={handleSubmit}
			handleUpdateDepartment={handleUpdateDepartment}
			handleUpdateEnabled={handleUpdateEnabled}
		/>
	);
};

export default CreateContainer;
