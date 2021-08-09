import {SignInCard} from "../components/signin-card";
import {AuthRedirect} from "../components/auth-redirect";

export default function SignIn() {
    return <AuthRedirect>
        <div className="self-stretch flex flex-col justify-evenly flex-1 text-center font-mono text-primary-900">
            <h1 className="text-6xl font-bold">
                SkyNET
            </h1>
            <h3 className="text-xl">
                Please identify yourself or be immediately dismantled for further investigation
            </h3>
        </div>
        <SignInCard className="ml-4 w-96"/>
    </AuthRedirect>
}
