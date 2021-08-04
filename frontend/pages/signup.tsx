import auth from "../services/auth";
import {useState} from "react";
import {useRouter} from "next/router";
import {SignUpCard} from "../components/signup-card";

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
        <div className="flex w-full h-screen -mt-16">
             <SignUpCard className="w-1/3 mx-auto my-auto xl:w-1/4"/>
        </div>
    )
}
