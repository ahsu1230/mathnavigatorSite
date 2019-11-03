'use strict';
require('./studentProjects.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import {
  SiteDescription,
  Projects
} from './dataProjects.js';
const classnames = require('classnames');
const srcBlank = require('../../assets/blank.png');

const STUDENTS_DOMAIN = "https://www.students.andymathnavigator.com";

export class StudentProjectsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      description: SiteDescription,
      projects: Projects || []
    }
  }

  componentDidMount() {
		if (process.env.NODE_ENV === 'production') {
      mixpanel.track("student-projects");
    }
  }

	render() {
    var sections;
    if (Projects && Projects.length > 0) {
      sections = Projects.map((project, index) =>
        <Section key={index} section={project}/>
      );
    } else {
      sections = <h3>Content Coming Soon</h3>
    }

		return (
      <div id="view-students-projects">
        <div id="view-students-projects-container">
          <h1>Student Website Projects</h1>
          <p>{this.state.description}</p>
          {sections}
        </div>
      </div>
		);
	}
}

class Section extends React.Component {
  render() {
    const section = this.props.section;
    const items = section.projects.map((project, index) => (
      <li key={index}>
        <ProjectCard project={project}/>
      </li>
    ));

    return (
      <div className="section">
        <h3>{section.sectionTitle}</h3>
        <ul>
          {items}
        </ul>
      </div>
    );
  }
}

class ProjectCard extends React.Component {
  render() {
    const project = this.props.project;
    return (
      <div className="project-card">
        <div className="img-container">
          <div className="overlay"></div>
          <OverlayContent project={project}/>
          <img src={project.imgSrc || srcBlank}/>
        </div>
        <div className="info">
          <h4>{project.student1}</h4>
          <h4>{project.title}</h4>
          <div>{project.school}, {project.grade}</div>
        </div>
      </div>
    );
  }
}

class OverlayContent extends React.Component {
  render() {
    const project = this.props.project;
    const content = createContent(project);
    return (
      <div>
        {content}
      </div>
    );
  }
}

function createContent(project) {
  const fullUrl = STUDENTS_DOMAIN + project.url;
  if (project.url) {
    return (
      <div className="card-content">
        <h4>{project.title}</h4>
        <p>{project.description}</p>
        <a href={fullUrl} target="_blank">
          <button>View</button>
        </a>
      </div>
    );
  } else {
    return (
      <div className="card-content">
        <h4>Coming Soon</h4>
      </div>
    );
  }
}
