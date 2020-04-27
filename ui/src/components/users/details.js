import React from "react";
import { Row, Col, Input, FormGroup, Label } from "reactstrap";

const Details = ({ error, user }) => (
	<>
		<FormGroup>
			<Row>
				<Col md="3">
					<Label>Username</Label>
				</Col>
				<Col md="9">
					<Input value={user.username} disabled={true} />
				</Col>
			</Row>
		</FormGroup>
		<FormGroup>
			<Row>
				<Col md="3">
					<Label>Email</Label>
				</Col>
				<Col md="9">
					<Input value={user.email} disabled={true} />
				</Col>
			</Row>
		</FormGroup>
	</>
);

export default Details;
