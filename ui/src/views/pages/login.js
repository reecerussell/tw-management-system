import React from "react";
import { Card, CardBody, Col, Container, Row } from "reactstrap";
import { LoginContainer } from "../../containers/auth";

const Login = () => (
	<div className="app flex-row align-items-center">
		<Container>
			<Row className="justify-content-center">
				<Col md="6">
					<Card className="p-4">
						<CardBody>
							<LoginContainer />
						</CardBody>
					</Card>
				</Col>
			</Row>
		</Container>
	</div>
);

export default Login;
