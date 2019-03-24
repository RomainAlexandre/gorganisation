import React, { Component } from 'react';

class Background extends Component {
  render() {
    let imageUrl = this.props.imageUrl;

    return (
      <img src={imageUrl} alt="Background" />
    );
  }
}

export default Background; 