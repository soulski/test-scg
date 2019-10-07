import React from 'react'
import Head from 'next/head'

const Home = () => (
  <span>Redirecting...</span>
)

Home.getInitialProps = async ({ res }) => {
  res.writeHead(302, { Location: '/scg' })
  res.end()
}

export default Home
