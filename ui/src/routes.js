import React from "react";

const Dashboard = React.lazy(() => import("./views/pages/dashboard"));

const UserList = React.lazy(() => import("./views/users/list"));
const UserDetails = React.lazy(() => import("./views/users/details"));
const UserEdit = React.lazy(() => import("./views/users/edit"));
const UserCreate = React.lazy(() => import("./views/users/create"));
const UserChangePassword = React.lazy(() =>
	import("./views/users/changePassword")
);

const QueueBustersList = React.lazy(() => import("./views/queueBusters/list"));
const QueueBustersDetails = React.lazy(() =>
	import("./views/queueBusters/details")
);
const QueueBustersCreate = React.lazy(() =>
	import("./views/queueBusters/create")
);

const routes = [
	{
		path: "/",
		name: "Management System",
		exact: true,
	},
	{
		path: "/dashboard",
		name: "Dashboard",
		component: Dashboard,
	},
	{
		path: "/users",
		exact: true,
		name: "Users",
		component: UserList,
	},
	{
		path: "/users/create",
		name: "Create",
		component: UserCreate,
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
	{
		path: "/changePassword",
		name: "Change Password",
		exact: true,
		component: UserChangePassword,
	},
	{
		path: "/queueBusters",
		name: "Queue Busters",
		exact: true,
		component: QueueBustersList,
	},
	{
		path: "/queueBusters/create",
		name: "Create",
		component: QueueBustersCreate,
	},
	{
		path: "/queueBusters/:department/details",
		name: "Details",
		component: QueueBustersDetails,
	},
];

export default routes;
