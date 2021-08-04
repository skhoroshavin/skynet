import Link from "next/link";

export const FormInput = (props: any) => {
    return <input {...props} className={
        `p-2
         border rounded border-primary-400 ring-primary-200
         hover:ring-2
         focus:outline-none focus:ring
         ${props.className}`
    }/>
}

export const FormButton = (props: any) => {
    return <button {...props} className={
        `p-2
         rounded bg-primary-500
         hover:bg-primary-600
         ${props.className}`
    }/>
}

type LinkButtonProps = {
    title: string,
    href: string,
    className?: string
}

export const LinkButton = ({ title, href, className }: LinkButtonProps) => {
    return <Link href={href} passHref>
        <a className={`block text-center p-2 rounded bg-primary-500 hover:bg-primary-600 ${className}`}>
            {title}
        </a>
    </Link>
}
