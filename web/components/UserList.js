import React from 'react'
import { Table } from 'reactstrap';

const UserList = ({ users, onClick = ()=>{} }) => {
  const handleOnClick = (id) => (e) => {
    e.preventDefault();
    onClick(id)
  }

  const userItems = users.map(user => (
    <tr key={user.Id}>
      <th>
        <a href="#" onClick={handleOnClick(user.Id)}>{user.Id}</a>
      </th>
    </tr>
  ))

  return (
    <Table>
      <thead>
        <tr>
          <th>Id</th>
        </tr>
      </thead>
      <tbody>
        {userItems}
      </tbody>
    </Table>
  )
}

export default UserList
