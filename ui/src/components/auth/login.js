import React from "react";
import {
	Form,
	InputGroup,
	InputGroupAddon,
	InputGroupText,
	Spinner,
	Input,
	Button,
	Row,
	Col,
	UncontrolledAlert,
} from "reactstrap";

const Login = ({
	error,
	loading,
	handleSubmit,
	handleUpdateText,
	username,
	password,
}) => {
	return (
		<Form onSubmit={handleSubmit}>
			<h1>Login</h1>
			<p className="text-muted">Sign In to your account</p>
			{error !== null ? (
				<UncontrolledAlert color="danger">{error}</UncontrolledAlert>
			) : null}
			<InputGroup className="mb-3">
				<InputGroupAddon addonType="prepend">
					<InputGroupText>
						<i className="icon-user"></i>
					</InputGroupText>
				</InputGroupAddon>
				<Input
					type="text"
					placeholder="Username"
					autoComplete="username"
					name="username"
					value={username}
					onChange={handleUpdateText}
				/>
			</InputGroup>
			<InputGroup className="mb-4">
				<InputGroupAddon addonType="prepend">
					<InputGroupText>
						<i className="icon-lock"></i>
					</InputGroupText>
				</InputGroupAddon>
				<Input
					type="password"
					placeholder="Password"
					autoComplete="current-password"
					name="password"
					value={password}
					onChange={handleUpdateText}
				/>
			</InputGroup>
			<Row>
				<Col xs="6">
					<Button color="primary" className="px-4" type="submit">
						{loading ? (
							<Spinner size="sm" color="default" />
						) : (
							"Login"
						)}
					</Button>
				</Col>
			</Row>
		</Form>
	);
};

export default Login;
