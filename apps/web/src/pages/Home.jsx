import { useState, useEffect } from 'react'
import {
  AppBar,
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  Container,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  FormControl,
  Grid,
  InputAdornment,
  InputLabel,
  LinearProgress,
  MenuItem,
  Select,
  Stack,
  TextField,
  Toolbar,
  Typography,
} from '@mui/material'
import SearchIcon from '@mui/icons-material/Search'
import SchoolIcon from '@mui/icons-material/School'
import PlaceIcon from '@mui/icons-material/Place'
import AttachMoneyIcon from '@mui/icons-material/AttachMoney'
import VerifiedIcon from '@mui/icons-material/Verified'
import { healthCheck } from '@services/api'

const listingTemplates = [
  {
    id: 1,
    title: 'Campus Walk Studios',
    distance: '6 min walk',
    price: 980,
    beds: '1 bed',
    bath: '1 bath',
    furnished: true,
    neighborhood: 'Downtown',
  },
  {
    id: 2,
    title: 'Maple Student Lofts',
    distance: '12 min bike',
    price: 760,
    beds: 'Shared 2 bed',
    bath: '1 bath',
    furnished: false,
    neighborhood: 'North Campus',
  },
  {
    id: 3,
    title: 'Library View Residences',
    distance: '9 min transit',
    price: 1120,
    beds: '1 bed',
    bath: '1 bath',
    furnished: true,
    neighborhood: 'Midtown',
  },
]

function Home() {
  const [apiStatus, setApiStatus] = useState('Checking...')
  const [search, setSearch] = useState('')
  const [budget, setBudget] = useState('1200')
  const [openApplyDialog, setOpenApplyDialog] = useState(false)

  useEffect(() => {
    healthCheck()
      .then(data => setApiStatus(data.message))
      .catch(() => setApiStatus('API not available'))
  }, [])

  const filteredListings = listingTemplates.filter(listing => {
    const withinBudget = listing.price <= Number(budget)
    const matchText =
      listing.title.toLowerCase().includes(search.toLowerCase()) ||
      listing.neighborhood.toLowerCase().includes(search.toLowerCase())

    return withinBudget && matchText
  })

  return (
    <Box sx={{ minHeight: '100vh', background: 'linear-gradient(180deg, #eaf2ff 0%, #f7f9fd 45%, #f4f7fb 100%)' }}>
      <AppBar position="static" elevation={0} sx={{ background: 'transparent', color: 'text.primary' }}>
        <Toolbar>
          <SchoolIcon sx={{ mr: 1 }} />
          <Typography variant="h6" sx={{ flexGrow: 1 }}>
            Sanctor Housing
          </Typography>
          <Chip size="small" color={apiStatus.includes('not') ? 'error' : 'success'} label={`API: ${apiStatus}`} />
        </Toolbar>
      </AppBar>

      <Container maxWidth="lg" sx={{ py: 4 }}>
        <Stack spacing={3}>
          <Box>
            <Typography variant="h4" gutterBottom>
              Find Your Student Home
            </Typography>
            <Typography color="text.secondary">
              Template #1: Listing Explorer with quick filters and one-click apply action.
            </Typography>
          </Box>

          <Card elevation={0} sx={{ border: '1px solid #d6e2f1' }}>
            <CardContent>
              <Grid container spacing={2}>
                <Grid item xs={12} md={8}>
                  <TextField
                    fullWidth
                    label="Search by neighborhood or property"
                    value={search}
                    onChange={e => setSearch(e.target.value)}
                    InputProps={{
                      startAdornment: (
                        <InputAdornment position="start">
                          <SearchIcon />
                        </InputAdornment>
                      ),
                    }}
                  />
                </Grid>
                <Grid item xs={12} md={4}>
                  <FormControl fullWidth>
                    <InputLabel id="budget-label">Max Monthly Budget</InputLabel>
                    <Select
                      labelId="budget-label"
                      value={budget}
                      label="Max Monthly Budget"
                      onChange={e => setBudget(e.target.value)}
                    >
                      <MenuItem value="800">$800</MenuItem>
                      <MenuItem value="1000">$1,000</MenuItem>
                      <MenuItem value="1200">$1,200</MenuItem>
                      <MenuItem value="1500">$1,500</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
              </Grid>
            </CardContent>
          </Card>

          <Grid container spacing={2}>
            {filteredListings.map(listing => (
              <Grid item xs={12} md={6} lg={4} key={listing.id}>
                <Card sx={{ height: '100%' }}>
                  <CardContent>
                    <Stack spacing={1.25}>
                      <Typography variant="h6">{listing.title}</Typography>
                      <Stack direction="row" spacing={1} alignItems="center">
                        <PlaceIcon fontSize="small" color="action" />
                        <Typography variant="body2" color="text.secondary">
                          {listing.neighborhood} • {listing.distance}
                        </Typography>
                      </Stack>
                      <Stack direction="row" spacing={1} alignItems="center">
                        <AttachMoneyIcon fontSize="small" color="action" />
                        <Typography variant="body2" color="text.secondary">
                          ${listing.price}/month • {listing.beds} • {listing.bath}
                        </Typography>
                      </Stack>
                      <Stack direction="row" spacing={1}>
                        <Chip size="small" label={listing.furnished ? 'Furnished' : 'Unfurnished'} />
                        <Chip size="small" color="success" icon={<VerifiedIcon />} label="Verified" />
                      </Stack>
                      <Button variant="contained" onClick={() => setOpenApplyDialog(true)}>
                        Start Application
                      </Button>
                    </Stack>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>

          {filteredListings.length === 0 && (
            <Card>
              <CardContent>
                <Typography>No listings match your filters yet.</Typography>
              </CardContent>
            </Card>
          )}

          <Card elevation={0} sx={{ border: '1px solid #d6e2f1' }}>
            <CardContent>
              <Typography variant="subtitle1" gutterBottom>
                Template #2: Application Progress Strip
              </Typography>
              <Typography variant="body2" color="text.secondary" sx={{ mb: 1 }}>
                Documents uploaded and profile completion can be tracked with a lightweight progress UI.
              </Typography>
              <LinearProgress variant="determinate" value={65} sx={{ height: 10, borderRadius: 99 }} />
            </CardContent>
          </Card>
        </Stack>
      </Container>

      <Dialog open={openApplyDialog} onClose={() => setOpenApplyDialog(false)} fullWidth maxWidth="sm">
        <DialogTitle>Quick Application</DialogTitle>
        <DialogContent>
          <Stack spacing={2} sx={{ mt: 1 }}>
            <TextField fullWidth label="Full name" />
            <TextField fullWidth label="School email" type="email" />
            <TextField fullWidth label="Preferred move-in date" type="date" InputLabelProps={{ shrink: true }} />
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenApplyDialog(false)}>Cancel</Button>
          <Button variant="contained" onClick={() => setOpenApplyDialog(false)}>
            Submit Template Flow
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  )
}

export default Home
