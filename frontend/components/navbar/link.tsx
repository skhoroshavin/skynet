import Link from "next/link";
import React, {ReactFragment} from "react";

type NavLinkProps = {
    className?: string
    href: string
    text: ReactFragment
}

export const NavLink = ({ className, href, text }: NavLinkProps) => {
    return <Link href={href} passHref>
        <a className={`h-full p-2 flex items-center hover:bg-primary-700 font-bold transition ${className}`}>
            {text}
        </a>
    </Link>
}
