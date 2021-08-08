import {FieldValues} from "react-hook-form/dist/types/fields";
import {FieldPath} from "react-hook-form/dist/types/utils";
import {FieldErrors, RegisterOptions, UseFormRegister} from "react-hook-form";
import {Spinner} from "./spinner";

type FormInputProps<TFieldValues extends FieldValues> = {
    className?: string
    id: FieldPath<TFieldValues>
    type?: string
    placeholder?: string
    options?: RegisterOptions<TFieldValues>
    register: UseFormRegister<TFieldValues>
    errors: FieldErrors<TFieldValues>
}

export const FormInput = <TFieldValues extends FieldValues>(props: FormInputProps<TFieldValues>) => {
    const type = props.type || "text"
    const options = props.options || {}
    const error: any = props.errors[props.id]

    return <>
        <input className={`h-10
                           p-2 rounded
                           border border-primary-400
                           ring-primary-200
                           hover:ring-2
                           focus:outline-none focus:rings
                           ${error && "bg-error-100"} ${props.className}`}
               type={type} placeholder={props.placeholder}
               {...props.register(props.id, options)}/>
        <div className={`mt-0.5 h-4 text-xs text-error-700`}>
            {error?.message}
        </div>
    </>
}

type FormButtonProps = {
    className?: string
    text: string
    isSubmitting: boolean
}

export const FormButton = ({ className, text, isSubmitting }: FormButtonProps) => {
    return <button className={`p-2 rounded
                               bg-primary-500
                               hover:bg-primary-600
                               flex justify-center
                               ${isSubmitting && "bg-primary-600"}
                               ${className}`}
                   disabled={isSubmitting}>
        <Spinner className={`h-5 w-5 my-0.5 mr-2 text-white ${isSubmitting || "hidden"}`}/>
        <span className={isSubmitting ? "hidden" : ""}>{text}</span>
    </button>
}
