import auth from "../services/auth";
import {useState} from "react";
import {useRouter} from "next/router";
import {SignUpCard} from "../components/signup-card";

export default function SignUp() {
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
        <div className="fullscreen flex">
             <SignUpCard className="w-96 m-auto"/>
        </div>
    )
}
