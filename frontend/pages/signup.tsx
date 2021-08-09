import {SignUpCard} from "../components/signup-card";
import {AuthRedirect} from "../components/auth-redirect";

export default function SignUp() {
    return <AuthRedirect>
        <div className="fullscreen flex">
            <SignUpCard className="w-96 m-auto"/>
        </div>
    </AuthRedirect>
}
