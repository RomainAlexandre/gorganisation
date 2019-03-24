import React, { Component } from 'react';
import Unsplash from 'unsplash-js';

import Quote from './Quote.js'
import Clock from './Clock.js'
import Background from './Background.js'

import logo from '../images/logo.svg';
import '../css/App.css';

const unsplash = new Unsplash({
  applicationId: UNSPLASH_API_ID,
  secret: UNSPLASH_API_SECRET
});

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      quote: null,
      isLoading: false,
      imageUrl: null,
    };
  }

  componentDidMount() {
    this.setState({isLoading: true});
    let quotePromise = fetch('http://quotes.rest/qod.json?category=inspire')
    let backgroundPromise = unsplash.photos.getRandomPhoto({
      collections: [827743],
      width: 1920,
      height: 1080,
      featured: true
    })

    Promise.all([quotePromise, backgroundPromise])
      .then(([quoteData, backgroundData]) => {
        debugger
        let quoteJson = quoteData.json();
        let backgroundJson = backgroundData.json();
        this.setState({
          // quote: quoteJson.contents.quotes[0],
          imageUrl: backgroundJson.urls.full,
          isLoading: false,
        })
      });
  }

  render() {

    let isLoading = this.state.isLoading;
    let quote = this.state.quote;
    let imageUrl = this.state.imageUrl;

    if (isLoading) {
      return <p>Loading ...</p>;
    }

    return (
      <div>
        <div className="App">
          <Background imageUrl={imageUrl}>
            <Clock/>
            <Quote quote={quote}/>
          </Background>
        </div>
      </div>
    );
  }
}

export default App; 