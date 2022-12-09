import * as React from 'react'
import { createStyles, Header, Container, Group, Text, Anchor } from '@mantine/core'
import { useHref } from 'react-router-dom'

const useStyles = createStyles((theme) => ({
  header: {
    backgroundColor: theme.fn.variant({ variant: 'filled', color: theme.primaryColor }).background,
    borderBottom: 0,
  },
  inner: {
    height: 56,
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  link: {
    display: 'block',
    lineHeight: 1,
    padding: '8px 12px',
    borderRadius: theme.radius.sm,
    textDecoration: 'none',
    color: theme.white,
    fontSize: theme.fontSizes.sm,
    fontWeight: 500,

    '&:hover': {
      backgroundColor: theme.fn.lighten(
        theme.fn.variant({ variant: 'filled', color: theme.primaryColor }).background!,
        0.1
      ),
    }
  }
}))

export const Navbar: React.FC = () => {
  const { classes } = useStyles()

  return (
    <Header height={56} className={classes.header}>
      <Container size="xl">
        <div className={classes.inner}>
          <Anchor href={useHref('/')}>
            <Text weight={500} size="lg" color="white">Aquareo</Text>
          </Anchor>
          <Group spacing={5}>
            <a href={useHref('/')} className={classes.link}>Dashboard</a>
            <a href={useHref('/metrics')} className={classes.link}>Metrics</a>
          </Group>
        </div>
      </Container>
    </Header>
  )
}