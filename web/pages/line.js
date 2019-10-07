import React, { useState, useEffect } from 'react'
import Router from 'next/router'
import axios from 'axios'
import UserList from '../components/UserList'

const Line = () => {
  let [loaded, setLoaded] = useState(false);
  let [users, setUsers] = useState([]);
  useEffect(() => {
    if (!loaded) {
      axios.get('http://localhost:8080/api/line/users')
        .then((output) => {
          setLoaded(true)
          setUsers(output.data)
        })
    }
  }, [loaded])

  const onUserClick = (userId) => {
    Router.push('/conversation/' + userId)
  }

  return (
    <UserList users={users} onClick={onUserClick}/>
  )
}

export default Line
