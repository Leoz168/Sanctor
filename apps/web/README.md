# Sanctor Frontend (React + Vite)

A modern React frontend built with Vite, featuring a clean architecture with Zustand for state management.

## Project Structure

```
frontend/
├── src/
│   ├── pages/              # Page components (routes)
│   │   └── Home.jsx
│   │
│   ├── components/         # Reusable UI components
│   │   ├── Button.jsx
│   │   ├── Card.jsx
│   │   ├── Loading.jsx
│   │   └── index.js
│   │
│   ├── hooks/              # Custom React hooks
│   │   ├── useUsers.js
│   │   ├── useApi.js
│   │   └── index.js
│   │
│   ├── services/           # API service layer
│   │   ├── api.js          # Axios instance & interceptors
│   │   ├── authService.js  # Auth API calls
│   │   ├── userService.js  # User API calls
│   │   └── index.js
│   │
│   ├── store/              # Zustand state stores
│   │   ├── authStore.js    # Authentication state
│   │   └── userStore.js    # User state
│   │
│   ├── styles/             # CSS modules & global styles
│   │   ├── global.css
│   │   ├── Home.css
│   │   ├── Button.css
│   │   ├── Card.css
│   │   └── Loading.css
│   │
│   ├── App.jsx             # Main App with Router
│   └── main.jsx            # Entry point
│
├── public/
│   └── index.html
│
├── vite.config.js          # Vite configuration
├── package.json
├── Dockerfile
└── .env.example
```

## Features

### ✅ Modern Stack
- **Vite** - Lightning fast dev server & build tool
- **React 18** - Latest React features
- **React Router** - Client-side routing
- **Zustand** - Lightweight state management
- **Axios** - HTTP client with interceptors

### ✅ Clean Architecture
- **Pages** - Route components
- **Components** - Reusable UI elements
- **Hooks** - Custom React hooks for logic reuse
- **Services** - API integration layer
- **Store** - Global state management

### ✅ Developer Experience
- Path aliases (`@components`, `@services`, etc.)
- Hot Module Replacement (HMR)
- Fast builds with Vite
- Environment variables

## Development

### Local Development

```bash
cd apps/web
npm install
npm run dev
```

The app will open at [http://localhost:3000](http://localhost:3000)

### Building

```bash
npm run build
npm run preview    # Preview production build
```

## Docker

### Production Build
```bash
docker build -t sanctor-frontend .
docker run -p 3000:80 sanctor-frontend
```

### Development with Docker Compose
```bash
docker-compose -f docker-compose.dev.yml up
```

## Path Aliases

Configure imports with `@` prefix:

```javascript
// Instead of:
import Button from '../../../components/Button'

// Use:
import Button from '@components/Button'
```

Available aliases:
- `@` → `src/`
- `@components` → `src/components/`
- `@pages` → `src/pages/`
- `@hooks` → `src/hooks/`
- `@services` → `src/services/`
- `@store` → `src/store/`
- `@styles` → `src/styles/`

## Environment Variables

Create a `.env.local` file (see `.env.example`):

```bash
VITE_API_URL=http://localhost:8080
```

Access in code:
```javascript
const apiUrl = import.meta.env.VITE_API_URL
```

## API Integration

### Using Services

```javascript
import { getUsers, createUser } from '@services/userService'

// Fetch users
const users = await getUsers()

// Create user
const newUser = await createUser({
  email: 'test@example.com',
  username: 'testuser',
  firstName: 'Test',
  lastName: 'User'
})
```

### Using Custom Hooks

```javascript
import { useUsers } from '@hooks'

function UserList() {
  const { users, loading, error, refetch } = useUsers()
  
  if (loading) return <Loading />
  if (error) return <div>Error: {error}</div>
  
  return (
    <div>
      {users.map(user => (
        <Card key={user.id} title={user.username}>
          {user.email}
        </Card>
      ))}
    </div>
  )
}
```

### Using Zustand Store

```javascript
import useAuthStore from '@store/authStore'

function Profile() {
  const { user, isAuthenticated, logout } = useAuthStore()
  
  return (
    <div>
      <p>Welcome, {user?.username}</p>
      <Button onClick={logout}>Logout</Button>
    </div>
  )
}
```

## Adding New Features

### New Page
1. Create `src/pages/MyPage.jsx`
2. Add route in `src/App.jsx`:
```javascript
<Route path="/my-page" element={<MyPage />} />
```

### New Component
1. Create `src/components/MyComponent.jsx`
2. Create `src/styles/MyComponent.css`
3. Export from `src/components/index.js`

### New API Service
1. Add functions to existing service or create new one
2. Export from `src/services/index.js`

### New Hook
1. Create `src/hooks/useMyHook.js`
2. Export from `src/hooks/index.js`

## Testing

```bash
npm test
```

## License

MIT
