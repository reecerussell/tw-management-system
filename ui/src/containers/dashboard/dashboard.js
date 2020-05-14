import React, { useState, useEffect } from "react";
import * as Api from "../../utils/api";
import { Dashboard } from "../../components/dashboard";

const DashboardContainer = () => {
	const [error, setError] = useState(null);
	const [items, setItems] = useState([]);

	const fetchItems = async () =>
		await Api.QueueBusters.GetAll(
			async (res) => setItems(await res.json()),
			setError
		);

	useEffect(() => {
		fetchItems();
	}, []);

	return <Dashboard error={error} items={items} />;
};

export default DashboardContainer;
