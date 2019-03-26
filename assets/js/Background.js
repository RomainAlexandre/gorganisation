import React, { Component } from 'react';

import '../css/Background.css';

class Background extends Component {
  render() {
    let imageUrl = this.props.imageUrl;

    return (
      <div className="img-box">
        <img className="background-img" src={imageUrl} alt="Background" />
      </div>
    );
  }
}

export default Background; 