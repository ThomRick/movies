import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import './home.view.css';

export class HomeView extends Component {
  render() {
    return (
      <div className="home">
        <Link to="/game">Jouer</Link>
      </div>
    )
  }
}
