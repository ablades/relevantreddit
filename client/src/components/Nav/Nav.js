import React from 'react'
import axios from 'axios'
import { useSelector } from 'react-redux'

//Navigation for website
function Nav(props) {
    const {endpoint, userName} = props

    // name of reducer (its where the state is stored)
    const name = useSelector(state => state.name)

    //Authentication request
    function onAuth () {
        axios.get(endpoint + "/r/login").then((response) => {
            console.log(response)

            window.location.assign(response.request.responseURL);
            //window.open(response.request.responseURL, '_blank');
        })
     } 

    return (
        <nav className="navbar">
            {userName !== "" ? 
            <span>Welcome! {userName}</span> 
            : 
            <button type="submit" onClick={onAuth}>Authenticate with Reddit</button>
            }
        </nav>
    )
}

export default Nav