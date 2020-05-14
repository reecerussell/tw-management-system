import React from "react";
import { Row, Col, Card, CardHeader, CardBody, CardFooter } from "reactstrap";
import { ListContainer } from "../../containers/queueBusters";
import { Link } from "react-router-dom";

const List = () => (
	<Row>
		<Col lg="6">
			<Card>
				<CardHeader>Queue Busters</CardHeader>
				<CardBody>
					<ListContainer />
				</CardBody>
				<CardFooter>
					<Link to="/queueBusters/create">
						Create new queue buster
					</Link>
				</CardFooter>
			</Card>
		</Col>
	</Row>
);

export default List;
