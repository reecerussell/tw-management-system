import React from "react";
import { Link, NavLink } from "react-router-dom";
import {
	UncontrolledDropdown,
	DropdownItem,
	DropdownMenu,
	DropdownToggle,
	Nav,
	NavItem,
} from "reactstrap";
import { AppNavbarBrand, AppSidebarToggler } from "@coreui/react";
import logo from "../../assets/img/brand/logo.svg";

const propTypes = {};
const defaultProps = {};

const Header = ({ onLogout }) => (
	<>
		<AppSidebarToggler className="d-lg-none" display="md" mobile />
		<AppNavbarBrand
			full={{ src: logo, width: 89, height: 25, alt: "Thames Water" }}
			minimized={{
				src: logo,
				width: 30,
				height: 30,
				alt: "Thames Water",
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
					<DropdownItem tag={Link} to="/changePassword">
						<i className="fa fa-key"></i> Change Password
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
