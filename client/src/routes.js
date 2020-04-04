import React from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'

import Dashboard from './components/Dashboard';
import NotFound from './components/NotFound';


export default function AppRouter() {
  return (
    <Router>
      <Switch>
        <Route exact path='/' component={Dashboard} />
        <Route exact component={NotFound} />
      </Switch>
    </Router>
  )
}