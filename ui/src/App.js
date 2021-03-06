import React, { Suspense } from "react";
import { HashRouter, Route, Switch } from "react-router-dom";
import { IsAuthenticated, Logout } from "./utils/user";
import "./App.scss";

const Authorize = React.lazy(() => import("./containers/auth/authorize"));
const Layout = React.lazy(() => import("./containers/layout"));

const App = () => {
	const loadingFallback = (
		<div className="animated fadeIn pt-3 text-center">Loading...</div>
	);

	setInterval(() => {
		/*
			Timeout handler. Every 15 seconds check if the user is still
			authenticated. If the user isn't call Logout, which triggers
			the Login page to show.
		*/

		if (!IsAuthenticated()) {
			console.log("Not Authenticated");
			Logout();
		} else {
			console.log("Authenticated");
		}
	}, 5000);

	return (
		<HashRouter>
			<Suspense fallback={loadingFallback}>
				<Authorize>
					<Switch>
						<Route
							path="/"
							name="Home"
							render={(props) => <Layout {...props} />}
						/>
					</Switch>
				</Authorize>
			</Suspense>
		</HashRouter>
	);
};

export default App;
