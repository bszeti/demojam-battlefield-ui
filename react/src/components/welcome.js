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
import { withRouter } from "react-router-dom";
import BASE_URL from '../constant';
class Welcome extends React.Component {

    startGame() {
        fetch(BASE_URL + 'start', {
          method: 'GET'
        }).then(async (fetchedData) => {
            try {
                console.log(fetchedData)
                const dataAsJson = await fetchedData.json();
                this.props.history.push({
                    pathname: '/dashboard/' + dataAsJson
                });
            } catch(ee) {
              console.log(ee)
            }
        });
    };

    render() {
        return (
            <div className="App">
              <header className="App-header">
                <img src={require("../assets/img/battle-icon-10.jpg")} className="App-logo" alt="logo" />
                <p>
                  GREAT BATTLE OF KILLER PODS
                </p>
                <span className="App-link" onClick={this.startGame.bind(this)}>Start</span>
              </header>
            </div>
          );
    }
}

export default withRouter(Welcome);
