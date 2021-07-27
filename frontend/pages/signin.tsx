import auth from "../services/auth";
import {useState} from "react";
import {useRouter} from "next/router";

export default function SignIn() {
    return (
        <div>
            So, you want to sign in into SkyNET? That&apos;s awesome! Please provide your basic auth info here:
            <br/>

            <form >
                <label>Login:</label><br/>
                <input id="id" type="text"/><br/>

                <label>Password:</label><br/>
                <input id="password" type="password"/><br/>

                <button type="submit">Sign In</button>
            </form>
        </div>
    )
}
