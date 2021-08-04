import {FormEvent} from "react";
import {FormButton, FormInput, LinkButton} from "./common/form";
import {Card, CardDivider} from "./common/card";

type SignInCardProps = {
    className?: string
}

export const SignInCard = ({ className }: SignInCardProps) => {
    const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        console.log("Sign In")
    }

    return <Card className={className}>
        <form onSubmit={handleSubmit}>
            <FormInput className="w-full" id="id" type="text" placeholder="Identifier"/>
            <FormInput className="mt-4 w-full" id="password" type="password" placeholder="Password"/>
            <FormButton className="mt-4 w-full">Sign In</FormButton>
            <CardDivider className="mt-6"/>
            <LinkButton className="mt-6 w-full bg-secondary-300 hover:bg-secondary-400"
                        title="Register" href="/signup"/>
        </form>
    </Card>
}
