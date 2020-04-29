import React from "react";
import { Row, Col, Card, CardHeader, CardBody } from "reactstrap";
import { ListContainer } from "../../containers/queueBusters";

const List = () => (
	<Row>
		<Col lg="6">
			<Card>
				<CardHeader>Queue Busters</CardHeader>
				<CardBody>
					<ListContainer />
				</CardBody>
			</Card>
		</Col>
	</Row>
);

export default List;
