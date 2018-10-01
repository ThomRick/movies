import React, { Component } from 'react';
import './game.view.css';

const gameFixture = {
  videos: [
    {
      source: 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
    },
    {
      source: 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4',
    },
    {
      source: 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerBlazes.mp4',
    },
    {
      source: 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerEscapes.mp4',
    },
    {
      source: 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerFun.mp4',
    },
  ],
}

export class GameView extends Component {
  state = {
    // videoSrc: 'test-video.mp4',
    // videoSrc: 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
    video: null,
    response: '',
  }

  componentDidMount() {
    const iterator = gameFixture.videos[Symbol.iterator]();
    this.setState({
      video: iterator.next().value,
    });
    this.interval = setInterval(() => {
      const next = iterator.next();
      if (!next.done) {
        this.setState({
          video: next.value,
        });
      } else {
        clearInterval(this.interval);
      }
    }, 2 * 1000);
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  _handleOnResponseChange = (event) => {
    const response = event.target.value;
    this.setState({
      response,
    });
  }

  render() {
    return (
      <div className="game">
        <div className="content">
          <pre>
            { JSON.stringify(this.state.video, null, 2) }
          </pre>
          {
            this.state.video !== null ?
            <video autoPlay width="640" height="380">
              <source src={ this.state.video.source } type="video/mp4"/>
            </video>
            :
            null
          }
          <form>
            <input 
              type="text"
              value={ this.state.response }
              onChange={ this._handleOnResponseChange }
            />
          </form>
        </div>
      </div>
    )
  }
}
