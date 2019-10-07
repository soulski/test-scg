import React, { useState, useCallback } from 'react'
import Head from 'next/head'
import {
  Row,
  Col,
  Form,
  FormGroup,
  Input,
  Label,
  Button,
} from 'reactstrap'
import axios from 'axios'
import qs from 'qs'


const SCG = () => {
  let [xyz, setXYZ] = useState(['', '', '', '', '', '', '']);
  let computeXYZ = useCallback(() => {
    axios.get('http://localhost:8080/api/xyz', {
      params: {
        numbers: [5, 9, 15, 23]
      },
      paramsSerializer: params => {
        return qs.stringify(params, { arrayFormat: 'repeat' })
      }
    }).then((output) => {
      setXYZ(output.data)
    })
  }, [xyz])

  return (
    <Form style={{ 'min-height': '500px' }}>
      <Head>
        <title>SCG</title>
        <link rel='icon' href='/static/favicon.ico' importance='low' />
      </Head>


      <FormGroup className="justify-content-md-center" row>
        <Col xs={2}>
          <Label>Calculate X Y Z</Label>
        </Col>
      </FormGroup>

      <FormGroup className="justify-content-md-center" row>
        <Col xs={1}><Input type="text" placeholder="X" readOnly value={xyz[0]}/></Col>
        <Col xs={1}><Label>5</Label></Col>
        <Col xs={1}><Label>9</Label></Col>
        <Col xs={1}><Label>15</Label></Col>
        <Col xs={1}><Label>23</Label></Col>
        <Col xs={1}><Input type="text" placeholder="Y" readOnly value={xyz[5]}/></Col>
        <Col xs={1}><Input type="text" placeholder="Z" readOnly value={xyz[6]}/></Col>
      </FormGroup>
      <FormGroup className="justify-content-md-center" row>
        <Col xs={2}>
          <Button onClick={computeXYZ}>Calculate</Button>
        </Col>
      </FormGroup>
    </Form>
  )
}

export default SCG
