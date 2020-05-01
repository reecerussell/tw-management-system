import React from "react";
import {
	Modal,
	ModalHeader,
	ModalBody,
	ModalFooter,
	ButtonGroup,
	Button,
	UncontrolledAlert,
	Spinner,
} from "reactstrap";

const Delete = ({ loading, error, handleSubmit, toggle }) => (
	<Modal isOpen={true} toggle={toggle}>
		<ModalHeader toggle={toggle}>Delete Queue Buster</ModalHeader>
		<ModalBody>
			{error !== null ? (
				<UncontrolledAlert color="danger">{error}</UncontrolledAlert>
			) : null}
			<h4>Are you sure?</h4>
			<p>
				This action is permanent and cannot be reversed. Are you sure
				you want to delete this queue buster?
			</p>
		</ModalBody>
		<ModalFooter>
			<ButtonGroup>
				<Button type="button" onClick={toggle} color="secondary">
					Cancel
				</Button>
				<Button type="button" onClick={handleSubmit} color="danger">
					{loading ? <Spinner size="sm" color="danger" /> : "Delete"}
				</Button>
			</ButtonGroup>
		</ModalFooter>
	</Modal>
);

export default Delete;
