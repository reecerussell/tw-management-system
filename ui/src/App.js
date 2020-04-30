import React, { Component } from "react";
import { HashRouter, Route, Switch } from "react-router-dom";
import "./App.scss";

const loading = () => (
	<div className="animated fadeIn pt-3 text-center">Loading...</div>
);

// Containers
const Authorize = React.lazy(() => import("./containers/auth/authorize"));
const Layout = React.lazy(() => import("./containers/layout"));

// Pages
const Login = React.lazy(() => import("./views/pages/login"));

class App extends Component {
	render() {
		return (
			<HashRouter>
				<React.Suspense fallback={loading()}>
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
				</React.Suspense>
			</HashRouter>
		);
	}
}

export default App;
