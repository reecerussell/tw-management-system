import React from "react";
import {
	Form,
	FormGroup,
	Label,
	Input,
	Button,
	Spinner,
	Row,
	Col,
	Card,
	CardHeader,
	CardBody,
	CardFooter,
	UncontrolledAlert,
} from "reactstrap";
import { Redirect } from "react-router-dom";
import { AppSwitch } from "@coreui/react";

const Create = ({
	error,
	loading,
	department,
	enabled,
	announcements,
	handleSubmit,
	handleUpdateDepartment,
	handleUpdateEnabled,
	handleUpdateAnnouncements,
	redirect,
}) => {
	if (redirect !== null) {
		return <Redirect to={redirect} />;
	}

	return (
		<Form onSubmit={handleSubmit}>
			<Card>
				<CardHeader>Queue Buster</CardHeader>
				<CardBody>
					{error !== null ? (
						<UncontrolledAlert color="danger">
							{error}
						</UncontrolledAlert>
					) : null}
					<FormGroup>
						<Row>
							<Col md="3">
								<Label>Department</Label>
							</Col>
							<Col md="9">
								<Input
									type="text"
									value={department}
									onChange={handleUpdateDepartment}
								/>
							</Col>
						</Row>
					</FormGroup>
					<FormGroup>
						<Row>
							<Col md="3">
								<Label>Enabled</Label>
							</Col>
							<Col md="9">
								<AppSwitch
									className={"float-lg-right"}
									variant={"pill"}
									label
									color={"info"}
									checked={enabled}
									onChange={handleUpdateEnabled}
									dataOn="Yes"
									dataOff="No"
								/>
							</Col>
						</Row>
					</FormGroup>
					<FormGroup>
						<Row>
							<Col md="3">
								<Label>Queue Announcements</Label>
							</Col>
							<Col md="9">
								<AppSwitch
									className={"float-lg-right"}
									variant={"pill"}
									label
									color={"info"}
									checked={announcements}
									onChange={handleUpdateAnnouncements}
									dataOn="Yes"
									dataOff="No"
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
};

export default Create;
