import Link from "next/link";
import React, {FormEvent} from "react";
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
            <input className="input w-full" id="id" type="text" placeholder="Identifier"/>
            <input className="input mt-4 w-full" id="password" type="password" placeholder="Password"/>
            <button className="button mt-4 w-full">Sign In</button>
            <CardDivider className="mt-6"/>
            <Link href={"/signup"} passHref>
                <a className="button mt-6 w-full bg-secondary-300 hover:bg-secondary-400">
                    Register
                </a>
            </Link>
        </form>
    </Card>
}
