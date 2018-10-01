import React from 'react';
import { Route } from 'react-router-dom';
import { HomeView, GameView } from './views';

export const AppRoutes = () => (
  <main className="content">
    <Route exact path="/" component={ HomeView }/>
    <Route exact path="/game" component={ GameView } />
  </main>
);
