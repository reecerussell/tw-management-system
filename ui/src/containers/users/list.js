import React, { useState, useEffect } from "react";
import * as Api from "../../utils/api";
import { List } from "../../components/users";

const ListContainer = () => {
	const [users, setUsers] = useState([]);
	const [error, setError] = useState(null);
	const [loading, setLoading] = useState(false);

	const fetchUsers = async () => {
		if (loading) {
			return;
		}

		setLoading(true);

		Api.Users.All(async (res) => setUsers(await res.json()), setError);

		setLoading(false);
	};

	useEffect(() => {
		fetchUsers();
	}, []);

	return <List users={users} error={error} />;
};

export default ListContainer;
