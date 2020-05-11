import React, { Suspense } from "react";
import { HashRouter, Route, Switch } from "react-router-dom";
import { IsAuthenticated, Logout } from "./utils/user";
import "./App.scss";

const Authorize = React.lazy(() => import("./containers/auth/authorize"));
const Layout = React.lazy(() => import("./containers/layout"));
const Login = React.lazy(() => import("./views/pages/login"));

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
			Logout();
		}
	}, 15000);

	return (
		<HashRouter>
			<Suspense fallback={loadingFallback}>
				<Switch>
					<Route
						exact
						path="/login"
						name="Login Page"
						render={(props) => <Login {...props} />}
					/>
					<Route
						path="/"
						name="Home"
						render={(props) => (
							<Authorize>
								<Layout {...props} />
							</Authorize>
						)}
					/>
				</Switch>
			</Suspense>
		</HashRouter>
	);
};

export default App;
