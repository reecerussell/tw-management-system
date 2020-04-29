import React from "react";
import { Row, Col } from "reactstrap";
import { DetailsContainer } from "../../containers/queueBusters";
import { useParams } from "react-router-dom";

const Details = () => {
	const { department } = useParams();

	return (
		<Row>
			<Col lg="4">
				<DetailsContainer department={department} />
			</Col>
		</Row>
	);
};

export default Details;
