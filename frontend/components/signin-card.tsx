import Link from "next/link";
import React from "react";
import {Card, CardDivider} from "./common/card";
import {useForm} from "react-hook-form";
import auth from "../services/auth";
import {useRouter} from "next/router";
import {FormButton, FormInput} from "./common/form";

type SignInCardProps = {
    className?: string
}

type SignInInputs = {
    id: string,
    password: string
}

export const SignInCard = ({ className }: SignInCardProps) => {
    const { handleSubmit, register, formState: { errors, isSubmitting } } = useForm<SignInInputs>({
        mode: "onTouched"
    });

    const router = useRouter()
    const onSubmit = async (data: SignInInputs) => {
        const result = await auth.signIn(data.id, data.password)
        if (result == null)
            await router.push(`/users/${data.id}`)
        else
            alert(`Failed to sign in: ${result}`)
    }

    return <Card className={className}>
        <form onSubmit={handleSubmit(onSubmit)}>
            <FormInput className="w-full" register={register} errors={errors}
                       id="id" type="text" placeholder="Identifier"
                       options={{ required: "Please provide your identifier" }}/>

            <FormInput className="mt-2 w-full" register={register} errors={errors}
                       id="password" type="password" placeholder="Password"
                       options={{ required: "Please provide your password" }}/>

            <FormButton className="mt-2 w-full"
                        text="Sign In" isSubmitting={isSubmitting}/>

            <CardDivider className="mt-4"/>
            <Link href={"/signup"} passHref>
                <a className="button mt-4 w-full bg-secondary-300 hover:bg-secondary-400">
                    Register
                </a>
            </Link>
        </form>
    </Card>
}
