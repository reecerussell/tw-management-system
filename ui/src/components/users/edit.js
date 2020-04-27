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
import { DeleteContainer } from "../../containers/users";

const Details = ({
	error,
	user,
	loading,
	handleSubmit,
	handleUpdateText,
	deleteModal,
	toggleModal,
	deleteSuccess,
	redirect,
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
									value={user.username}
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
									value={user.email}
									name="email"
									type="email"
									onChange={handleUpdateText}
								/>
							</Col>
						</Row>
					</FormGroup>
				</CardBody>
				<CardFooter>
					<Link
						className="btn btn-link pl-0"
						to={`/users/${user.id}/details`}
					>
						Back
					</Link>
					<ButtonGroup style={{ float: "right" }}>
						<Button
							type="button"
							color="danger"
							onClick={toggleModal}
						>
							Delete
						</Button>
						{deleteModal ? (
							<DeleteContainer
								id={user.id}
								onSuccess={deleteSuccess}
								toggle={toggleModal}
							/>
						) : null}
						<Button type="submit" color="success">
							{loading ? (
								<Spinner size="sm" color="default" />
							) : (
								"Save"
							)}
						</Button>
					</ButtonGroup>
				</CardFooter>
			</Card>
		</Form>
	);
};
export default Details;
