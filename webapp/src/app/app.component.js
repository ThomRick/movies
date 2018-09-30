import React, { Component } from 'react';
import { GoogleLogin } from 'react-google-login';
import './app.component.css';
import { BrowserRouter } from 'react-router-dom';
import { AppRoutes } from './app.routes';

export class AppComponent extends Component {
  state = {
    isSignUpDialogOpen: false,
    isSignInDialogOpen: false,
    signUp: {
      nickname: '',
      email: '',
      password: '',
      passwordCheck: '',
    },
    signIn: {
      email: '',
      password: '',
    },
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

  _handleOnGoogleSignUpResponse = (response) => {
    console.log(response);
    this.setState({
      isSignUpDialogOpen: false,
    });
  }

  _handleOnGoogleSignInResponse = (response) => {
    console.log(response);
    this.setState({
      isSignInDialogOpen: false,
    });
  }

  _handleOnSignUpNicknameChange = (event) => {
    const nickname = event.target.value;
    this.setState({
      signUp: {
        ...this.state.signUp,
        nickname,
      },
    });
  }

  _handleOnSignUpEmailChange = (event) => {
    const email = event.target.value;
    this.setState({
      signUp: {
        ...this.state.signUp,
        email,
      },
    });
  }

  _handleOnSignUpPasswordChange = (event) => {
    const password = event.target.value;
    this.setState({
      signUp: {
        ...this.state.signUp,
        password,
      },
    });
  }

  _handleOnSignUpPasswordCheckChange = (event) => {
    const passwordCheck = event.target.value;
    this.setState({
      signUp: {
        ...this.state.signUp,
        passwordCheck,
      },
    });
  }

  _handleOnSignInEmailChange = (event) => {
    const email = event.target.value;
    this.setState({
      signIn: {
        ...this.state.signIn,
        email,
      },
    });
  }

  _handleOnSignInPasswordChange = (event) => {
    const password = event.target.value;
    this.setState({
      signIn: {
        ...this.state.signIn,
        password,
      },
    });
  }

  render() {
    return (
      <BrowserRouter>
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
          <AppRoutes />
          <footer></footer>
          {
            this.state.isSignUpDialogOpen ?
            <div className="dialog">
              <div className="background" onClick={ () => this.setState({ isSignUpDialogOpen: false }) }></div>
              <div className="content">
                <header>
                  <h2>Inscription</h2>
                </header>
                <main>
                  <div className="sign-social">
                    <GoogleLogin
                      clientId="857123691814-2gmks9c09okm1mk86nachs1vbpbk33nr.apps.googleusercontent.com"
                      onSuccess={ this._handleOnGoogleSignUpResponse }
                      onFailure={ this._handleOnGoogleSignUpResponse }
                      buttonText="Se connecter avec google"
                    />
                  </div>
                  <form onSubmit={ this._handlOnSignUpDialogClosed }>
                    <p>ou utiliser une inscription classique</p>
                    <div className="row">
                      <label>Pseudo :</label>
                      <input
                        type="text"
                        value={ this.state.signUp.nickname }
                        onChange={ this._handleOnSignUpNicknameChange }
                      />
                    </div>
                    <div className="row">
                      <label>Email :</label>
                      <input
                        type="email"
                        value={ this.state.signUp.email }
                        onChange={ this._handleOnSignUpEmailChange }
                      />
                    </div>
                    <div className="row">
                      <label>Mot de passe :</label>
                      <input
                        type="password" 
                        value={ this.state.signUp.password } 
                        onChange={ this._handleOnSignUpPasswordChange }
                      />
                    </div>
                    <div className="row">
                      <label>Confirmation mot de pass :</label>
                      <input 
                        type="password" 
                        value={ this.state.signUp.passwordCheck } 
                        onChange={ this._handleOnSignUpPasswordCheckChange }
                      />
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
              <div className="background" onClick={ () => this.setState({ isSignInDialogOpen: false }) }></div>
              <div className="content">
                <header>
                  <h2>Connexion</h2>
                </header>
                <main>
                  <div className="sign-social">
                    <GoogleLogin
                      clientId="857123691814-2gmks9c09okm1mk86nachs1vbpbk33nr.apps.googleusercontent.com"
                      onSuccess={ this._handleOnGoogleSignInResponse }
                      onFailure={ this._handleOnGoogleSignInResponse }
                      buttonText="Se connecter avec google"
                    />
                  </div>
                  <form onSubmit={ this._handleOnSignInDialogCloed }>
                    <p>ou utiliser une identification classique</p>
                    <div className="row">
                      <label>Email :</label>
                      <input
                        type="email"
                        value={ this.state.signIn.email }
                        onChange={ this._handleOnSignInEmailChange }
                      />
                    </div>
                    <div className="row">
                      <label>Mot de passe :</label>
                      <input
                        type="password"
                        value={ this.state.signIn.password }
                        onChange={ this._handleOnSignInPasswordChange }
                      />
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
        </div>
      </BrowserRouter>
    );
  }
}
