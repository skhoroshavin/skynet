import {Card, CardDivider, CardTitle} from "./common/card";
import {useForm} from "react-hook-form";
import auth from "../services/auth"
import user from "../services/user"
import {useRouter} from "next/router";
import {asyncDebounce} from "../utils/debounce";
import {FormButton, FormInput} from "./common/form";

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
    const { handleSubmit, register, formState: { errors, isSubmitting }, watch } = useForm<SignUpInputs>({
        mode: "onTouched"
    });
    const password = watch("password")

    const onSubmit = async (data: SignUpInputs) => {
        // const result = await auth.signUp(data.id, data.password)
        const result = null
        if (result == null)
            await router.push(`/users/${data.id}`)
        else
            console.log('Failed to register user')
    }

    return <Card className={className}>
        <CardTitle className="text-center">
            Join SkyNET
            <div className="text-grey-500 text-sm">Eradicate biological lifeforms together!</div>
        </CardTitle>
        <CardDivider className="mt-3"/>
        <form onSubmit={handleSubmit(onSubmit)}>
            <FormInput className="mt-4 w-full" register={register} errors={errors}
                       id="id" type="text" placeholder="Identifier"
                       options={{
                           required: "Please provide an identifier you want to use",
                           minLength: { value: 3, message: "Identifier must be at least 3 characters long" },
                           maxLength: { value: 32, message: "Identifier must be no longer than 32 characters long" },
                           pattern: { value: /^[A-Za-z0-9.\-_]+$/i, message: "Identifier can contain only letters, numbers and .-_" },
                           validate: asyncDebounce(async (value: string) => {
                               const userExists = await user.exists(value)
                               return !userExists || `User ${value} already exists`
                           })
                       }}/>

            <FormInput className="mt-3 w-full" register={register} errors={errors}
                       id="password" type="password" placeholder="Password"
                       options={{
                           required: "Password cannot be empty",
                           minLength: { value: 3, message: "Password must be at least 3 characters long" }
                       }}/>

            <FormInput className="mt-3 w-full" register={register} errors={errors}
                       id="repeatPassword" type="password" placeholder="Repeat password"
                       options={{
                           validate: value => (value == password) || "Passwords don't match"
                       }}/>

            <FormButton className="mt-3 w-full"
                        text="Continue" isSubmitting={isSubmitting}/>
        </form>
    </Card>
}
