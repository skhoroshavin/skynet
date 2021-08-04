import {FormEvent} from "react";
import {FormButton, FormInput, LinkButton} from "./common/form";
import {Card, CardDivider, CardTitle} from "./common/card";

type SignUpCardProps = {
    className?: string
}

export const SignUpCard = ({ className }: SignUpCardProps) => {
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
            <FormInput className="mt-5 w-full" id="id" type="text" placeholder="Identifier"/>
            <FormInput className="mt-4 w-full" id="password" type="password" placeholder="Password"/>
            <FormInput className="mt-4 w-full" id="password" type="password" placeholder="Repeat password"/>
            <FormButton className="mt-4 w-full">Continue</FormButton>
        </form>
    </Card>
}
