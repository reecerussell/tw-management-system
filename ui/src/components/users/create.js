import React from "react";
import {
	Row,
	Col,
	Input,
	Form,
	FormGroup,
	Label,
	UncontrolledAlert,
	ButtonGroup,
	Button,
	Spinner,
	Card,
	CardHeader,
	CardBody,
	CardFooter,
} from "reactstrap";
import { Link, Redirect } from "react-router-dom";

const Create = ({
	error,
	loading,
	handleSubmit,
	handleUpdateText,
	redirect,
	username,
	email,
	password,
	confirm,
}) => {
	if (redirect) {
		return <Redirect to={redirect} />;
	}

	return (
		<Form onSubmit={handleSubmit}>
			<Card>
				<CardHeader>User</CardHeader>
				<CardBody>
					{error !== null ? (
						<UncontrolledAlert color="danger">
							{error}
						</UncontrolledAlert>
					) : null}
					<FormGroup>
						<Row>
							<Col md="3">
								<Label>Username</Label>
							</Col>
							<Col md="9">
								<Input
									value={username}
									name="username"
									type="text"
									onChange={handleUpdateText}
								/>
							</Col>
						</Row>
					</FormGroup>
					<FormGroup>
						<Row>
							<Col md="3">
								<Label>Email</Label>
							</Col>
							<Col md="9">
								<Input
									value={email}
									name="email"
									type="email"
									onChange={handleUpdateText}
								/>
							</Col>
						</Row>
					</FormGroup>
					<FormGroup>
						<Row>
							<Col md="3">
								<Label>Password</Label>
							</Col>
							<Col md="9">
								<Input
									value={password}
									name="password"
									type="password"
									onChange={handleUpdateText}
								/>
							</Col>
						</Row>
					</FormGroup>
					<FormGroup>
						<Row>
							<Col md="3">
								<Label>Confirm Password</Label>
							</Col>
							<Col md="9">
								<Input
									value={confirm}
									name="confirm"
									type="password"
									onChange={handleUpdateText}
								/>
							</Col>
						</Row>
					</FormGroup>
				</CardBody>
				<CardFooter>
					<Link className="btn btn-link pl-0" to={`/users`}>
						Back to users
					</Link>
					<ButtonGroup style={{ float: "right" }}>
						<Button type="submit" color="success">
							{loading ? (
								<Spinner size="sm" color="default" />
							) : (
								"Create"
							)}
						</Button>
					</ButtonGroup>
				</CardFooter>
			</Card>
		</Form>
	);
};
export default Create;
