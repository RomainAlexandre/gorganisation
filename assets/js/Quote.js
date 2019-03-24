import React, { Component } from 'react';

class Quote extends Component {
  render() {
    let quote = this.props.quote;

    return (
      <div className="Quote">
        <h1>
          {
            quote
            ? `${quote.quote} - ${quote.author}`
            : ''
          }
        </h1>
      </div>  
    );
  }
}

export default Quote; 