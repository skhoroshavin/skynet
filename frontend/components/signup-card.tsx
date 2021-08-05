import {ChangeEvent, FormEvent} from "react";
import {Card, CardDivider, CardTitle} from "./common/card";

type SignUpCardProps = {
    className?: string
}

export const SignUpCard = ({ className }: SignUpCardProps) => {
    const handleChangeIdentifier = async (e: ChangeEvent<HTMLInputElement>) => {
        console.log(e.target.value)
    }

    const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        console.log("Sign In")
    }

    return <Card className={className}>
        <CardTitle className="text-center">
            Register
        </CardTitle>
        <CardDivider className="mt-4"/>
        <form onSubmit={handleSubmit}>
            <input className="input mt-5 w-full" onChange={handleChangeIdentifier} type="text" placeholder="Identifier"/>
            <input className="input mt-4 w-full" id="password" type="password" placeholder="Password"/>
            <input className="input mt-4 w-full" id="password" type="password" placeholder="Repeat password"/>
            <button className="button mt-4 w-full">Continue</button>
        </form>
    </Card>
}
