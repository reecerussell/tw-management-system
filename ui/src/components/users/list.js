import React from "react";
import { Link } from "react-router-dom";
import { Table, UncontrolledAlert } from "reactstrap";

const List = ({ error, users }) => (
	<>
		{error !== null ? (
			<UncontrolledAlert color="danger">{error}</UncontrolledAlert>
		) : null}
		<Table striped responsive>
			<thead>
				<tr>
					<th>Username</th>
					<th>Email</th>
					<th />
				</tr>
			</thead>
			<tbody>
				{users.map((user, idx) => (
					<tr key={idx}>
						<td>{user.username}</td>
						<td>{user.email}</td>
						<td>
							<Link to={`/users/${user.id}/details`}>View</Link>
						</td>
					</tr>
				))}
			</tbody>
		</Table>
	</>
);

export default List;
