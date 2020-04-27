import React from "react";

const UserList = React.lazy(() => import("./views/users/list"));
const UserDetails = React.lazy(() => import("./views/users/details"));
const UserEdit = React.lazy(() => import("./views/users/edit"));

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
	{
		path: "/users/:id/edit",
		name: "Edit",
		component: UserEdit,
	},
];

export default routes;
