import { useState, useEffect } from 'react'

/**
 * Custom hook for API calls with loading and error states
 */
export const useApi = (apiFunc, immediate = true) => {
  const [data, setData] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const execute = async (...params) => {
    try {
      setLoading(true)
      setError(null)
      const result = await apiFunc(...params)
      setData(result)
      return result
    } catch (err) {
      setError(err.message || 'An error occurred')
      throw err
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (immediate) {
      execute()
    }
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  return { data, loading, error, execute }
}
