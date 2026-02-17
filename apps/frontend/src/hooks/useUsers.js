import { useState, useEffect } from 'react'
import { getUsers } from '@services/userService'

/**
 * Custom hook to fetch and manage users
 */
export const useUsers = () => {
  const [users, setUsers] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    fetchUsers()
  }, [])

  const fetchUsers = async () => {
    try {
      setLoading(true)
      const data = await getUsers()
      setUsers(data)
      setError(null)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const refetch = () => {
    fetchUsers()
  }

  return { users, loading, error, refetch }
}
