import React from "react";

const propTypes = {};
const defaultProps = {};

const Footer = () => (
	<>
		<span>
			<a href="https://coreui.io">CoreUI</a> &copy; 2019 creativeLabs.
		</span>
		<span className="ml-auto">
			Powered by <a href="https://coreui.io/react">CoreUI for React</a>
		</span>
	</>
);

Footer.propTypes = propTypes;
Footer.defaultProps = defaultProps;

export default Footer;
