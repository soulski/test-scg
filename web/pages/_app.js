import React from 'react'
import App from 'next/app'
import {
  Container,
  Row,
  Col,
  Navbar,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink
} from 'reactstrap'

import 'bootstrap/dist/css/bootstrap.css';


class MyApp extends App {

  render() {
    const { Component, pageProps } = this.props
    return (
      <Container>
        <Row>
          <Col>
            <Navbar color="dark" dark expand="md">
              <NavbarBrand href="/" className="mr-auto">SCG</NavbarBrand>
              <Nav navbar>
                <NavItem>
                  <NavLink href="/scg">XYZ</NavLink>
                </NavItem>
                <NavItem>
                  <NavLink href="/restaurants">Restaurants</NavLink>
                </NavItem>
                <NavItem>
                  <NavLink href="/line">Line</NavLink>
                </NavItem>
              </Nav>
            </Navbar>
          </Col>
        </Row>
        <Row>
          <Col>
            <Component {...pageProps} />
          </Col>
        </Row>
      </Container>
    )
  }

}

export default MyApp
