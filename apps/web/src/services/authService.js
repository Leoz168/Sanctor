import api from './api'

// Login
export const login = async (email, password) => {
  const response = await api.post('/api/auth/login', { email, password })
  if (response.data.token) {
    localStorage.setItem('token', response.data.token)
  }
  return response.data
}

// Register
export const register = async (userData) => {
  const response = await api.post('/api/auth/register', userData)
  if (response.data.token) {
    localStorage.setItem('token', response.data.token)
  }
  return response.data
}

// Logout
export const logout = () => {
  localStorage.removeItem('token')
  window.location.href = '/login'
}

// Check if user is authenticated
export const isAuthenticated = () => {
  return !!localStorage.getItem('token')
}
