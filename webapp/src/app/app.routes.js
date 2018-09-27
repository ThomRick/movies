import React from 'react';
import { Route } from 'react-router-dom';
import { HomeView } from './views';

export const AppRoutes = () => (
  <main>
    <Route exact path="/" component={ HomeView }/>
  </main>
);
