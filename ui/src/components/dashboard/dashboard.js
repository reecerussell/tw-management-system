import React from "react";
import {
	Row,
	Col,
	Card,
	CardHeader,
	CardBody,
	Table,
	UncontrolledAlert,
} from "reactstrap";
import { Link } from "react-router-dom";

const sort = (a, b) => {
	if (a.enabled === b.enabled) {
		return 0;
	}

	if (a.enabled) {
		return -1;
	}

	if (b.enabled) {
		return 1;
	}
};

const Dashboard = ({ error, items }) => (
	<Row>
		<Col lg="6">
			<Card>
				<CardHeader>
					<i className="icon-list"></i> Queue Busters
				</CardHeader>
				<CardBody>
					{error !== null ? (
						<UncontrolledAlert color="danger">
							{error}
						</UncontrolledAlert>
					) : null}
					<Row>
						<Col sm="6">
							<div className="callout callout-info">
								<small className="text-muted">Active</small>
								<br />
								<strong className="h4">
									{items.filter((x) => x.enabled).length}
								</strong>
							</div>
						</Col>
						<Col sm="6">
							<div className="callout callout-danger">
								<small className="text-muted">Disabled</small>
								<br />
								<strong className="h4">
									{items.filter((x) => !x.enabled).length}
								</strong>
							</div>
						</Col>
					</Row>
					<hr className="mt-0" />
					<Table hover responsive className="table-outline mb-0">
						<thead className="thead-light">
							<tr>
								<th>Department</th>
								<th>Enabled</th>
								<th />
							</tr>
						</thead>
						<tbody>
							{items.sort(sort).map((item, idx) => (
								<tr key={idx}>
									<td>{item.department}</td>
									<td>{item.enabled ? "Yes" : "No"}</td>
									<td>
										<Link
											to={`/queueBusters/${item.department}/details`}
										>
											View
										</Link>
									</td>
								</tr>
							))}
						</tbody>
					</Table>
				</CardBody>
			</Card>
		</Col>
	</Row>
);

export default Dashboard;
