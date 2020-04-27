import React, { Suspense, useState } from "react";
import { IsAuthenticated, Listen } from "../../utils/user";
import PropTypes from "prop-types";

const LoginView = React.lazy(() => import("../../views/pages/login"));

const propTypes = {
	children: PropTypes.node,
};
const defaultProps = {};

const AuthorizeContainer = ({ children }) => {
	const [state, setState] = useState(0);
	Listen(() => setState(state + 1));

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
