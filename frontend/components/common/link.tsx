import React, {FC} from "react";
import Link from "next/link";

type LinkButtonProps = {
    className?: string,
    href: string,
    text: string
}

export const LinkButton: FC<LinkButtonProps> = ({className, href, text}) => {
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
