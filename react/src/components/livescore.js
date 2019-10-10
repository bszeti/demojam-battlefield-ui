
// reactstrap components
import {
    Row
} from "reactstrap";
import Countdown from 'react-countdown-now';

import Player from './player';
import BASE_URL from '../constant';
import React, { useEffect, useState, useRef } from 'react';

const _ = require('lodash');

function LiveScore(props) {
    const [playersData, setPlayersData] = useState({});
    const isGameOver = useRef(false);

    const timerToClearSomewhere = useRef(false)
    console.log(props);
    const audioPlayer = useRef(false);

    let baseUrl;
    if (
        typeof(props.match) === 'undefined' || 
        typeof(props.match.params) === 'undefined' ||
        typeof(props.match.params.game) === 'undefined'
        ) {
        props.history.push('/');
    } else {
        baseUrl = `${BASE_URL}/battlefield/${props.match.params.game}`;
    }

    function getWinner() {
        let playerWinner = _.maxBy(playersData.data, function(o) { return o.kill; })
        return playerWinner['name'];
    }

    function getWinner() {
        let playerWinner = _.maxBy(playersData.data, function(o) { return o.kill; })
        return playerWinner['name'];
    }


    // Random component
    const Completionist = () => {
        return (
            <div className="grid-container">
                <div className="grid-item">
                    <img
                        title={'Kills'}
                        alt="..."
                        src={require("../assets/img/giphy.gif")}
                    />
                </div>
                <div className="grid-item winner-css">
                    <span style={{fontSize: '40px'}}>Winner is {getWinner()}</span>
                    <br></br>
                    <span className="App-link" onClick={() => props.history.push({pathname: '/'})}>Reset</span>
                </div> 
            </div>
        )
    };

    // Renderer callback with condition
    const renderer = ({ hours, minutes, seconds, completed }) => {
        if (completed) {
        // Render a completed state
        return <Completionist />;
        } else {
        // Render a countdown
        return <span>{minutes < 10 ? `0${minutes}` : minutes }:{seconds < 10 ? `0${seconds}` : seconds}</span>;
        }
    };
    

    useEffect(() => {
        timerToClearSomewhere.current = setInterval(() => {
            fetch(baseUrl, {
                method: 'GET',
                headers: new Headers({'content-type': 'application/json'})
            }).then(async (fetchedData) => {
                console.log(fetchedData);
                const dataAsJson = await fetchedData.json();
                var merged = _.merge(_.keyBy(dataAsJson['spec']['players'], 'name'), _.keyBy(dataAsJson['status']['scores'], 'name'));
                if (dataAsJson['status']['phase'] === 'done') {
                    isGameOver.current = true;
                }
                if (
                    !_.isEmpty(playersData) &&
                    JSON.stringify(dataAsJson['status']) !== JSON.stringify(playersData['rawData']['status']) &&
                    JSON.stringify(dataAsJson['spec']) !== JSON.stringify(playersData['rawData']['spec'])
                ) {
                    console.log('dsads')
                    try{
                        if (audioPlayer.current) {
                            console.log(audioPlayer.current)
                            console.log('THIS IS AUDIO')
                            // audioPlayer.current.stop();
                        }
                        audioPlayer.current = document.getElementById('soundBar').play();
                        // In browsers that don’t yet support this functionality,
                        // playPromise won’t be defined.
                        if (audioPlayer.current !== undefined) {
                            audioPlayer.current.then(function(data) {
                                console.log(data);
                                console.log('THIS IS SUCCESS')
                            // Automatic playback started!
                            }).catch(function(error) {
                                console.log(error)
                                console.log('THIS IS FAIL')
                            // Automatic playback failed.
                            // Show a UI element to let the user manually start playback.
                            });
                        }
                    } catch(e){
                        console.log(e);
                    }
                    // audioPlayer.current = 
                }
                setPlayersData({'rawData': dataAsJson, 'data': _.values(merged)});
            });
        }, 500); // 1.5 second polling
        if (isGameOver.current) {
            clearTimeout(timerToClearSomewhere.current);
        }
        return () => {
            clearTimeout(timerToClearSomewhere.current);
        };
    }, [playersData, baseUrl]);

    return (
        <div className="main-panel">
            <div className="content">
                <Row>
                    {playersData.data && playersData['data'].map((player, key) => {
                        return <Player key={key} game = {playersData.rawData.metadata.name} details = {player}></Player>
                    })}
                </Row>
                {playersData['rawData'] && playersData['rawData']['spec'] &&
                    <div className="blockquote blockquote-primary countdown-timer">
                        <Countdown renderer={renderer} date={new Date(playersData['rawData']['status']['startTime']).getTime() + playersData['rawData']['spec']['duration'] * 1000} />
                    </div>
                }
            </div>
            <DrumPad
                //   name={key} 
                  soundFile="https://s3.amazonaws.com/freecodecamp/simonSound4.mp3"
                //   key={key}
                />
        </div>
    );
}

function DrumPad(props) {
    return (
      <div className="drum-pad">
        <audio className="clip" id='soundBar' src={props.soundFile} type="audio/mp3" ></audio>
      </div>
    )
  }

export default LiveScore;