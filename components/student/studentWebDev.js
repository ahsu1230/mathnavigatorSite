'use strict';
require('./studentWebDev.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import {
  description,
  sectionIg,
  sectionTesla,
  sectionProfile,
  sectionYelp
} from './webdev.js';
import { debounce } from 'lodash';
const classnames = require('classnames');

export class StudentWebDevPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      description: description,
      webdevUrl: "/class/web_design_summer_a",
      showPropzIg: false,
      showPropzTesla: false,
      showPropzProfile1: false,
      showPropzProfile2: false,
      showPropzYelp: false
    }
    this.handleScroll = this.handleScroll.bind(this);
  }

  componentDidMount() {
    this.unbindTimeout = setTimeout(function() {
      this.setState({ showPropzIg: true });
    }.bind(this), 1500);

    window.addEventListener('scroll', debounce(this.handleScroll, 200));

    if (process.env.NODE_ENV === 'production') {
      mixpanel.track("student-webdev");
    }
  }

  componentWillUnmount() {
    if (this.unbindTimeout) {
      clearTimeout(this.unbindTimeout);
    }
    window.removeEventListener('scroll', debounce(this.handleScroll, 200));
  }

  handleScroll(event) {
    var statePropzTesla = this.state.showPropzTesla;
    var statePropzProfile1 = this.state.showPropzProfile1;
    var statePropzProfile2 = this.state.showPropzProfile2;
    var statePropzYelp = this.state.showPropzYelp;

    if (statePropzTesla &&
        statePropzProfile1 &&
        statePropzProfile2 &&
        statePropzYelp) {
      return;
    }

    var showPropzTesla = statePropzTesla;
    var showPropzProfile1 = statePropzProfile1;
    var showPropzProfile2 = statePropzProfile2;
    var showPropzYelp = statePropzYelp;

    var scrollY = event.target.scrollingElement.scrollTop;
    if (!showPropzTesla && scrollY > 320) {
      showPropzTesla = true;
    }
    if (!showPropzProfile1 && scrollY > 600) {
      showPropzProfile1 = true;
    }
    if (!showPropzProfile2 && scrollY > 1100) {
      showPropzProfile2 = true;
    }
    if (!showPropzYelp && scrollY > 1500) {
      showPropzYelp = true;
    }

    var needToSet = statePropzTesla ^ showPropzTesla;
    needToSet = needToSet || (statePropzProfile1 ^ showPropzProfile1);
    needToSet = needToSet || (statePropzProfile2 ^ showPropzProfile2);
    needToSet = needToSet || (statePropzYelp ^ showPropzYelp);

    if (needToSet) {
      this.setState({
        showPropzTesla: showPropzTesla,
        showPropzProfile1: showPropzProfile1,
        showPropzProfile2: showPropzProfile2,
        showPropzYelp: showPropzYelp,
      });
    }
  }

	render() {
		return (
      <div id="view-students-wb">
        <div id="view-students-wb-container">
          <h1>
            <Link to="/programs">Programs</Link> >
            <Link to={this.state.webdevUrl}> Web Development</Link> >
            Overview
          </h1>
          <p> {this.state.description} </p>

          <SectionIG showPropz={this.state.showPropzIg}/>
          <SectionTesla showPropz={this.state.showPropzTesla}/>
          <SectionProfile showPropz1={this.state.showPropzProfile1}
            showPropz2={this.state.showPropzProfile2}/>
          <SectionYelp showPropz={this.state.showPropzYelp}/>
        </div>
      </div>
		);
	}
}

class SectionIG extends React.Component {
  render() {
    const info = sectionIg;
    var propzInfo = generatePropz(info);
    var propzStyle = {
      top: "10px",
      left: "45%"
    };

    return (
      <div className="section ig-header">
        <h2>{info.title}</h2>
        <div className="propz-container">
          <img src={info.imgSrc}/>
          <Propz info={propzInfo} style={propzStyle} direction="up"
            show={this.props.showPropz}/>
        </div>
        <div className="credit">Created by {info.student1}, {info.student2}</div>
        <p>{info.description}</p>
      </div>
    );
  }
}

class SectionTesla extends React.Component {
  render() {
    const info = sectionTesla;
    var propzInfo = generatePropz(info);
    var propzStyle = {
      top: "156px",
      left: "132%"
    };
    return (
      <div className="section tesla">
        <h2>{info.title}</h2>
        <div className="tesla-container">
          <div className="propz-container">
            <img src={info.imgSrc}/>
            <Propz info={propzInfo} style={propzStyle} direction="right"
              show={this.props.showPropz}/>
          </div>
          <div className="social-links">
            <a href="https://www.facebook.com/TeslaMoto" target="_blank">
              <div className="icon social-fb"/>
            </a>
            <a href="https://www.instagram.com/teslamotors/" target="_blank">
              <div className="icon social-ig"/>
            </a>
            <a href="https://twitter.com/tesla" target="_blank">
              <div className="icon social-tw"/>
            </a>
            <a href="https://www.youtube.com/channel/UC5WjFrtBdufl6CZojX3D8dQ" target="_blank">
              <div className="icon social-yt"/>
            </a>
          </div>
        </div>
        <div className="credit">Created by {info.student1}, {info.student2}</div>
        <p>{info.description}</p>
      </div>
    );
  }
}

class SectionProfile extends React.Component {
  render() {
    const info = sectionProfile;
    var propzInfo1 = generatePropz(info.info1);
    var propzStyle1 = {
      top: "0px",
      left: "70%"
    };
    var propzInfo2 = generatePropz(info.info2);
    var propzStyle2 = {
      top: "90%",
      left: "30%"
    };
    return (
      <div className="section profiles">
        <h2>{info.title}</h2>
        <div className="propz-container">
          <img className="img1" src={info.info1.imgSrc}/>
          <Propz info={propzInfo1} style={propzStyle1} direction="up"
            show={this.props.showPropz1}/>
          <div className="credit">Created by {info.info1.student1}, {info.info1.student2}</div>
        </div>
        <div>
          <div className="descriptions">
            <p>{info.info1.description}</p>
            <p>{info.info2.description}</p>
          </div>
          <div className="propz-container">
            <img className="img2" src={info.info2.imgSrc}/>
            <Propz info={propzInfo2} style={propzStyle2} direction="down"
              show={this.props.showPropz2}/>
            <div className="credit">Created by {info.info2.student1}, {info.info2.student2}</div>
          </div>
        </div>
      </div>
    );
  }
}

class SectionYelp extends React.Component {
  render() {
    const info = sectionYelp;
    var propzInfo = generatePropz(info);
    var propzStyle = {
      top: "164px",
      left: "80%"
    };
    return (
      <div className="section yelp">
        <h2>{info.title}</h2>
        <p>{info.description}</p>
        <div className="propz-container">
          <img src={info.imgSrc}/>
          <Propz info={propzInfo} style={propzStyle} direction="right"
            show={this.props.showPropz}/>
        </div>
        <div className="credit">Created by {info.student1}, {info.student2}</div>
      </div>
    );
  }
}

class Propz extends React.Component {
  render() {
    const info = this.props.info;
    const style = this.props.style;
    var propzClasses = classnames("propz", this.props.direction, {
      "show": this.props.show
    });

    return (
      <div className={propzClasses} style={style}>
        <div className="tri"/>
        <div>Created by {info.line1}</div>
        <div>{info.line2}</div>
      </div>
    );
  }
}

function generatePropz(info) {
  return {
    line1: info.student1,
    line2: info.student2
  };
}
