import Link from "next/link";

type LinkButtonProps = {
    className?: string,
    href: string,
    text: string
}

export const LinkButton = ({className, href, text}: LinkButtonProps) => {
    return <Link href={href} passHref>
        <a className={`block
                       p-2 rounded
                       bg-primary-500
                       hover:bg-primary-600
                       text-center 
                       ${className}`}>
            {text}
        </a>
    </Link>
}
