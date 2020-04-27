import React from "react";
import { EditContainer } from "../../containers/users";
import { Col, Row } from "reactstrap";
import { useParams } from "react-router-dom";

const List = () => {
	const { id } = useParams();

	return (
		<Row>
			<Col md="4">
				<EditContainer id={id} />
			</Col>
		</Row>
	);
};

export default List;
