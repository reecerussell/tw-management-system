import React from "react";
import { ChangePasswordContainer } from "../../containers/users";
import { Col, Row } from "reactstrap";
import { GetId } from "../../utils/user";

const ChangePassword = () => (
	<Row>
		<Col md="6">
			<ChangePasswordContainer id={GetId()} />
		</Col>
	</Row>
);

export default ChangePassword;
