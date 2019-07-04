import React from 'react';
import '../styl/header.styl';

function Header() {
  return (
    <div id="header">
      <div className="header-container">

        <div className="logo-container">
          <a href="https://www.google.com">
            <div className="logo-svg"></div>
            <div className="logo-title">Math Navigator</div>
          </a>
        </div>

      </div>
    </div>
  );
}

export default Header;
