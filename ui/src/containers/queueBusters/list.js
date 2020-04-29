import React, { useState, useEffect } from "react";
import * as Api from "../../utils/api";
import { List } from "../../components/queueBusters";

const ListContainer = () => {
	const [items, setItems] = useState([]);
	const [error, setError] = useState(null);
	const [loading, setLoading] = useState(false);

	const fetchItems = async () => {
		if (loading) {
			return;
		}

		setLoading(true);

		await Api.QueueBusters.GetAll(
			async (res) => setItems(await res.json()),
			setError
		);

		setLoading(false);
	};

	useState(() => {
		fetchItems();
	}, []);

	useState(() => {
		if (error !== null) {
			setTimeout(() => setError(null), 4000);
		}
	}, [error]);

	return <List error={error} loading={loading} items={items} />;
};

export default ListContainer;
