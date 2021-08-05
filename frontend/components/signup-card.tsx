import {Card, CardDivider, CardTitle} from "./common/card";
import {useForm} from "react-hook-form";
import auth from "../services/auth"
import user from "../services/user"
import {useRouter} from "next/router";

type SignUpCardProps = {
    className?: string
}

type SignUpInputs = {
    id: string,
    password: string,
    repeatPassword: string
}

export const SignUpCard = ({ className }: SignUpCardProps) => {
    const router = useRouter()
    const { register, handleSubmit, watch, formState: { errors } } = useForm<SignUpInputs>();
    const password = watch("password")

    const onSubmit = async (data: SignUpInputs) => {
        const result = await auth.signUp(data.id, data.password)
        if (result == null)
            await router.push(`/users/${data.id}`)
        else
            console.log('Failed to register user')
    }

    const validateUserDoesntExist = async(value: string) => {
        const userData = await user.data(value)
        if (userData instanceof Error)
            return
        return `User ${value} already exists`
    }

    return <Card className={className}>
        <CardTitle className="text-center">
            Register
        </CardTitle>
        <CardDivider className="mt-4"/>
        <form onSubmit={handleSubmit(onSubmit)}>

            <input className={`input mt-4 w-full ${errors.id && "bg-error-100"}`}
                   type="text" placeholder="Identifier"
                   {...register("id", {
                       required: {
                           value: true,
                           message: "Please provide an identifier you want to use"
                       },
                       minLength: {
                           value: 3,
                           message: "Identifier must be at least 3 characters long"
                       },
                       maxLength: {
                           value: 32,
                           message: "Identifier must be no longer than 32 characters long"
                       },
                       pattern: {
                           value: /^[A-Za-z0-9.\-_]+$/i,
                           message: "Identifier can contain only letters, numbers and ., - or _"
                       },
                       validate: validateUserDoesntExist
                   })}/>
            <div className={`absolute mt-0.5 text-xs text-error-700`}>
                {errors.id?.message}
            </div>

            <input className={`input mt-7 w-full ${errors.password && "bg-error-100"}`}
                   id="password" type="password" placeholder="Password"
                   {...register("password", {
                       required: {
                           value: true,
                           message: "Password cannot be empty"
                       },
                       minLength: {
                           value: 3,
                           message: "Password must be at least 3 characters long"
                       }
                   })}/>
            <div className={`absolute mt-0.5 text-xs text-error-700`}>
                {errors.password?.message}
            </div>

            <input className={`input mt-7 w-full ${errors.repeatPassword && "bg-error-100"}`}
                   id="password" type="password" placeholder="Repeat password"
                   {...register("repeatPassword", {
                       validate: value => (value != password) ? "Passwords don't match" : undefined
                   })}/>
            <div className={`absolute mt-0.5 text-xs text-error-700`}>
                {errors.repeatPassword?.message}
            </div>

            <button className="button mt-8 w-full">Continue</button>
        </form>
    </Card>
}
