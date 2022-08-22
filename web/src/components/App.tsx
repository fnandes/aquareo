import AppBar from '@mui/material/AppBar'
import Box from '@mui/material/Box'
import {deepPurple, indigo} from '@mui/material/colors'
import Container from '@mui/material/Container'
import CssBaseline from '@mui/material/CssBaseline'
import {createTheme, ThemeProvider} from '@mui/material/styles'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import * as React from 'react'
import {Dashboard} from './Dasboard'

const theme = createTheme({
  palette: {
    mode: 'dark',
    primary: deepPurple,
    secondary: indigo
  }
})

export const App: React.FC = () => (
  <ThemeProvider theme={theme}>
    <div>
      <CssBaseline/>
      <AppBar position="static" color="primary" enableColorOnDark>
        <Toolbar>
          <Typography variant="h6" component="div">
            Test
          </Typography>
        </Toolbar>
      </AppBar>
      <Box sx={{my: 4}}>
        <Container maxWidth="xl">
          <Dashboard/>
        </Container>
      </Box>
    </div>
  </ThemeProvider>
)