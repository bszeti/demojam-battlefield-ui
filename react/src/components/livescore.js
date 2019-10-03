
// reactstrap components
import {
    Row
} from "reactstrap";
import Countdown from 'react-countdown-now';

import Player from './player';

import React, { useEffect, useState, useRef } from 'react';

const _ = require('lodash');

function LiveScore() {
    const [playersData, setPlayersData] = useState({});
    const isGameOver = useRef(false);

    const timerToClearSomewhere = useRef(false)

    const audioPlayer = useRef(false)

    // const baseUrl = 'https://api.te2019.aws.redhat-demo.com:6443/apis/rhte.demojam.battlefield/v1alpha1/namespaces/visual/battlefields/'
    // const baseUrl = 'https://5d916e4c741bd4001411625c.mockapi.io/players/1';
    // const baseUrl = 'https://api.te2019.aws.redhat-demo.com:6443/apis/rhte.demojam.battlefield/v1alpha1/namespaces/visual/battlefields/demofield';
    const baseUrl = 'api/battlefield/demofield';

    function getData() {
        // fetch(baseUrl, {
        //     method: 'GET',
        //     headers: new Headers({
        //         'Authorization': 'Bearer mbLDcSbkR0uQE0Ic6QZkkVtnuuHiAgilV9SXATr6yGw'
        //     })
        // }).then(async (fetchedData) => {
        //     const dataAsJson = await fetchedData.json();
        //     var merged = _.merge(_.keyBy(dataAsJson['spec']['players'], 'name'), _.keyBy(dataAsJson['status']['scores'], 'name'));
        //     if (dataAsJson['status']['phase'] === 'done') {
        //         isGameOver.current = true;
        //     }
        //     if (
        //         !_.isEmpty(playersData) &&
        //         JSON.stringify(dataAsJson['status']) !== JSON.stringify(playersData['rawData']['status']) &&
        //         JSON.stringify(dataAsJson['spec']) !== JSON.stringify(playersData['rawData']['spec'])
        //     ) {
        //         document.getElementById('soundBar').play();
        //         setPlayersData({'rawData': dataAsJson, 'data': _.values(merged)});
        //     }
            
        // });
    }

    // Random component
    const Completionist = () => <span>Game Over !!!!!!</span>;

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
        getData();
        timerToClearSomewhere.current = setInterval(() => {
            fetch(baseUrl, {
                method: 'GET'
                //, headers: new Headers({
                //     'Authorization': 'Bearer mbLDcSbkR0uQE0Ic6QZkkVtnuuHiAgilV9SXATr6yGw'
                // }
                // )
            }).then(async (fetchedData) => {
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
        }, 1500);
        if (isGameOver.current) {
            clearTimeout(timerToClearSomewhere.current);
        }
        return () => {
            clearTimeout(timerToClearSomewhere.current);
        };
    }, [playersData]);

    return (
        <div className="main-panel">
            <div className="content">
                <Row>
                    {playersData.data && playersData['data'].map((player, key) => {
                        return <Player key={key} details = {player}></Player>
                    })}
                </Row>
                {playersData['rawData'] && playersData['rawData']['spec'] &&
                    <p className="blockquote blockquote-primary countdown-timer">
                        <Countdown renderer={renderer} date={new Date(playersData['rawData']['status']['startTime']).getTime() + playersData['rawData']['spec']['duration'] * 1000} />
                    </p>
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