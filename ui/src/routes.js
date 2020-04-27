import React from "react";

const UserList = React.lazy(() => import("./views/users/list"));

const routes = [
	{
		path: "/users",
		exact: true,
		name: "Users",
		component: UserList,
	},
];

export default routes;
