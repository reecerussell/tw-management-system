import React from "react";
import { Table, UncontrolledAlert } from "reactstrap";
import { Link } from "react-router-dom";

const List = ({ error, loading, items }) => (
	<>
		{error !== null ? (
			<UncontrolledAlert color="danger">{error}</UncontrolledAlert>
		) : null}
		<Table striped>
			<thead>
				<tr>
					<th>Department</th>
					<th>Enabled</th>
					<th></th>
				</tr>
			</thead>
			<tbody>
				{loading ? (
					<tr>
						<td colSpan="3">Loading...</td>
					</tr>
				) : (
					items.map((qb, idx) => (
						<tr key={idx}>
							<td>{qb.department}</td>
							<td>{qb.enabled ? "yes" : "no"}</td>
							<td>
								<Link
									to={`queueBuster/${qb.department}/details`}
								>
									View
								</Link>
							</td>
						</tr>
					))
				)}
			</tbody>
		</Table>
	</>
);

export default List;
