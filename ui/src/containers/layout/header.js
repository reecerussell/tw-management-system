import React from "react";
import { Link, NavLink } from "react-router-dom";
import {
	Badge,
	UncontrolledDropdown,
	DropdownItem,
	DropdownMenu,
	DropdownToggle,
	Nav,
	NavItem,
} from "reactstrap";
import {
	AppAsideToggler,
	AppNavbarBrand,
	AppSidebarToggler,
} from "@coreui/react";
import logo from "../../assets/img/brand/logo.svg";
import sygnet from "../../assets/img/brand/sygnet.svg";

const propTypes = {};
const defaultProps = {};

const Header = ({ onLogout }) => (
	<>
		<AppSidebarToggler className="d-lg-none" display="md" mobile />
		<AppNavbarBrand
			full={{ src: logo, width: 89, height: 25, alt: "CoreUI Logo" }}
			minimized={{
				src: sygnet,
				width: 30,
				height: 30,
				alt: "CoreUI Logo",
			}}
		/>
		<AppSidebarToggler className="d-md-down-none" display="lg" />

		<Nav className="d-md-down-none" navbar>
			<NavItem className="px-3">
				<NavLink to="/dashboard" className="nav-link">
					Dashboard
				</NavLink>
			</NavItem>
		</Nav>
		<Nav className="ml-auto" navbar>
			<UncontrolledDropdown nav direction="down">
				<DropdownToggle nav className="pr-3">
					Account
				</DropdownToggle>
				<DropdownMenu right>
					<DropdownItem header tag="div" className="text-center">
						<strong>Account</strong>
					</DropdownItem>
					<DropdownItem onClick={(e) => onLogout(e)}>
						<i className="fa fa-lock"></i> Logout
					</DropdownItem>
				</DropdownMenu>
			</UncontrolledDropdown>
		</Nav>
	</>
);

Header.propTypes = propTypes;
Header.defaultProps = defaultProps;

export default Header;
