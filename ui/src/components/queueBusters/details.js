import React from "react";
import {
	Card,
	CardHeader,
	CardBody,
	CardFooter,
	UncontrolledAlert,
} from "reactstrap";
import { Link } from "react-router-dom";
import { AppSwitch } from "@coreui/react";

const Details = ({ error, queueBuster, handleToggle }) => (
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
			<p>
				<b>Enabled</b>
				<br />
				<AppSwitch
					className="pt-1"
					variant={"pill"}
					label
					color={"info"}
					checked={queueBuster.enabled}
					onChange={handleToggle}
					dataOn="Yes"
					dataOff="No"
				/>
			</p>
		</CardBody>
		<CardFooter>
			<Link to="/queueBusters">Queue Busters</Link>
		</CardFooter>
	</Card>
);

export default Details;
