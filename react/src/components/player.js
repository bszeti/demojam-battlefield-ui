/*!

=========================================================
* Black Dashboard React v1.0.0
=========================================================

* Product Page: https://www.creative-tim.com/product/black-dashboard-react
* Copyright 2019 Creative Tim (https://www.creative-tim.com)
* Licensed under MIT (https://github.com/creativetimofficial/black-dashboard-react/blob/master/LICENSE.md)

* Coded by Creative Tim

=========================================================

* The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

*/
import React from "react";

// reactstrap components
import {
    Button,
    Card,
    CardBody,
    CardFooter,
    CardText,
    Col,
    Progress
} from "reactstrap";

const nameList = [
    'python', 'java', 'javascript', 'react'
]

const getImage = (name) => {
    let index = 0;
    try {
        index = parseInt(name.replace(/[^\d.]/g, ''))
    } catch (ee) { }
    return nameList[index];
}

const getProgressBarClassName = (value, maxHealth) => {
    const perValue = (value / maxHealth) * 100
    if (perValue > 0 && perValue < 25) {
        return 'progress-bar-very-low';
    } else if (perValue >= 25 && perValue < 50) {
        return 'progress-bar-low';
    } else if (perValue >= 50 && perValue < 75) {
        return 'progress-bar-high';
    } else {
        return 'progress-bar-very-high';
    }
}

class Player extends React.Component {

    componentDidUpdate(props) {
        // console.log(props);
    }

    render() {
        const { details } = this.props;
        // const healthIndex = Math.ceil(Math.random()*5);
        const healthIndex = details.currentHealth;
        return (
            <Col md="3">
                <img
                    alt="..."
                    title={details.ready ? 'Ready' : 'Not Ready'}
                    className={details.ready ? 'status-sign status-sign' : 'status-sign-waiting status-sign'}
                    
                    // src={require("../assets/img/" + details.ready ? "ready.png" : "waiting.png")}
                    src={require(details.ready ? "../assets/img/ready.png" : "../assets/img/waiting.png")}
                />
                {details.shield && <img
                    title={'Shield'}
                    className="shield-sign"
                    alt="..."
                    src={require("../assets/img/shield.png")}
                />}
                <Card className="card-user">
                    <CardBody>
                        <CardText />
                        <div className="author">
                            <div className="block block-one" />
                            <div className="block block-two" />
                            <div className="block block-three" />
                            <div className="block block-four" />
                            <a href="#pablo" onClick={e => e.preventDefault()}>
                                <img
                                    alt="..."
                                    className="avatar"
                                    src={require("../assets/img/icons/" + getImage(details.name) + ".svg")}
                                />
                                <h5 className="title">{getImage(details.name)}</h5>
                            </a>
                        </div>
                        <div className="card-description">
                            <Progress title={'Health'} max={details.maxhealth} color={getProgressBarClassName(healthIndex, details.maxhealth)} animated value={healthIndex} />
                        </div>
                    </CardBody>
                    <CardFooter>
                        <div className={details.disqualified ? 'disqualified-text' : 'non-visible'}>
                            <p>Disqualified</p>
                        </div>
                        <div className="button-container">
                            <Button className="btn-icon btn-round">
                                <img
                                    title={'Kills'}
                                    alt="..."
                                    src={require("../assets/img/aim.png")}
                                />
                            </Button>
                            <Button className="btn-icon btn-round">
                                <img
                                    title={'Deaths'}
                                    alt="..."
                                    src={require("../assets/img/kill.jpg")}
                                />
                            </Button>
                        </div>
                        <div className="button-container">
                            <Button className="btn-icon btn-round">
                                <span className="kill-stat">{details.kill}</span>
                            </Button>
                            <Button className="btn-icon btn-round">
                                <span className="kill-stat">{details.death}</span>
                            </Button>
                        </div>
                    </CardFooter>
                    <div className={details.killedby !== '' ? 'custom-overlay display' : 'custom-overlay'}>
                        <div className="custom-overlay-text">
                            <p>KILLED</p>
                            <p>BY</p>
                            <p>{getImage(details.killedby)}</p>
                        </div>
                    </div>
                </Card>
            </Col>
        );
    }
}

export default Player;
