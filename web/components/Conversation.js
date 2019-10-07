import React, { useState } from 'react'
import {
  Container,
  Row,
  Col,
  Toast,
  ToastBody,
  ToastHeader,
  Button,
  Form,
  FormGroup,
  Input
} from 'reactstrap';


const UserMessage = ({ message }) => (
  <Row>
    <Col>
      <Toast >
        <ToastHeader>User</ToastHeader>
        <ToastBody>{message.Text}</ToastBody>
      </Toast>
    </Col>
  </Row>
)

const SystemMessage = ({ message }) => (
  <Row>
    <Col>
      <div className="float-right">
        <Toast className="float-right">
          <ToastHeader>System</ToastHeader>
          <ToastBody>{message.Text}</ToastBody>
        </Toast>
      </div>
    </Col>
  </Row>
)

const Message = ({message}) => {
  if (message.Sender === "user") {
    return <UserMessage message={message}/>
  }
  else {
    return <SystemMessage message={message}/>
  }
}

const Conversation = ({
  conversation = { User: {}, Messages: [] },
  onSubmit = () => {}
}) => {
  const [message, setMessage] = useState("")

  const conversationItem = conversation.Messages.map(message => (
    <Message key={message.CreateDate} message={message}/>
  ))

  const handlerSubmit = (e) => {
    e.preventDefault();
    onSubmit(message)
  }

  return (
    <Container>
      <Row><Col>Conversation with User {conversation.User.Id}</Col></Row>
      {conversationItem}
      <Row>
        <Col>
          <Form inline style={{ 'width': '100%', 'minHeight': '100px' }}>
            <FormGroup row>
              <Input 
                type="textarea" 
                placeholder="Message send to user" 
                onChange={(e) => setMessage(e.target.value)}
                value={message} />
              <Button onClick={handlerSubmit}>Submit</Button>
            </FormGroup>
          </Form>
        </Col>
      </Row>
    </Container>
  )
}

export default Conversation

