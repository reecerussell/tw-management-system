import React from "react";
import {
	Card,
	CardHeader,
	CardBody,
	CardFooter,
	UncontrolledAlert,
	Button,
} from "reactstrap";
import { Link, Redirect } from "react-router-dom";
import { AppSwitch } from "@coreui/react";
import { DeleteContainer } from "../../containers/queueBusters";

const Details = ({
	error,
	queueBuster,
	handleToggle,
	toggleModal,
	deleteModal,
	deleteSuccess,
	redirect,
}) => {
	if (redirect !== null) {
		return <Redirect to={redirect} />;
	}

	return (
		<Card>
			<CardHeader>Queue Buster</CardHeader>
			<CardBody>
				{error !== null ? (
					<UncontrolledAlert color="danger">
						{error}
					</UncontrolledAlert>
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
				<Link to="/queueBusters" className="btn btn-link pl-0">
					Queue Busters
				</Link>

				<div className="float-right">
					<Button type="button" color="danger" onClick={toggleModal}>
						Delete
					</Button>
					{deleteModal ? (
						<DeleteContainer
							department={queueBuster.department}
							onSuccess={deleteSuccess}
							toggle={toggleModal}
						/>
					) : null}
				</div>
			</CardFooter>
		</Card>
	);
};

export default Details;
