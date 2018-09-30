import React, { Component } from 'react';
import { GoogleLogin } from 'react-google-login';
import './app.component.css';

export class AppComponent extends Component {
  state = {
    isSignUpDialogOpen: false,
    isSignInDialogOpen: false,
  }

  _handleOnSignUp = (event) => {
    event.preventDefault();
    this.setState({ isSignUpDialogOpen: true });
  }

  _handleOnSignIn = (event) => {
    event.preventDefault();
    this.setState({ isSignInDialogOpen: true });
  }

  _handlOnSignUpDialogClosed = (event) => {
    event.preventDefault();
    this.setState({ isSignUpDialogOpen: false });
  }

  _handleOnSignInDialogCloed = (event) => {
    event.preventDefault();
    this.setState({ isSignInDialogOpen: false });
  }

  _handleGoogleSignUpResponse = (response) => {
    console.log(response);
  }

  render() {
    return (
      <div className="app">
        <header className="menu">
          <nav>
            <h1>Movies</h1>
            <div className="sign">
              <button className="secondary" onClick={ this._handleOnSignUp }>Inscription</button>
              <button className="primary" onClick={ this._handleOnSignIn }>Connexion</button>
            </div>
          </nav>
        </header>
        <main className="content">
          {
            this.state.isSignUpDialogOpen ?
            <div className="dialog">
              <div className="content">
                <header>
                  <h2>Inscription</h2>
                </header>
                <main>
                  <div className="sign-social">
                    <GoogleLogin
                      clientId="857123691814-2gmks9c09okm1mk86nachs1vbpbk33nr.apps.googleusercontent.com"
                      onSuccess={ this._handleGoogleSignUpResponse }
                      onFailure={ this._handleGoogleSignUpResponse }
                      buttonText="Se connecter avec google"
                    >

                    </GoogleLogin>
                  </div>
                  <form onSubmit={ this._handlOnSignUpDialogClosed }>
                    <p>ou utiliser une inscription classique</p>
                    <div className="row">
                      <label>Pseudo :</label>
                      <input type="text"/>
                    </div>
                    <div className="row">
                      <label>Email :</label>
                      <input type="text" />
                    </div>
                    <div className="row">
                      <label>Mot de passe :</label>
                      <input type="password" />
                    </div>
                    <div className="row">
                      <label>Confirmation mot de pass :</label>
                      <input type="password" />
                    </div>
                    <div className="actions">
                      <button type="submit">Valider</button>
                    </div>
                  </form>
                </main>
              </div>
            </div>
            :
            null
          }
          {
            this.state.isSignInDialogOpen ?
            <div className="dialog">
              <div className="content">
                <header>
                  <h2>Connexion</h2>
                </header>
                <main>
                  <div>

                  </div>
                  <form onSubmit={ this._handleOnSignInDialogCloed }>
                    <p>ou utiliser une identification classique</p>
                    <div className="row">
                      <label>Email :</label>
                      <input type="text" />
                    </div>
                    <div className="row">
                      <label>Mot de passe :</label>
                      <input type="password" />
                    </div>
                    <div className="actions">
                      <button type="submit">Valider</button>
                    </div>
                  </form>
                </main>
              </div>
            </div>
            :
            null
          }
        </main>
        <footer>
          
        </footer>
      </div>
    );
  }
}
