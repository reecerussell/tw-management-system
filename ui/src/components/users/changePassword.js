import React from "react";
import {
	Row,
	Col,
	Form,
	FormGroup,
	Label,
	Input,
	Button,
	Spinner,
	UncontrolledAlert,
	Card,
	CardHeader,
	CardBody,
	CardFooter,
} from "reactstrap";

const ChangePassword = ({
	error,
	success,
	loading,
	current,
	password,
	confirm,
	handleUpdateText,
	handleSubmit,
}) => (
	<Form onSubmit={handleSubmit}>
		<Card>
			<CardHeader>Change Password</CardHeader>
			<CardBody>
				{error !== null ? (
					<UncontrolledAlert color="danger">
						{error}
					</UncontrolledAlert>
				) : null}
				{success !== null ? (
					<UncontrolledAlert color="success">
						{success}
					</UncontrolledAlert>
				) : null}
				<FormGroup>
					<Row>
						<Col md="3">
							<Label>Current Password</Label>
						</Col>
						<Col md="9">
							<Input
								type="password"
								value={current}
								name="current"
								onChange={handleUpdateText}
								autoComplete="current-password"
							/>
						</Col>
					</Row>
				</FormGroup>
				<FormGroup>
					<Row>
						<Col md="3">
							<Label>New Password</Label>
						</Col>
						<Col md="9">
							<Input
								type="password"
								value={password}
								name="password"
								onChange={handleUpdateText}
								autoComplete="new-password"
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
								type="password"
								value={confirm}
								name="confirm"
								onChange={handleUpdateText}
								autoComplete="new-password confirm-password"
							/>
						</Col>
					</Row>
				</FormGroup>
			</CardBody>
			<CardFooter>
				<div style={{ float: "right" }}>
					<Button type="submit" color="success">
						{loading ? (
							<Spinner size="sm" color="default" />
						) : (
							"Save"
						)}
					</Button>
				</div>
			</CardFooter>
		</Card>
	</Form>
);

export default ChangePassword;
