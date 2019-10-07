import React, { useState, useEffect, useCallback } from 'react'
import { useRouter } from 'next/router'
import axios from 'axios'
import Conversation from '../../components/Conversation'

const ConversationPage = () => {
  const router = useRouter()
  const { cid } = router.query

  let [loaded, setLoaded] = useState(false);
  useEffect(() => {
    if (!loaded && cid) {
      axios.get('http://localhost:8080/api/line/conversations/' + cid)
        .then((output) => {
          setLoaded(true)
          setConversation(output.data)
        })
    }
  }, [loaded, cid])

  let [conversation, setConversation] = useState({ User: {}, Messages: []});
  let sendMessage = useCallback((message) => {
    console.log(message)
    axios.post('http://localhost:8080/api/line/conversations/' + cid, {
      Message: message
    }).then((output) => {
      setConversation(output.data)
    })
  }, [conversation])

  const handleSubmit = (message) => {
    sendMessage(message)
  }

  return (
    <Conversation conversation={conversation} onSubmit={handleSubmit}/>
  )
}

export default ConversationPage 
