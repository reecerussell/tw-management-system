import React from "react";
import { ListContainer } from "../../containers/users";
import { Card, Col, Row, CardHeader, CardBody, CardFooter } from "reactstrap";
import { Link } from "react-router-dom";

const List = () => (
	<Row>
		<Col md="8" lg="6">
			<Card>
				<CardHeader>Users</CardHeader>
				<CardBody>
					<ListContainer />
				</CardBody>
				<CardFooter>
					<Link to="/users/create">Create new user</Link>
				</CardFooter>
			</Card>
		</Col>
	</Row>
);

export default List;
