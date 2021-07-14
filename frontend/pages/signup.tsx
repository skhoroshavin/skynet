import auth from "../services/auth";
import {useState} from "react";
import {useRouter} from "next/router";

export default function SignUp() {
    const [error, setError] = useState<string>()
    const router = useRouter()

    const signUp = async (event: any) => {
        event.preventDefault()
        const id = event.target.id.value
        const password = event.target.password.value
        const err = await auth.signUp(id, password)
        if (err != null) {
            setError(`Failed to signup: ${err}`)
        }
        else {
            await router.push('/')
        }
    }

    return (
        <div>
            So, you want to sign in into SkyNET? That&apos;s awesome! Please provide your basic auth info here:
            <br/>

            <form onSubmit={signUp}>
                <label>Login:</label><br/>
                <input id="id" type="text"/><br/>

                <label>Password:</label><br/>
                <input id="password" type="password"/><br/>

                <label>Password (again):</label><br/>
                <input id="password_check" type="password"/><br/>

                <button type="submit">Sign Up</button>
            </form>

            {error && <div className="error"> {error} </div>}
        </div>
    )
}
