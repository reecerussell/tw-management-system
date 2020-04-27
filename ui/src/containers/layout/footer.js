import React from "react";

const propTypes = {};
const defaultProps = {};

const Footer = () => (
	<>
		<span>
			<a href="#">Management System</a> &copy; 2020 Thames Water.
		</span>
		<span className="ml-auto">
			Developed by <a href="https://reece-russell.co.uk">Reece Russell</a>
		</span>
	</>
);

Footer.propTypes = propTypes;
Footer.defaultProps = defaultProps;

export default Footer;
