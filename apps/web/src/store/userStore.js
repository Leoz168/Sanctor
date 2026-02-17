import { create } from 'zustand'

const useUserStore = create((set) => ({
  users: [],
  loading: false,
  error: null,

  setUsers: (users) => set({ users }),
  
  addUser: (user) =>
    set((state) => ({
      users: [...state.users, user],
    })),

  updateUser: (id, updatedUser) =>
    set((state) => ({
      users: state.users.map((user) =>
        user.id === id ? { ...user, ...updatedUser } : user
      ),
    })),

  deleteUser: (id) =>
    set((state) => ({
      users: state.users.filter((user) => user.id !== id),
    })),

  setLoading: (loading) => set({ loading }),
  
  setError: (error) => set({ error }),
}))

export default useUserStore
