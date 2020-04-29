import React from "react";
import {
	Card,
	CardHeader,
	CardBody,
	CardFooter,
	Row,
	Col,
	UncontrolledAlert,
} from "reactstrap";
import { Link } from "react-router-dom";

const Details = ({ error, queueBuster, handleEnable, handleDisable }) => (
	<Card>
		<CardHeader>Queue Buster</CardHeader>
		<CardBody>
			{error !== null ? (
				<UncontrolledAlert color="danger">{error}</UncontrolledAlert>
			) : null}
			<p>
				<b>Department</b>
				<br />
				{queueBuster.department}
			</p>
			<Row>
				<Col sm="6">
					<p>
						<b>Enabled</b>
						<br />
						{queueBuster.enabled ? "Yes" : "No"}
					</p>
				</Col>
				<Col sm="6">
					<p>
						<br />
						{queueBuster.enabled ? (
							<a href="#" onClick={handleDisable}>
								Click to disable
							</a>
						) : (
							<a href="#" onClick={handleEnable}>
								Click to enable
							</a>
						)}
					</p>
				</Col>
			</Row>
		</CardBody>
		<CardFooter>
			<Link to="/queueBusters">Queue Busters</Link>
		</CardFooter>
	</Card>
);

export default Details;
