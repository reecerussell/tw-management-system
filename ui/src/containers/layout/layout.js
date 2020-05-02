import React, { Suspense } from "react";
import { Redirect, Route, Switch } from "react-router-dom";
import * as router from "react-router-dom";
import { Container } from "reactstrap";
import { Logout } from "../../utils/user";

import {
	AppFooter,
	AppHeader,
	AppSidebar,
	AppSidebarFooter,
	AppSidebarForm,
	AppSidebarHeader,
	AppSidebarMinimizer,
	AppBreadcrumb2 as AppBreadcrumb,
	AppSidebarNav2 as AppSidebarNav,
} from "@coreui/react";
// sidebar nav config
import navigation from "../../_nav";
// routes config
import routes from "../../routes";

const Footer = React.lazy(() => import("./footer"));
const Header = React.lazy(() => import("./header"));

const Layout = (props) => {
	const loading = () => (
		<div className="animated fadeIn pt-1 text-center">Loading...</div>
	);

	const signOut = (e) => {
		e.preventDefault();
		Logout();
	};

	return (
		<div className="app">
			<AppHeader fixed>
				<Suspense fallback={loading()}>
					<Header onLogout={(e) => signOut(e)} />
				</Suspense>
			</AppHeader>
			<div className="app-body">
				<AppSidebar fixed display="lg">
					<AppSidebarHeader />
					<AppSidebarForm />
					<Suspense>
						<AppSidebarNav
							navConfig={navigation}
							router={router}
							{...props}
						/>
					</Suspense>
					<AppSidebarFooter />
					<AppSidebarMinimizer />
				</AppSidebar>
				<main className="main">
					<AppBreadcrumb appRoutes={routes} router={router} />
					<Container fluid>
						<Suspense fallback={loading()}>
							<Switch>
								{routes.map((route, idx) => {
									return route.component ? (
										<Route
											key={idx}
											path={route.path}
											exact={route.exact}
											name={route.name}
											render={(props) => (
												<route.component {...props} />
											)}
										/>
									) : null;
								})}
								<Redirect from="/" to="/dashboard" />
							</Switch>
						</Suspense>
					</Container>
				</main>
			</div>
			<AppFooter>
				<Suspense fallback={loading()}>
					<Footer />
				</Suspense>
			</AppFooter>
		</div>
	);
};

export default Layout;
