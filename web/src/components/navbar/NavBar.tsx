import { AppBar, Toolbar, Typography } from '@mui/material'
import * as React from 'react'

export const NavBar: React.FC = () => (
  <AppBar position="fixed">
    <Toolbar>
      <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
        Aquareo
      </Typography>
    </Toolbar>
  </AppBar>
)