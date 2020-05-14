import React, { Suspense, useState, useEffect } from "react";
import { IsAuthenticated, Listen, Unlisten } from "../../utils/user";
import PropTypes from "prop-types";

const LoginView = React.lazy(() => import("../../views/pages/login"));

const propTypes = {
	children: PropTypes.node,
};
const defaultProps = {};

const AuthorizeContainer = ({ children }) => {
	const [state, update] = useState(0);
	const listenerCallback = () => update(state + 1);

	useEffect(() => {
		Listen("auth", listenerCallback);

		return () => Unlisten("auth");
	}, []);

	if (!IsAuthenticated()) {
		return (
			<Suspense fallback={<p>Loading...</p>}>
				<LoginView />
			</Suspense>
		);
	}

	return children;
};

AuthorizeContainer.propTypes = propTypes;
AuthorizeContainer.defaultProps = defaultProps;

export default AuthorizeContainer;
