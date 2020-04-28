import React from "react";
import { DetailsContainer } from "../../containers/users";
import { Card, Col, Row, CardHeader, CardBody, CardFooter } from "reactstrap";
import { useParams, Link } from "react-router-dom";

const Details = () => {
	const { id } = useParams();

	return (
		<Row>
			<Col lg="4">
				<Card>
					<CardHeader>User</CardHeader>
					<CardBody>
						<DetailsContainer id={id} />
					</CardBody>
					<CardFooter>
						<span style={{ float: "right" }}>
							<Link to={`/users/${id}/edit`}>Edit</Link>
						</span>
					</CardFooter>
				</Card>
			</Col>
		</Row>
	);
};

export default Details;
