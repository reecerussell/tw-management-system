import React from "react";
import { CreateContainer } from "../../containers/users";
import { Col, Row } from "reactstrap";

const Create = () => {
	return (
		<Row>
			<Col lg="4">
				<CreateContainer />
			</Col>
		</Row>
	);
};

export default Create;
