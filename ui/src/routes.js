import React from "react";

const UserList = React.lazy(() => import("./views/users/list"));
const UserDetails = React.lazy(() => import("./views/users/details"));

const routes = [
	{
		path: "/",
		name: "Management System",
		exact: true,
	},
	{
		path: "/dashboard",
		name: "Dashboard",
		component: UserList,
	},
	{
		path: "/users",
		exact: true,
		name: "Users",
		component: UserList,
	},
	{
		path: "/users/:id/details",
		name: "Details",
		component: UserDetails,
	},
];

export default routes;
