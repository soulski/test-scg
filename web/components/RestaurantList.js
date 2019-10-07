import React from 'react'
import { Table, Badge } from 'reactstrap';

const RestaurantList = ({ places }) => {
  const placeItems = places.map(place => (
    <tr key={place.Id}>
      <th>{place.Id}</th>
      <td>{place.Name}</td>
      <td>
        <Badge color="dark" pill>Lat : {place.Location.Lat}</Badge>
        <Badge color="dark" pill>Lng : {place.Location.Lng}</Badge>
      </td>
    </tr>
  ))

  return (
    <div>
      <div>Restaurants in Bangsue area</div>
      <Table>
        <thead>
          <tr>
            <th>Id</th>
            <th>Name</th>
            <th>Location</th>
          </tr>
        </thead>
        <tbody>
          {placeItems}
        </tbody>
      </Table>
    </div>
  )
}

export default RestaurantList
