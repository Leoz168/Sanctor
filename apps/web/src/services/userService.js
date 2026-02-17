import api from './api'

// Get all users
export const getUsers = async () => {
  const response = await api.get('/api/users')
  return response.data
}

// Get user by ID
export const getUser = async (id) => {
  const response = await api.get(`/api/users/get?id=${id}`)
  return response.data
}

// Create new user
export const createUser = async (userData) => {
  const response = await api.post('/api/users/create', userData)
  return response.data
}

// Update user
export const updateUser = async (id, userData) => {
  const response = await api.put(`/api/users/update?id=${id}`, userData)
  return response.data
}

// Delete user
export const deleteUser = async (id) => {
  const response = await api.delete(`/api/users/delete?id=${id}`)
  return response.data
}
