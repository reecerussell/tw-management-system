import React from "react";
import { EditContainer } from "../../containers/users";
import { Col, Row } from "reactstrap";
import { useParams } from "react-router-dom";

const Edit = () => {
	const { id } = useParams();

	return (
		<Row>
			<Col lg="4">
				<EditContainer id={id} />
			</Col>
		</Row>
	);
};

export default Edit;
