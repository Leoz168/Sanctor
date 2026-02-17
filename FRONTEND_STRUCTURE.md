# Frontend Structure - Complete

Your frontend has been successfully restructured with Vite and modern best practices!

## ✅ Directory Structure

```
apps/web/
├── src/
│   ├── pages/                    # Page components (routes)
│   │   └── Home.jsx              # ✓ Main landing page
│   │
│   ├── components/               # Reusable UI components
│   │   ├── Button.jsx            # ✓ Styled button component
│   │   ├── Card.jsx              # ✓ Card container
│   │   ├── Loading.jsx           # ✓ Loading spinner
│   │   └── index.js              # ✓ Barrel exports
│   │
│   ├── hooks/                    # Custom React hooks
│   │   ├── useUsers.js           # ✓ User data fetching hook
│   │   ├── useApi.js             # ✓ Generic API hook
│   │   └── index.js              # ✓ Barrel exports
│   │
│   ├── services/                 # API integration layer
│   │   ├── api.js                # ✓ Axios instance + interceptors
│   │   ├── authService.js        # ✓ Auth API calls
│   │   ├── userService.js        # ✓ User CRUD operations
│   │   └── index.js              # ✓ Barrel exports
│   │
│   ├── store/                    # State management (Zustand)
│   │   ├── authStore.js          # ✓ Authentication state
│   │   └── userStore.js          # ✓ User state
│   │
│   ├── styles/                   # CSS files
│   │   ├── global.css            # ✓ Global styles
│   │   ├── Home.css              # ✓ Page styles
│   │   ├── Button.css            # ✓ Component styles
│   │   ├── Card.css
│   │   └── Loading.css
│   │
│   ├── App.jsx                   # ✓ Router setup
│   └── main.jsx                  # ✓ Entry point
│
├── public/                       # Static assets
├── index.html                    # ✓ HTML template (Vite entry)
├── vite.config.js                # ✓ Vite configuration
├── package.json                  # ✓ Updated for Vite
├── Dockerfile                    # ✓ Updated for Vite build
└── .env.example                  # ✓ Environment variables
```

## 🚀 Tech Stack

### Build Tool
- **Vite** - Lightning-fast dev server & optimized builds
- **Hot Module Replacement (HMR)** - Instant updates

### Frontend Framework
- **React 18** - Latest React with concurrent features
- **React Router v6** - Client-side routing

### State Management
- **Zustand** - Lightweight, hook-based state
- Persistent storage with localStorage

### HTTP Client
- **Axios** - Promise-based HTTP client
- Request/response interceptors
- Automatic token injection

### Styling
- **CSS Modules** - Component-scoped styles
- Modern CSS with variables

## 🎯 Key Features

### Path Aliases
Use clean imports instead of relative paths:
```javascript
// ❌ Before
import Button from '../../../components/Button'

// ✅ After
import Button from '@components/Button'
```

### API Services Layer
Centralized API calls with error handling:
```javascript
import { getUsers, createUser } from '@services/userService'

const users = await getUsers()
```

### Custom Hooks
Reusable logic with hooks:
```javascript
import { useUsers } from '@hooks'

const { users, loading, error, refetch } = useUsers()
```

### Global State with Zustand
Simple, performant state management:
```javascript
import useAuthStore from '@store/authStore'

const { user, isAuthenticated, logout } = useAuthStore()
```

### Axios Interceptors
- **Request**: Auto-inject JWT tokens
- **Response**: Handle 401 errors globally

## 📦 Available Components

### Button
```javascript
import { Button } from '@components'

<Button variant="primary" onClick={handleClick}>
  Click Me
</Button>
```

### Card
```javascript
import { Card } from '@components'

<Card title="User Info">
  <p>Content here</p>
</Card>
```

### Loading
```javascript
import { Loading } from '@components'

{loading && <Loading message="Fetching data..." />}
```

## 🔧 Development

### Start Development Server
```bash
cd apps/web
npm install
npm run dev
```

### Build for Production
```bash
npm run build      # Output to dist/
npm run preview    # Preview production build
```

### Environment Variables
Create `.env.local`:
```bash
VITE_API_URL=http://localhost:8080
```

Access in code:
```javascript
const apiUrl = import.meta.env.VITE_API_URL
```

## 🐳 Docker

### Production
```bash
docker compose up --build
```

### Development (Hot Reload)
```bash
docker compose -f docker-compose.dev.yml up
```

## 📝 Adding New Features

### New Page
1. Create `src/pages/NewPage.jsx`
2. Add route in `src/App.jsx`:
```javascript
<Route path="/new-page" element={<NewPage />} />
```

### New Component
1. Create `src/components/MyComponent.jsx`
2. Create `src/styles/MyComponent.css`
3. Export from `src/components/index.js`:
```javascript
export { default as MyComponent } from './MyComponent'
```

### New API Service
1. Add to existing service or create new file in `src/services/`
2. Export from `src/services/index.js`

### New Hook
1. Create `src/hooks/useMyHook.js`
2. Export from `src/hooks/index.js`

### New Store
1. Create `src/store/myStore.js` using Zustand
2. Use in components:
```javascript
import useMyStore from '@store/myStore'
```

## 🌐 API Integration

The frontend is configured to work with your Go backend:

### User Operations
```javascript
import { getUsers, createUser, updateUser, deleteUser } from '@services'

// List all users
const users = await getUsers()

// Create user
const user = await createUser({
  email: 'user@example.com',
  username: 'username',
  firstName: 'First',
  lastName: 'Last'
})

// Update user
await updateUser(userId, { firstName: 'Updated' })

// Delete user
await deleteUser(userId)
```

### Authentication (Scaffolded)
```javascript
import { login, register, logout, isAuthenticated } from '@services'

// Login
const { token, user } = await login('email@example.com', 'password')

// Register
await register(userData)

// Check auth status
if (isAuthenticated()) {
  // User is logged in
}
```

## 🎨 Styling Guide

### Global Styles
Edit `src/styles/global.css` for app-wide styles

### Component Styles
Each component has its own CSS file:
- Scoped to component
- No naming conflicts
- Easy to maintain

### CSS Variables
Define in `global.css`:
```css
:root {
  --primary-color: #61dafb;
  --background: #282c34;
}
```

## 🔥 Benefits of New Structure

✅ **Vite** - 10-100x faster than CRA  
✅ **Clean Architecture** - Organized by feature  
✅ **Path Aliases** - No more `../../../`  
✅ **Zustand** - Simpler than Redux  
✅ **Axios Interceptors** - Global auth handling  
✅ **Custom Hooks** - Reusable logic  
✅ **Service Layer** - Separated concerns  

## 📚 Resources

- [Vite Documentation](https://vitejs.dev/)
- [React Router](https://reactrouter.com/)
- [Zustand](https://github.com/pmndrs/zustand)
- [Axios](https://axios-http.com/)

## 🎉 You're All Set!

Your frontend is now:
- ⚡ Lightning fast with Vite
- 🎯 Well-organized and scalable
- 🔌 Ready to integrate with backend
- 🎨 Modern and maintainable

Access your app at: http://localhost:3000
