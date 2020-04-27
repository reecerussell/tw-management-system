import React from "react";
import { ListContainer } from "../../containers/users";
import { Card, Col, Row, CardHeader, CardBody } from "reactstrap";

const List = () => (
	<Row>
		<Col md="8" lg="6">
			<Card>
				<CardHeader>Users</CardHeader>
				<CardBody>
					<ListContainer />
				</CardBody>
			</Card>
		</Col>
	</Row>
);

export default List;
