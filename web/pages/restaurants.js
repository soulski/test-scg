import React, { useState, useEffect } from 'react'
import axios from 'axios'
import RestaurantList from '../components/RestaurantList'

const Restaurant = () => {
  let [loaded, setLoaded] = useState(false);
  let [places, setPlaces] = useState([]);
  useEffect(() => {
    if (!loaded) {
      axios.get('http://localhost:8080/api/restaurant')
        .then((output) => {
          setLoaded(true)
          setPlaces(output.data.Places)
        })
    }
  }, [loaded])

  return (
    <RestaurantList places={places} />
  )
}

export default Restaurant
